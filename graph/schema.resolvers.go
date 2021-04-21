package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"inventory_management/graph/generated"
	"inventory_management/graph/model"
)

func (r *queryResolver) Item(ctx context.Context) (*model.Item, error) {
	var item = model.Item{
		ID:       "1",
		Sku:      "white-shirt",
		Quantity: 20,
	}

	return &item, nil
	// panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
