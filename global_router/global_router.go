package global_router

import "net/http"

func NewGlobalRouter(mux *http.ServeMux) http.Handler {
	handlerAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
		} else {
			mux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(handlerAllReq)
}
