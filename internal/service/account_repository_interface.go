package service

import "context"

type AccountRepositoryInterface interface {
	Create(ctx context.Context, id int64, balance string) error 
	GetByID(ctx context.Context, id int64) (int64, string, error)
}
