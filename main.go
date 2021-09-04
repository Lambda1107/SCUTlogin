package main

import (
	"SCUTlogin/server"
	"SCUTlogin/utils"
)

func main() {
	config := utils.GetConfig()

	// core
	// httpClient := server.Login("url", "id", "passwd")
	// fmt.Println(httpClient.Jar)

	// healthy report
	server.HealthReport(config.Scode, config.Password)
}
