package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paincake00/order-service/internal/db"
	"github.com/paincake00/order-service/internal/domain/service"
	errorsUtil "github.com/paincake00/order-service/internal/errors"
	"github.com/paincake00/order-service/internal/jsonutil"
	"go.uber.org/zap"
)

type OrderHandler struct {
	OrderService service.OrderService
	ErrorService *errorsUtil.WebErrorService
	Logger       *zap.SugaredLogger
}

func (o *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	ctx := r.Context()

	order, err := o.OrderService.GetOrderByID(ctx, idParam)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrNotFound):
			o.ErrorService.NotFoundResponseError(w, r, err)
		default:
			o.ErrorService.InternalServerError(w, r, err)
		}
		return
	}

	if err = jsonutil.WriteJSONResponse(w, http.StatusOK, order); err != nil {
		o.ErrorService.InternalServerError(w, r, err)
		return
	}
}
