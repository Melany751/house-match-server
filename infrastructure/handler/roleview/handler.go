package roleview

import (
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	roleView "github.com/Melany751/house-match-server/domain/services/roleview"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	useCase roleView.UseCaseRoleView
}

func newHandler(useCase roleView.UseCaseRoleView) handler {
	return handler{useCase}
}

func (h handler) getByIds(c *gin.Context) {
	roleID := c.Param("roleID")
	roleUid, err := uuid.Parse(roleID)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}
	viewID := c.Param("viewID")
	viewUid, err := uuid.Parse(viewID)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.GetByIDs(roleUid, viewUid)
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

func (h handler) assignment(c *gin.Context) {
	var req model.RoleView
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Error read body")
	}

	id, err := h.useCase.Assignment(req)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, id)
}

func (h handler) update(c *gin.Context) {
	var req model.RoleView
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Error read body")
	}

	created, err := h.useCase.Update(req)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, created)
}

func (h handler) delete(c *gin.Context) {
	roleID := c.Param("roleID")
	roleUid, err := uuid.Parse(roleID)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}
	viewID := c.Param("viewID")
	viewUid, err := uuid.Parse(viewID)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	deleted, err := h.useCase.Delete(roleUid, viewUid)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, deleted)
}
