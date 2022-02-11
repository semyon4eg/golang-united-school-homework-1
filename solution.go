package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	helloMessage := emoji.Sprint(GetMessage())
	fmt.Println(helloMessage)
}

func GetMessage() string {
	return "Hello :world_map:!"
}
