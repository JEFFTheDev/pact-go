package main

import "github.com/JEFFTheDev/pact-go/examples/consumer/goconsumer"

func main() {
	client := goconsumer.Client{
		Host: "http://localhost:8080",
	}
	client.Run()
}
