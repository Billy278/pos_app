package controllers

import (
	"net/http"
	"pos_app/payment_services/modules/services"
	"pos_app/payment_services/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CtrlPaymentImpl struct {
	PaymentSrv services.PaymentSrv
}

func (ctrl *CtrlPaymentImpl) List(ctx *gin.Context) {
	limit, offset, err := ctrl.GetQuery(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidQuery,
		})
		return
	}
	res, err := ctrl.PaymentSrv.SrvList(ctx, limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: responses.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, responses.Response{
		Code:    http.StatusAccepted,
		Success: true,
		Message: responses.Success,
		Data:    res,
	})

}
func (ctrl *CtrlPaymentImpl) FindByid(ctx *gin.Context) {}
func (ctrl *CtrlPaymentImpl) Created(ctx *gin.Context)
func (ctrl *CtrlPaymentImpl) Updated(ctx *gin.Context)
func (ctrl *CtrlPaymentImpl) Deleted(ctx *gin.Context)
func (ctrl *CtrlPaymentImpl) GetQuery(ctx *gin.Context) (limit, offset uint64, err error) {
	limitQry, ok := ctx.GetQuery("limit")
	if !ok {
		limitQry = "10"
	}
	offsetQry, ok := ctx.GetQuery("offset")
	if !ok {
		offsetQry = "1"
	}
	limit, err = strconv.ParseUint(limitQry, 10, 64)
	if err != nil {
		return
	}

	offset_, err := strconv.ParseUint(offsetQry, 10, 64)
	if err != nil {
		return
	}
	if offset_ > 1 {
		offset_ *= 10
	}
	offset = offset_

	return
}
