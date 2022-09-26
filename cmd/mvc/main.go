package main

import "github.com/tsadamori/go-app/controller"

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}