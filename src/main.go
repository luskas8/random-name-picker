package main

import (
	"fmt"
	"module/internal/request"
	"module/setup/settings"
)

func main() {
	settings.Load()

	client := request.Client{Base: "https://randommer.io", Token: settings.RANDOMMER_API_KEY}
	response := client.Get("/api/Name?nameType=fullname&quantity=5")

	fmt.Println(response)
}
