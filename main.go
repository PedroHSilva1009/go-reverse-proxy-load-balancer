package main

import (
	"log"
	"net/http"
	"proxy-reverso-go/proxy"
)

func main() {
	log.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", proxy.NewProxy())
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}