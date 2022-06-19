package handler

import (
	"bytes"
	"net/http"

	update "github.com/sudak-91/telegrambotgo/Service"
)

type Handler struct {
	Updater update.Updater
}

func NewHandler(Upd update.Updater) *Handler {
	return &Handler{Updater: Upd}
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var b []byte
	buffer := bytes.NewBuffer(b)
	buffer.ReadFrom(r.Body)
	k, err := h.Updater.Update(buffer.Bytes())
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json")
	_, err = w.Write(k)
	if err != nil {
		panic(err.Error())
	}
}
