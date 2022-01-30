package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-graphql-test/src/types"
)

func (r *queryResolver) GetUser(ctx context.Context) (*string, error) {
	temp := "hello world"
	return &temp, nil
}

func (r *queryResolver) SearchUser(ctx context.Context, payload *types.SearchUserInput) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetBook(ctx context.Context, id *string) (*types.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
