package service

import (
	"context"
	"pruebas/graph/model"
	"pruebas/prisma/db"
	"pruebas/tools"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	// User res
	var User *model.User = nil

	// Connect to database
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "unable to connect to database",
		})
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	pctx := context.Background()

	// Hash Password
	input.Password = tools.HashPassword(input.Password)

	// Create user
	createUser, err := client.User.CreateOne(
		db.User.Name.Set(input.Name),
		db.User.Email.Set(input.Email),
		db.User.Password.Set(input.Password),
	).Exec(pctx)
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Error to create an user",
		})
	}

	// Append data to res
	if createUser != nil {
		User = &model.User{
			ID:       createUser.ID,
			Name:     createUser.Name,
			Email:    createUser.Email,
			Password: createUser.Password,
		}
	}

	return User, nil
}

func GetUserID(ctx context.Context, id int) (*model.User, error) {
	// User res
	var User *model.User = nil

	// Connect to database
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "unable to connect to database",
		})
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	pctx := context.Background()

	user, err := client.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(pctx)
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "this user does not exist",
		})
	}

	// Append data to res
	if user != nil {
		User = &model.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	return User, nil
}

func GetUserEmail(ctx context.Context, email string) (*model.User, error) {
	// User res
	var User *model.User = nil

	// Connect to database
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "unable to connect to database",
		})
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	pctx := context.Background()

	user, err := client.User.FindFirst(
		db.User.Email.Equals(email),
	).Exec(pctx)
	if err != nil {
		print("Email not found\n")
	}

	// Append data to res
	if user != nil {
		User = &model.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	return User, nil
}
