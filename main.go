package main

import (
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/api"
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/internal/config"
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/internal/database"
	"github.com/joho/godotenv"
)

func init() {
	//** load .env file anywhere from any directory **//
	godotenv.Load()
}

func main() {
	//** Database connection **/
	config.GetDBConfig()
	database.Client.ConnectDb()
	defer database.Client.DisconnetDb()

	//** API connect (router) **/
	api.Init()
	api.Router.RoutersEstablishment()
	api.Router.Run()

}
