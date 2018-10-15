package main

import (
	"context"
	"fmt"
	"helloservice"
	"net/http"
	"os"
	"strconv"
)

func main() {
	client := hello.NewGreetingProtobufClient("http://localhost:8080", &http.Client{})
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("something go wrong")
	}
	for i := 0; i < arg; i++ {
		greetings, err := client.GetRandomGreeting(context.Background(), &hello.Arguments{})
		if err != nil {
			fmt.Printf("oh no: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%+v \n", greetings.Greetings)
	}
}
