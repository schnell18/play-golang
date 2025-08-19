package api

import (
	"encoding/json"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
