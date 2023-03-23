package graph

import "knowledge/api/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserStore map[string]model.User
	TodoStore map[string]model.Todo
}
