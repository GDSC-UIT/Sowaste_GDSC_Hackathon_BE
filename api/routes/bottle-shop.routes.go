package routes

import (
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/api/handlers"
	"github.com/GDSC-UIT/Sowaste_GDSC_Hackathon_BE/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BottleShopRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.BottleShopHandlers{
		Handler: services.BottleShopServices{
			Db: db,
		},
	}
	dictionaries := group.Group("/bottle-shops")
	{
		dictionaries.GET("", handler.GetBottleShops)
		dictionaries.GET("/:id", handler.GetABottleShop)
		dictionaries.POST("", handler.CreateABottleShop)
		dictionaries.PUT("/:id", handler.UpdateABottleShop)
		dictionaries.DELETE("/:id", handler.DeleteABottleShop)
	}
}
