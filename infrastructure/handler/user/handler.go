package user

import (
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	responseError "github.com/Melany751/house-match-server/domain/services/error"
	"github.com/Melany751/house-match-server/domain/services/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase user.UseCaseUser
}

func newHandler(useCase user.UseCaseUser) handler {
	return handler{useCase}
}

func (h handler) getById(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.GetById(uid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(responseError.OK(m))
}

func (h handler) getAll(c *gin.Context) {
	ms, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ms)
}

func (h handler) create(c *gin.Context) {
	var req model.User
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Error read body")
	}

	m, err := h.useCase.Create(req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(responseError.Created(m))
}

func (h handler) update(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	var req model.User
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Error read body")
	}

	created, err := h.useCase.Update(uid, req)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(responseError.Updated(created))
}

func (h handler) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	deleted, err := h.useCase.Delete(uid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(responseError.Deleted(deleted))
}
