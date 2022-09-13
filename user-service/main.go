package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"salazar/cheers/user/proto/userpb"
)

func main() {
	log.SetFlags(0)

	log.Printf("grpc-user: starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, newServer())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

//e := echo.New()
//
//log.SetFlags(0)
//
//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cc := &utils.CustomContext{Context: c, Driver: utils.GetDriver()}
//		return next(cc)
//	}
//})
//
//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//	AllowOrigins: []string{"https://web-cheers.web.app", "http://localhost:4200"},
//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
//}))
//
//e.Use(auth.UserInfoMiddleware)
//
//port := os.Getenv("PORT")
//if port == "" {
//	port = "8080"
//	log.Printf("defaulting to port %s", port)
//}
//
//e.GET("/", user.GetUser)
//e.PATCH("/", user.UpdateUser)
//e.Logger.Fatal(e.Start(":" + port))
