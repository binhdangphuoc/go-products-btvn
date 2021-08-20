package action

import (
	"encoding/json"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)
var tokenAdmin = "admin"
func AminCheckingMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		if reqToken!=tokenAdmin{
			logrus.Info(" Request not login admin")
			json.NewEncoder(w).Encode(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w,r)
	})
}

