package handler

import (
	"github.com/Melany751/house-match-server/domain/model"
	routerRole "github.com/Melany751/house-match-server/infrastructure/handler/role"
	routerUser "github.com/Melany751/house-match-server/infrastructure/handler/user"
)

func InitRoutes(specification model.RouterSpecification) {
	routerUser.NewRouter(specification)
	routerRole.NewRouter(specification)
}
