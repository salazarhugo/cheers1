package main

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}

	message := fmt.Sprintf("Hello %s", name)

	return message, nil
}
