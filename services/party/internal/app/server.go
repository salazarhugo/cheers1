package app

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"sync"
)

func NewServer() *MainApiServiceServer {
	s := &MainApiServiceServer{
		v1.UnimplementedMainServer{},
		sync.Mutex{},
		utils.GetSession(utils.GetDriver()),
	}
	fmt.Println(s)
	return s
}

type MainApiServiceServer struct {
	v1.UnimplementedMainServer
	mu sync.Mutex
	neo4j.Session
}
