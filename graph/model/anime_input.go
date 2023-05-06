package model

import "github.com/99designs/gqlgen/graphql"

type NewAnimeRankingInput struct {
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	Rank        int            `json:"rank"`
	AnimeImage  graphql.Upload `json:"animeImage"`
}

type UpdateAnimeRankingInput struct {
	Title       *string         `json:"title,omitempty"`
	Description *string         `json:"description,omitempty"`
	Rank        *int            `json:"rank,omitempty"`
	AnimeImage  *graphql.Upload `json:"animeImage,omitempty"`
}
