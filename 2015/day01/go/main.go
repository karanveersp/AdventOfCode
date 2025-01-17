package main

import "fmt"

func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}
