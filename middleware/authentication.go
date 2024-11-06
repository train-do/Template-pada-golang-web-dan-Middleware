package middleware

import (
	"database/sql"
	"net/http"

	"github.com/train-do/Template-pada-golang-web-dan-Middleware/service"
)

func Authentication(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("MASUK MIDDLEWARE")
		cookie, err := r.Cookie("access_token")
		if err != nil || err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		accessToken := cookie.Value
		serviceUser := service.ServiceUser{Db: db}
		if err = serviceUser.GetById(accessToken); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
