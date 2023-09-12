package repository

import (
	"context"
	"pos_app/payment_services/modules/models"
)

type Payment interface {
	RepoList(ctx context.Context, limit, offset uint64) (resPayment []models.Payment, err error)
	RepoFindByid(ctx context.Context, id uint64) (resPayment models.Payment, err error)
	RepoCreate(ctx context.Context, paymentIn models.Payment) (resPayment models.Payment, err error)
	RepoUpdate(ctx context.Context, paymentIn models.Payment) (resPayment models.Payment, err error)
	RepoDelete(ctx context.Context, id uint64) (err error)
}
