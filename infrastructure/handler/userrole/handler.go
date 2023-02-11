package userrole

import (
	"fmt"
	userRole "github.com/Melany751/house-match-server/domain/services/userrole"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase userRole.UseCaseUserRole
}

func newHandler(useCase userRole.UseCaseUserRole) handler {
	return handler{useCase}
}

func (h handler) getByIds(c *gin.Context) {
	userId := c.Param("userId")
	userUid, err := uuid.Parse(userId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}
	roleId := c.Param("roleId")
	roleUid, err := uuid.Parse(roleId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.GetByIDs(userUid, roleUid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, m)
}

func (h handler) getAll(c *gin.Context) {
	userId := c.Param("userId")
	userUid, err := uuid.Parse(userId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	ms, err := h.useCase.GetAllByUserID(userUid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ms)
}

func (h handler) create(c *gin.Context) {
	userId := c.Param("userId")
	userUid, err := uuid.Parse(userId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}
	roleId := c.Param("roleId")
	roleUid, err := uuid.Parse(roleId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	id, err := h.useCase.Assignment(userUid, roleUid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, id)
}

func (h handler) delete(c *gin.Context) {
	userId := c.Param("userId")
	userUid, err := uuid.Parse(userId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}
	roleId := c.Param("roleId")
	roleUid, err := uuid.Parse(roleId)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	deleted, err := h.useCase.Delete(userUid, roleUid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, deleted)
}