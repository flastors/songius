package apperror

import (
	"net/http"

	"github.com/flastors/songius/pkg/utils/logging"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger()
		logger.Infof("%s: %s %s", r.RemoteAddr, r.Method, r.URL)
		var err error
		if err = h(w, r); err != nil {
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte(err.Error()))
			logger.Warn(err.Error())
		}
	}
}
