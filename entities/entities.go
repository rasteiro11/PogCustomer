package entities

import "github.com/rasteiro11/PogCustomer/src/user/repository"

func GetEntities() []any {
	return []any{
		&repository.User{},
	}
}
