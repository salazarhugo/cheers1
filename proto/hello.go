package proto

import (
	"fmt"
	v1 "github.com/salazarhugo/cheers1/proto/out/cheers/api/v1"
	"sync"
)

func NewServer() *MainApiServiceServer {
	s := &MainApiServiceServer{}
	fmt.Println(s)
	return s
}

type MainApiServiceServer struct {
	v1.UnimplementedMainServer
	mu sync.Mutex
}
