package proto

import (
	"fmt"
	"sync"
)

func NewServer() *MainApiServiceServer {
	s := &MainApiServiceServer{}
	fmt.Println(s)
	return s
}

type MainApiServiceServer struct {
	UnimplementedMainServer
	mu sync.Mutex
}
