package handlers

import (
	shotball "hackthon_back"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoomCreate(c *gin.Context) {
	var input shotball.CreateData

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	data, err := h.service.Rooms.CreateRoom(input)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":             data.Id,
		"shortId":        data.ShortId,
		"roundDuration":  data.RoundDuration,
		"roomDuration":   data.RoomDuration,
		"roomStartDate":  data.RoomStartDate,
		"roundStartDate": data.RoundStartDate,
	})
}

func (h *Handler) RoomJoin(c *gin.Context) {
	var input shotball.JoinData

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.Rooms.JoinRoom(input)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":             data.Id,
		"shortId":        data.ShortId,
		"roundDuration":  data.RoundDuration,
		"roomDuration":   data.RoomDuration,
		"roomStartDate":  data.RoomStartDate,
		"roundStartDate": data.RoundStartDate,
	})
}

func (h *Handler) RoomIncrement(c *gin.Context) {
	var input shotball.IncrementData

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Rooms.IncrementingRoom(input)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
