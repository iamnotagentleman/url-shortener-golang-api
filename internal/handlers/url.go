package handlers

import (
	"awesome-url-shortener/internal/models"
	"awesome-url-shortener/internal/service"
	"encoding/json"
	"net/http"

	"github.com/ggicci/httpin"
)

type Handler struct {
	Service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func WithHTTPIn[T any](handler http.HandlerFunc, input T) http.Handler {
	return httpin.NewInput(input)(handler)
}

func (h *Handler) ShortURl(w http.ResponseWriter, req *http.Request) {
	input := req.Context().Value(httpin.Input).(*models.UrlShortCreateInput)
	val := h.Service.ShortUrl(req.Context(), *input.Payload)
	res, err := json.Marshal(val)

	if err != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *Handler) ResolveShortUrl(w http.ResponseWriter, req *http.Request) {
	key := req.PathValue("key")
	input := models.UrlShortGetInput{Key: key}
	val := h.Service.ResolveShortUrl(req.Context(), input)

	if val.Error != nil {
		internalError := http.StatusInternalServerError
		http.Error(w, val.Error.Error(), internalError)
	}

	http.Redirect(w, req, val.ShortUrl, http.StatusFound)
}
