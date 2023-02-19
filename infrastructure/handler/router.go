package handler

import (
	"github.com/Melany751/house-match-server/domain/model"
	routerLocationPerson "github.com/Melany751/house-match-server/infrastructure/handler/locationperson"
	routerModule "github.com/Melany751/house-match-server/infrastructure/handler/module"
	routerPerson "github.com/Melany751/house-match-server/infrastructure/handler/person"
	routerProperty "github.com/Melany751/house-match-server/infrastructure/handler/property"
	routerRole "github.com/Melany751/house-match-server/infrastructure/handler/role"
	routerRoleView "github.com/Melany751/house-match-server/infrastructure/handler/roleview"
	routerUser "github.com/Melany751/house-match-server/infrastructure/handler/user"
	routerUserRole "github.com/Melany751/house-match-server/infrastructure/handler/userrole"
	routerView "github.com/Melany751/house-match-server/infrastructure/handler/view"
)

func InitRoutes(specification model.RouterSpecification) {
	// User
	routerUser.NewRouter(specification)
	// Person
	routerPerson.NewRouter(specification)
	// LocationPerson
	routerLocationPerson.NewRouter(specification)
	// Role
	routerRole.NewRouter(specification)
	// Module
	routerModule.NewRouter(specification)
	// View
	routerView.NewRouter(specification)
	// RoleView
	routerRoleView.NewRouter(specification)
	// UserRole
	routerUserRole.NewRouter(specification)
	// Property
	routerProperty.NewRouter(specification)
}
