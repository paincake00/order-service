package errors

import (
	"net/http"

	"github.com/paincake00/order-service/internal/jsonutil"
	"go.uber.org/zap"
)

type WebErrorService struct {
	Logger *zap.SugaredLogger
}

func (wes *WebErrorService) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	wes.Logger.Errorf(
		"internal error: method=%s, path=%s, err=%s",
		r.Method,
		r.URL.Path,
		err.Error(),
	)

	_ = jsonutil.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
}

func (wes *WebErrorService) NotFoundResponseError(w http.ResponseWriter, r *http.Request, err error) {
	wes.Logger.Warnf(
		"not found error: method=%s, path=%s, err=%s",
		r.Method,
		r.URL.Path,
		err.Error(),
	)

	_ = jsonutil.WriteJSONError(w, http.StatusNotFound, "not found")
}
