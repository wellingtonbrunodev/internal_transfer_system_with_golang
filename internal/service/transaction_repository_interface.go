package service

import "context"

type TransactionRepositoryInterface interface {
	Transfer(
		ctx context.Context,
		sourceID int64,
		destinationID int64,
		amount string,
	) error
}
