package main

import (
	"context"
	"time"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, accountInput AccountInput) (*Account, error) {
	account := &Account{
		Name: accountInput.Name,
	}
	err := r.server.accountClient.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, productInput ProductInput) (*Product, error) {
	product := &Product{
		Name:        productInput.Name,
		Description: productInput.Description,
		Price:       productInput.Price,
	}
	err := r.server.catalogClient.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, orderInput OrderInput) (*Order, error) {
	order := &Order{
		CreatedAt: time.Now().UTC(),
	}
	err := r.server.orderClient.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
