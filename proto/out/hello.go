package out

import (
	"fmt"
	v1 "github.com/salazarhugo/cheers1/proto/out/cheers/api/v1"
	"sync"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func newServer() *mainApiServiceServer {
	s := &mainApiServiceServer{}
	fmt.Println(s)
	return s
}

type mainApiServiceServer struct {
	v1.UnimplementedMainServer
	mu sync.Mutex
}
