package handlers

import (
	shotball "hackthon_back"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoomCreate(c *gin.Context) {
	var input shotball.CreateData
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.service.Rooms.CreateRoom(input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
}
func (h *Handler) RoomJoin(c *gin.Content) {
	var input shotball.JoinData

}
func (h *Handler) RoomIncrement(c *gin.Content) {
	var input shotball.IncrementData

}
