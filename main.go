package main

import (
	"fmt"
	"github.com/rajhawaldar/learn-go/web"
	"github.com/rajhawaldar/learn-go/model"
)

func main() {
	fmt.Println("Hello Go!!")
	fmt.Println(web.MyProfile("raj"))
	fmt.Println(model.GetConversation())
}