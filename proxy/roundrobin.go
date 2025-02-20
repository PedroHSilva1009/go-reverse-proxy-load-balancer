package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)


var backendServers = []string{
	"http://localhost:8081",
	"http://localhost:8082",
	"http://localhost:8083",
}

var currentIndex uint64

func getNextBackend() string {
	index := atomic.AddUint64(&currentIndex, 1) % uint64(len(backendServers))
	return backendServers[index]
}

func NewRoundRobinProxy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backend := getNextBackend()

		target, err := url.Parse(backend)
		if err != nil {
			http.Error(w, "Erro ao escolher backend", http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		log.Printf("Encaminhando requisição para %s", backend)
		go proxy.ServeHTTP(w, r)
	})
}
