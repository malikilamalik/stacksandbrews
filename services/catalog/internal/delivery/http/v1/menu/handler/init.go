package menuHandler

import "github.com/malikilamalik/stacksandbrews/services/catalog/internal/repository/model"

type MenuHandler struct {
	store model.MenuData
}

func Init(store model.MenuData) MenuHandler {
	return MenuHandler{store: store}
}
