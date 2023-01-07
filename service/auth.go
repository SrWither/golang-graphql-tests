package service

import (
	"context"
	"pruebas/graph/model"
	"pruebas/tools"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func UserRegister(ctx context.Context, input model.NewUser) (interface{}, error) {
	// Check Email
	Gml, err := GetUserEmail(ctx, input.Email)
	if Gml != nil {
		print("Hay un email!\n")
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Email has been registered",
		})
		return nil, nil
	}

	createdUser, err := UserCreate(ctx, input)
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Unable to create user",
		})
		return nil, nil
	}

	token, err := JwtGenerate(ctx, strconv.Itoa(createdUser.ID))
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Error to generate token",
		})
		return nil, nil
	}

	return map[string]interface{}{
		"user":  createdUser.Name,
		"email": createdUser.Email,
		"token": token,
	}, nil
}

func UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := GetUserEmail(ctx, email)
	if getUser == nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Wrong email",
		})
		return nil, nil
	}

	if err := tools.ComparePassword(getUser.Password, password); err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Wrong password",
		})
		return nil, nil
	}

	token, err := JwtGenerate(ctx, strconv.Itoa(getUser.ID))
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "Error to generate token",
		})
		return nil, nil
	}

	return map[string]interface{}{
		"user":  getUser.Name,
		"email": getUser.Email,
		"token": token,
	}, nil
}
