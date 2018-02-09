package main

import (
	"rpc/haberdasher"

	"github.com/twitchtv/twirp"

	"context"
	"fmt"
	"net/http"
)

func main() {
	client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 12})
	fmt.Println(hat, err)
	if err != nil {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("Oh no " + twerr.Msg())
			default:
				fmt.Println(twerr.Error())
			}
		}
		fmt.Println(hat)
		return
	}
}
