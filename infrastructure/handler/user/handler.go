package user

import (
	"fmt"
	"github.com/Melany751/house-match-server/domain/services/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(useCase user.UseCase) handler {
	return handler{useCase}
}

func (h handler) getById(c *gin.Context) {
	//var req model.UserRequestById
	//if err := c.BindJSON(&req); err != nil {
	//	fmt.Printf("no se pudo convertir")
	//}

	id := c.Param("id")

	uuid, err := uuid.Parse(id)
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

func (h handler) getAll(c *gin.Context) {
	ms, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ms)
}
