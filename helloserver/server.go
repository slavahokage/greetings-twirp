package helloserver

import (
	"context"
	pb "helloservice"
	"math/rand"
	"time"
)


type Server struct {}

func (s *Server) GetRandomGreeting(ctx context.Context, arguments *pb.Arguments) (greeting *pb.GreetingMessage, err error) {
	hello := getVariousGreetings();
	return &pb.GreetingMessage{
		Greetings: hello,
	}, nil
}

func getVariousGreetings() string{
	greetings:= []string{"Hi","Hola","Привет","Salut","Hej","здрасти","Hallo","Здраво","Прывітанне","Zdravo","Hei","Ahoj"}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return  greetings[r.Intn(len(greetings)-1)]
}