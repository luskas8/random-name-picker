package main

import (
	"fmt"

	"github.com/luskas8/randomnamepicker/src/modules/request"
	"github.com/luskas8/randomnamepicker/src/modules/settings"
)

func main() {
	settings.Load()

	client := request.Client{Base: "https://randommer.io", Token: settings.RANDOMMER_API_KEY}
	response := client.Get("/api/Name?nameType=fullname&quantity=5")

	fmt.Println(response)
}
