package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(handler AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		start := time.Now()
		err := handler(w, r)
		if err != nil {
			log.Printf("%s %s err=%v took=%s", r.Method, r.URL.Path, err, time.Since(start))
		}
		return err
	}
}
