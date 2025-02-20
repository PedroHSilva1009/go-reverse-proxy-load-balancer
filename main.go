package main

import (
	"log"
	"net/http"
	"proxy-reverso-go/proxy"
)

func main() {
	proxyHandler, err := proxy.NewProxy()
	if err != nil {
		log.Fatal("Erro ao iniciar o proxy:", err)
	}

	log.Println("Servidor proxy rodando na porta 8080...")
	err = http.ListenAndServe(":8080", proxyHandler)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor proxy:", err)
	}
}
