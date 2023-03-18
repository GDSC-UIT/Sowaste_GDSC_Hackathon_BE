package routes

import (
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/api/handlers"
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ArticleRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.ArticleHandlers{
		Handler: services.ArticleServices{
			Db: db,
		},
	}
	articles := group.Group("/articles")
	{
		articles.GET("", handler.GetArticles)
		articles.GET("/:id", handler.GetAnArticle)
		articles.POST("", handler.CreateAnArticle)
		articles.PUT("/:id", handler.UpdateAnArticle)
		articles.DELETE("/:id", handler.DeleteAnArticle)
	}
}
