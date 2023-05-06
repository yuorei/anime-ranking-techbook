package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"
	"strings"

	"github.com/yuorei/anime-ranking-techbook/graph/generated"
	"github.com/yuorei/anime-ranking-techbook/graph/model"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	splitID := strings.Split(id, ":")

	kind := splitID[0]
	id = splitID[1]
	switch kind {
	case "user":
		return r.User(ctx, id)
	case "anime":
		return r.GetAnimeRanking(ctx, id)
	default:
		return nil, fmt.Errorf("そのようなIDは定義されていません")
	}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }