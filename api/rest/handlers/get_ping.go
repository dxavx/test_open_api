package handlers

import (
	"net/http"
	"test-open-api/internal/api"
	"time"
)

func (s *Server) GetPing(w http.ResponseWriter, r *http.Request) {
	ok := "sleep 1 sec"
	// Просто пишем в ответ сервера сгенерированную структур api.Ping
	time.Sleep(time.Second * 1)
	s.writeResponse(r.Context(), w, http.StatusOK, api.Ping{Status: &ok})
}
