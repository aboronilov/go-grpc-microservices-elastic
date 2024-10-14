package main

import "context"

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	accounts, err := r.server.accountClient.GetAllAccounts(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query, id *string) ([]*Product, error) {
	products, err := r.server.catalogClient.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *queryResolver) Orders(ctx context.Context, pagination *PaginationInput, query, id *string) ([]*Order, error) {
	orders, err := r.server.orderClient.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
