package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"fmt"
	"knowledge/api/database"
	graph "knowledge/api/graph/model"
	"knowledge/api/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input graph.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CreateUser is the resolver for the createUser field
func (r *mutationResolver) CreateUser(ctx context.Context, input graph.NewUser) (*model.User, error) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := database.ConnectDatabase(connectionString)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	user, err := database.InsertUser(db, input.Username)

	return user, err
}

// ID is the resolver for the id field.
func (r *projectResolver) ID(ctx context.Context, obj *model.Project) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := database.ConnectDatabase(connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	users, err := database.GetUsers(db)
	return users, err
}

// ID is the resolver for the id field.
func (r *todoResolver) ID(ctx context.Context, obj *model.Todo) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Project returns ProjectResolver implementation.
func (r *Resolver) Project() ProjectResolver { return &projectResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
