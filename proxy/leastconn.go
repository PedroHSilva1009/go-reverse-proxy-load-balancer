package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL         *url.URL
	ActiveConns int
	mu          sync.Mutex
}

var backendsLeastConn = []*Backend{
	{URL: mustParse("http://localhost:8081")},
	{URL: mustParse("http://localhost:8082")},
	{URL: mustParse("http://localhost:8083")},
}

func getLeastConnectionsBackend() *Backend {
	var best *Backend
	for _, backend := range backendsLeastConn {
		backend.mu.Lock()
		if best == nil || backend.ActiveConns < best.ActiveConns {
			best = backend
		}
		backend.mu.Unlock()
	}
	return best
}

func NewLeastConnectionsProxy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backend := getLeastConnectionsBackend()
		backend.mu.Lock()
		backend.ActiveConns++
		backend.mu.Unlock()

		defer func() {
			backend.mu.Lock()
			backend.ActiveConns--
			backend.mu.Unlock()
		}()

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)
		r.Host = backend.URL.Host
		proxy.ServeHTTP(w, r)
	})
}

func mustParse(rawURL string) *url.URL {
	parsed, _ := url.Parse(rawURL)
	return parsed
}
