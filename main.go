package main

import (
	"io"

	"github.com/dywoq/dywoqgame/stream"
)

func main() {
	s := stream.NewStream(10)
	io.WriteString(s, "Hi!")
	io.WriteString(s, "Look!")

	for !s.Written() {
		stream.Println(s)
	}
}
