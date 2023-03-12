package rechargecoins

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	firebase "firebase.google.com/go/v4"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/api/androidpublisher/v3"
	"log"
	"net/http"
	"strings"
)

type rechargeRequest struct {
	PackageName   string `json:"packageName"`
	ProductId     string `json:"productId"`
	PurchaseToken string `json:"purchaseToken"`
}

var app *firebase.App

func init() {
	functions.HTTP("RechargeCoins", RechargeCoins)
	app = utils.InitializeAppDefault()
}

func RechargeCoins(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	p := rechargeRequest{}

	// Parse the request body to get the parameters.
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("json.NewDecoder: %v", err)
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	// Validate parameters
	if p.ProductId == "" || p.PurchaseToken == "" {
		s := "missing 'productId' or 'purchaseToken' parameter"
		log.Println(s)
		http.Error(w, s, http.StatusBadRequest)
		return
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return
	}

	// Authenticate the user with the token passed in the header.
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Validate the idToken
	idToken := strings.TrimPrefix(authToken, "Bearer ")
	token, err := auth.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}
	userID := token.UID

	// Validate the purchase
	err = VerifyPurchase(p.PackageName, p.ProductId, p.PurchaseToken)
	if err != nil {
		http.Error(w, "invalid purchase: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Deliver the purchase
	err = DeliverContent(p.ProductId, userID)
	if err != nil {
		http.Error(w, "couldn't deliver content to user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Acknowledge the purchase
	err = AcknowledgePurchase(p.PackageName, p.ProductId, p.PurchaseToken)
	if err != nil {
		http.Error(w, "couldn't acknowledge purchase", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func DeliverContent(
	productId string,
	userID string,
) error {
	ctx := context.Background()

	// Get a Firestore client.
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	data := map[string]interface{}{
		"coins": firestore.Increment(productIdToCoins(productId)),
	}
	_, err = client.Collection("accounts").Doc(userID).Set(ctx, data, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func VerifyPurchase(
	packageName string,
	productId string,
	purchaseToken string,
) error {
	log.Println("Verifying purchase with parameters: ", packageName, productId, purchaseToken)
	// Initialize the Android Publisher API client.
	ctx := context.Background()
	svc, err := androidpublisher.NewService(ctx)
	if err != nil {
		return err
	}

	// Verify the purchase using the Android Publisher API.
	response, err := svc.Purchases.Products.Get(packageName, productId, purchaseToken).Do()
	if err != nil {
		return err
	}

	// Check that the product is purchased.
	if response.PurchaseState != 0 {
		return errors.New("product not purchased")
	}

	// Check that the product is not already acknowledged.
	if response.AcknowledgementState != 0 {
		return errors.New("product is already acknowledged")
	}

	return nil
}

func AcknowledgePurchase(
	packageName string,
	productId string,
	purchaseToken string,
) error {
	// Initialize the Android Publisher API client.
	ctx := context.Background()
	svc, err := androidpublisher.NewService(ctx)
	if err != nil {
		return err
	}

	// Acknowledge the purchase using the Android Publisher API.
	err = svc.Purchases.Products.Acknowledge(packageName, productId, purchaseToken, nil).Do()
	if err != nil {
		return err
	}

	// Purchase acknowledged successfully.
	return nil
}

func productIdToCoins(productId string) int64 {
	switch productId {
	case "coins_36":
		return 36
	case "coins_70":
		return 70
	case "coins_350":
		return 350
	case "coins_700":
		return 700
	case "coins_1400":
		return 1400
	case "coins_3500":
		return 3500
	case "coins_7000":
		return 7000
	case "coins_17500":
		return 17500
	default:
		return 0
	}
}
