package main

import (
	"err_test/client"
	"err_test/handler"
	"err_test/userstory"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	c := client.NewClient(log)
	uStory := userstory.NewUserStory(c)
	handler := handler.NewHandler(uStory)
	fmt.Println(handler.Handle())
}
