package model

import "github.com/99designs/gqlgen/graphql"

type UpdateUserInput struct {
	Name         *string         `json:"name,omitempty"`
	Description  *string         `json:"description,omitempty"`
	ProfileImage *graphql.Upload `json:"profileImage,omitempty"`
}

type UserInformationInput struct {
	Name         string         `json:"name"`
	Password     string         `json:"password"`
	Description  *string        `json:"description,omitempty"`
	ProfileImage graphql.Upload `json:"profileImage"`
}
