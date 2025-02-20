package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var backendURL = "http://localhost:8081"

// NewProxy cria um handler de proxy reverso
func NewProxy() http.Handler {
	target, err := url.Parse(backendURL)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Middleware pra customizar o comportamento do proxy
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host 
		proxy.ServeHTTP(w, r)
	})
}
