package services

import (
	"context"
	"fmt"
	"pos_app/payment_services/modules/models"
	"pos_app/payment_services/modules/repository"
	"time"
)

type PaymentSrvImpl struct {
	PaymentRepo repository.Payment
}

func NewPaymentSrvImpl(paymentrepo repository.Payment) PaymentSrv {
	return &PaymentSrvImpl{
		PaymentRepo: paymentrepo,
	}
}
func (serv *PaymentSrvImpl) SrvList(ctx context.Context, limit, offset uint64) (resPayment []models.Payment, err error) {
	fmt.Println("SrvList")
	resPayment, err = serv.PaymentRepo.RepoList(ctx, limit, offset)
	if err != nil {
		return
	}
	return
}
func (serv *PaymentSrvImpl) SrvFindByid(ctx context.Context, id uint64) (resPayment models.Payment, err error) {
	fmt.Println("SrvFindByid")
	resPayment, err = serv.PaymentRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	return
}
func (serv *PaymentSrvImpl) SrvCreate(ctx context.Context, paymentIn models.Payment) (resPayment models.Payment, err error) {
	fmt.Println("SrvCreate")
	tNow := time.Now()
	paymentIn.Created_At = &tNow
	resPayment, err = serv.PaymentRepo.RepoCreate(ctx, paymentIn)
	if err != nil {
		return
	}
	return
}
func (serv *PaymentSrvImpl) SrvUpdate(ctx context.Context, paymentIn models.Payment) (resPayment models.Payment, err error) {
	fmt.Println("SrvUpdate")
	_, err = serv.PaymentRepo.RepoFindByid(ctx, paymentIn.Id)
	if err != nil {
		return
	}
	tNow := time.Now()
	paymentIn.Updated_At = &tNow
	resPayment, err = serv.PaymentRepo.RepoUpdate(ctx, paymentIn)
	if err != nil {
		return
	}
	return
}
func (serv *PaymentSrvImpl) SrvDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("SrvDelete")
	_, err = serv.PaymentRepo.RepoFindByid(ctx, id)
	if err != nil {
		return
	}
	err = serv.PaymentRepo.RepoDelete(ctx, id)
	if err != nil {
		return
	}
	return
}
