package services

import (
	"context"
	"pos_app/payment_services/modules/models"
)

type PaymentSrv interface {
	SrvList(ctx context.Context, limit, offset uint64) (resPayment []models.Payment, err error)
	SrvFindByid(ctx context.Context, id uint64) (resPayment models.Payment, err error)
	SrvCreate(ctx context.Context, paymentIn models.Payment) (resPayment models.Payment, err error)
	SrvUpdate(ctx context.Context, paymentIn models.Payment) (resPayment models.Payment, err error)
	SrvDelete(ctx context.Context, id uint64) (err error)
}
