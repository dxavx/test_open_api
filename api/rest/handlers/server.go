package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"test-open-api/internal"
	"test-open-api/internal/api"
)

func NewServer() api.ServerInterface {
	return &Server{}
}

type Server struct {
}

type errorResponse struct {
	Code    internal.ErrorCode `json:"code"`
	Message string             `json:"message,omitempty"`
}

// readRequest
func (s *Server) readRequest(r *http.Request, model interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		return fmt.Errorf("json.Decode: %w", err)
	}
	return nil
}

// writeResponse
func (s *Server) writeResponse(
	ctx context.Context,
	w http.ResponseWriter,
	statusCode int,
	model interface{},
) {
	data, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if _, err = w.Write(data); err != nil {
		fmt.Printf("w.Write: %+v", err)
	}
}

// replayError
func (s *Server) replayError(ctx context.Context, w http.ResponseWriter, err error) {

	var internalError *internal.Error

	if errors.As(err, &internalError) {
		statusCode := internalError.StatusCode
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}

		s.writeResponse(ctx, w, statusCode, errorResponse{
			Code:    internal.ErrorUnknown,
			Message: "Unknown error",
		})
	}
}
