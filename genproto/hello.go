package proto

import (
	"fmt"
	proto "github.com/salazarhugo/cheers1/protogen/cheers/api/v1"
	"sync"
)

func NewServer() *MainApiServiceServer {
	s := &MainApiServiceServer{}
	fmt.Println(s)
	return s
}

type MainApiServiceServer struct {
	proto.UnimplementedMainServer
	mu sync.Mutex
}
