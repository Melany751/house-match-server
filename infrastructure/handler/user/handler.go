package user

import (
	"fmt"
	"github.com/Melany751/house-match-server/domain/services/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Request struct {
	Id string `json:"id"`
}

type handler struct {
	useCase user.UseCase
}

func newHandler(useCase user.UseCase) handler {
	return handler{useCase}
}

func (h handler) getById(c *gin.Context) {
	var req Request
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("no se pudo convertir")
	}

	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.GetById(uuid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, m)
}
