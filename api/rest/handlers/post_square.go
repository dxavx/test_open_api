package handlers

import (
	"net/http"
	"test-open-api/internal"
	"test-open-api/internal/api"
)

func (s *Server) PostSquare(w http.ResponseWriter, r *http.Request) {
	// Вычитываем запрос в сгенерированную структуру
	var request api.SquareRequest

	// JSON Парсинг запроса
	if err := s.readRequest(r, &request); err != nil {
		s.replayError(r.Context(), w, &internal.Error{
			Code:       internal.DecodingError,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	xSquare := *request.X * *request.X

	// Записываем ответ по запросу сгенерированной структуры
	s.writeResponse(r.Context(), w, http.StatusOK, api.SquareResponses{
		SquareX: &xSquare,
	})
}
