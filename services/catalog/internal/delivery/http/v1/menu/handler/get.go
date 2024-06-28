package menuHandler

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func (h *MenuHandler) Get(w http.ResponseWriter, r *http.Request) {
	getMenuRe := regexp.MustCompile(`^\/menu\/(\d+)$`)
	matches := getMenuRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		return
	}
	h.store.RLock()
	u, ok := h.store.Data[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
