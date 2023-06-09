package api

import (
	"os"

	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/api/routes"
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/internal/database"
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/transport"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Current *gin.Engine
}

var Router GinRouter

func Init() {
	Router.Current = gin.Default()

}

func (gr *GinRouter) RoutersEstablishment() {
	Router.Current.Use(transport.Recover(database.Client)).Use(transport.Cors)
	apiRouter := Router.Current.Group("/api")

	/*
		! Flow of the code:
		* 1. go\api\init.api.go
		* 2. go\api\routes\dictionary.routes.go
		* 3. go\api\handlers\dictionary.handles.go
		* 4. go\internal\services\dictionary.services.go
		TODO:
		* 1. Go to go\internal\services\dictionary.services.go to code services of dictionary
		* 2. Go to go\api\handlers\dictionary.handles.go to code handlers of dictionary (call services)
		* 3. Go to go\api\routes\dictionary.routes.go to code routes of dictionary (call handlers)
		* 4. Go to go\api\init.api.go to code routers of dictionary (call routes)

		! Similar flow for other services
	*/
	// routes.APIDocumentationRoutes(Router.Current)
	routes.RootRoutes(apiRouter)
	routes.DictonaryRoutes(apiRouter, database.Client.Source)
	routes.QuizRoutes(apiRouter, database.Client.Source)
	routes.QuestionRoutes(apiRouter, database.Client.Source)
	routes.BottleShopRoutes(apiRouter, database.Client.Source)
	routes.ArticleRoutes(apiRouter, database.Client.Source)
}

func (gr *GinRouter) Run() {
	port := os.Getenv("PORT")
	Router.Current.Run(":" + port)
}
