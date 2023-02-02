package handler

import (
	"github.com/Melany751/house-match-server/domain/model"
	routerModule "github.com/Melany751/house-match-server/infrastructure/handler/module"
	routerRole "github.com/Melany751/house-match-server/infrastructure/handler/role"
	routerRoleView "github.com/Melany751/house-match-server/infrastructure/handler/roleview"
	routerUser "github.com/Melany751/house-match-server/infrastructure/handler/user"
	routerUserRole "github.com/Melany751/house-match-server/infrastructure/handler/userrole"
	routerView "github.com/Melany751/house-match-server/infrastructure/handler/view"
)

func InitRoutes(specification model.RouterSpecification) {
	routerUser.NewRouter(specification)
	routerRole.NewRouter(specification)
	routerModule.NewRouter(specification)
	routerView.NewRouter(specification)
	routerRoleView.NewRouter(specification)
	routerUserRole.NewRouter(specification)
}
