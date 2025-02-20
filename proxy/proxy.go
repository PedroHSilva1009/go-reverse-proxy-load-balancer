package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Algorithm string `json:"algorithm"`
}

func loadConfig() (*Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(file, &config)
	return &config, err
}

func NewProxy() (http.Handler, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar configuração: %v", err)
	}

	switch config.Algorithm {
	case "roundrobin":
		return NewRoundRobinProxy(), nil
	case "leastconn":
		return NewLeastConnectionsProxy(), nil
	default:
		return nil, fmt.Errorf("algoritmo de balanceamento inválido: %s", config.Algorithm)
	}
}
