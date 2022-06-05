package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"salazar/cheers/notification/usereventpb"
	"strings"
	"sync"
)

func likePostNotification(c echo.Context, userEvent *usereventpb.UserEvent) {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
	defer session.Close()

	cypher := `MATCH (p:Post {id: $postId})<-[:POSTED]-(author:User)
				MATCH (u: User { id: $userId}) 
			    RETURN { authorId: author.id, tokens: author.registrationTokens, username: u.username, 
				avatar: u.profilePictureUrl }`

	params := map[string]interface{}{
		"userId": userEvent.UserId,
		"postId": userEvent.PostId,
	}

	result, err := session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return
	}

	if !result.Next() {
		return
	}
	record := result.Record()
	rMap := record.Values[0].(map[string]interface{})

	authorId := rMap["authorId"].(string)
	rTokens := rMap["tokens"].([]interface{})
	var tokens []string
	for _, token := range rTokens {
		tokens = append(tokens, token.(string))
	}

	var wg sync.WaitGroup
	wg.Add(1)

	notification := &Notification{
		ReceiverId: authorId,
		Tokens:     tokens,
		Title:      "Cheers",
		Body:       rMap["username"].(string) + " liked your post.",
		Avatar:     rMap["avatar"].(string),
	}

	go SendNotification(notification, &wg)

	wg.Wait()
	return
}

func postNotification(c echo.Context, userEvent *usereventpb.UserEvent) {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
	defer session.Close()

	cypher := `MATCH (p:Post {id: $postId})<-[:POSTED]-(author:User)
				MATCH (f:User)-[:FOLLOWS]->(author)
				WITH p, author, collect(DISTINCT [f.id, f.registrationTokens]) as tok
			    RETURN { followersWithTokens: tok, username: author.username, locationName: p.locationName,
				avatar: author.profilePictureUrl, beverage: p.beverage, caption: p.caption, type: p.type }`

	params := map[string]interface{}{
		"userId": userEvent.UserId,
		"postId": userEvent.PostId,
	}

	result, err := session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return
	}

	if !result.Next() {
		return
	}
	record := result.Record()
	rMap := record.Values[0].(map[string]interface{})
	followersWithTokens := rMap["followersWithTokens"].([]interface{})
	beverage := rMap["beverage"].(string)
	username := rMap["username"].(string)

	var wg sync.WaitGroup
	for _, followerWithTokens := range followersWithTokens {
		followerWithTokens := followerWithTokens.([]interface{})
		followerId := followerWithTokens[0].(string)
		tokens := make([]string, 0)
		for _, token := range followerWithTokens[1].([]interface{}) {
			tokens = append(tokens, token.(string))
		}
		if len(tokens) == 0 {
			continue
		}

		body := username + " just posted a photo."
		if beverage != "" && beverage != "NONE" {
			body = username + " is drinking " + strings.ToLower(beverage)
		}
		notification := &Notification{
			ReceiverId: followerId,
			Tokens:     tokens,
			Title:      "Cheers",
			Body:       body,
			Avatar:     rMap["avatar"].(string),
			ChannelId:  "NEW_POST_CHANNEL",
		}

		wg.Add(1)
		go SendNotification(notification, &wg)
	}

	wg.Wait()
	return
}

func followNotification(c echo.Context, userEvent *usereventpb.UserEvent) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
	defer session.Close()

	cypher := `MATCH (u: User {id: $userId})-[:FOLLOWS]->(f: User {id: $otherUserId}) 
                 RETURN { avatar: u.profilePictureUrl, username: u.username, tokens: f.registrationTokens, otherUserId: f.id }`

	params := map[string]interface{}{
		"userId":      userEvent.UserId,
		"otherUserId": userEvent.OtherUserId,
	}

	result, err := session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	if !result.Next() {
		return errors.New("expected one record was 0 or greater than 1")
	}
	rMap := result.Record().Values[0].(map[string]interface{})

	var tokens []string
	for _, token := range rMap["tokens"].([]interface{}) {
		tokens = append(tokens, token.(string))
	}
	notification := &Notification{
		ReceiverId: rMap["otherUserId"].(string),
		Tokens:     tokens,
		Title:      "Cheers",
		Body:       rMap["username"].(string) + " started following you.",
		Avatar:     rMap["avatar"].(string),
		ChannelId:  "NEW_FOLLOW_CHANNEL",
	}

	var wg sync.WaitGroup
	wg.Add(1)
	SendNotification(notification, &wg)

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
