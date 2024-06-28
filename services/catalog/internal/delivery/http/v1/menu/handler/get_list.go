package menuHandler

import (
	"encoding/json"
	"net/http"

	"github.com/malikilamalik/stacksandbrews/services/catalog/internal/repository/model"
)

func (h *MenuHandler) GetList(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	users := make([]model.Menu, 0, len(h.store.Data))
	for _, v := range h.store.Data {
		users = append(users, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
