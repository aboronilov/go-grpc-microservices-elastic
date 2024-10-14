package main

import "context"

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	orders, err := r.server.orderClient.GetOrdersByAccountID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *accountResolver) ID(ctx context.Context, obj *Account) (string, error) {
	account, err := r.server.accountClient.GetAccountByID(ctx, obj.ID)
	if err != nil {
		return "invalid_id", err
	}
	return account, nil
}
