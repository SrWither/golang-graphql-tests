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
		return nil, err
	}

	token, err := JwtGenerate(ctx, strconv.Itoa(createdUser.ID))
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}

func UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := GetUserEmail(ctx, email)
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "User not found",
		})
	}

	if err := tools.ComparePassword(getUser.Password, password); err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, strconv.Itoa(getUser.ID))
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user":  getUser.Name,
		"email": getUser.Email,
		"token": token,
	}, nil
}
