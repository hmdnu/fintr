package middleware

import (
	"log/slog"
	"net/http"

	"github.com/hmdnu/fintr/pkg/logger"
)

func Logger(handler AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := handler(w, r)
		if err != nil {
			slog.Error(err.Error())
			logger.AppLogger.ErrorLogger(err.Error())
			return err
		}
		return err
	}
}
