package main

import (
	"fmt"

	"github.com/aventurella/go-xpc/xpc"
)

type Sample struct{}

func (s *Sample) Test(args *[]interface{}, reply *interface{}) error {
	fmt.Println("CALLED SAMPLE.TEST!!!")
	fmt.Println(args)
	return nil
}

func main() {
	sample := new(Sample)
	xpc.Register(sample)
	xpc.Start()
}
