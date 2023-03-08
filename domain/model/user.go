package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	PersonID uuid.UUID `json:"person_id"`
}

type Users []User

type UserOutput struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	Person   Person    `json:"person"`
}

type UsersOutput []UserOutput

type UserWithRolesOutput struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	PersonID uuid.UUID `json:"person_id"`
	Roles    Roles     `json:"roles"`
}

type UsersWithRolesOutput []UserWithRolesOutput

type UserWithRole struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	PersonID uuid.UUID `json:"person_id"`
	Role     Role      `json:"role"`
}

type UsersWithRoles []UserWithRole

func (ur UsersWithRoles) GetUserWithRole() UsersWithRolesOutput {
	var results UsersWithRolesOutput
	mapa := make(map[uuid.UUID]bool)
	mapa_index := make(map[uuid.UUID]int)

	for i, element := range ur {
		if _, ok := mapa[element.ID]; !ok {
			mapa[element.ID] = true
			mapa_index[element.ID] = i
			var RolesI Roles
			results = append(results, UserWithRolesOutput{
				ID:       element.ID,
				User:     element.User,
				Password: element.Password,
				Email:    element.Email,
				Theme:    element.Theme,
				PersonID: element.PersonID,
				Roles: append(RolesI, Role{
					ID:          element.Role.ID,
					Name:        element.Role.Name,
					Description: element.Role.Description,
					Order:       element.Role.Order,
				}),
			})
		} else {
			results[mapa_index[element.ID]].Roles = append(results[mapa_index[element.ID]].Roles, Role{
				ID:          element.Role.ID,
				Name:        element.Role.Name,
				Description: element.Role.Description,
				Order:       element.Role.Order,
			})
			//results = BusquedaUsuario(results, element.ID, element.Role)
		}
	}

	return results
}

func BusquedaUsuario(ur UsersWithRolesOutput, idUser uuid.UUID, r Role) (results UsersWithRolesOutput) {
	for i, element := range ur {
		if element.ID == idUser {
			element.Roles = append(element.Roles, Role{
				ID:          r.ID,
				Name:        r.Name,
				Description: r.Description,
				Order:       r.Order,
			})
		}
		ur[i] = element
	}
	return ur
}
