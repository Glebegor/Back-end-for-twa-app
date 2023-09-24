package handlers

import (
	services "hackthon_back/package/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{service: serv}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORS())
	apiV2 := router.Group("/api/v2")
	{
		rooms := apiV2.Group("room")
		{
			rooms.POST("", h.)           // return room - wallet, telegram
			rooms.POST("/join", h.)      // join to the room - room_id, wallet, telegram
			rooms.POST("/increment", h.) // getting increment - room_id, wallet, telegram

		}
	}
	return router
}
