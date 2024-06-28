package menuRoute

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"

	menuHandler "github.com/malikilamalik/stacksandbrews/services/catalog/internal/delivery/http/v1/menu/handler"
	"github.com/malikilamalik/stacksandbrews/services/catalog/internal/repository/model"
)

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

type menuRoute struct {
	menuHandler menuHandler.MenuHandler
}

func Init() menuRoute {
	menuData := model.MenuData{
		Data: map[string]model.Menu{
			"1": {
				ID:   "OKE",
				Name: "Anice Baswedan",
			},
			"2": {
				ID:   "Nat Nat",
				Name: "Blooming Nat Nat",
			},
		},
		RWMutex: &sync.RWMutex{},
	}
	menu := menuHandler.Init(menuData)
	return menuRoute{menuHandler: menu}
}

var (
	listMenuRe   = regexp.MustCompile(`^\/menu[\/]*$`)
	getMenuRe    = regexp.MustCompile(`^\/menu\/(\d+)$`)
	createUserRe = regexp.MustCompile(`^\/menu[\/]*$`)
)

func (m menuRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(getMenuRe.MatchString(r.URL.Path))
	switch {
	//Get List
	case r.Method == http.MethodGet && listMenuRe.MatchString(r.URL.Path):
		m.menuHandler.GetList(w, r)
		return
	case r.Method == http.MethodGet && getMenuRe.MatchString(r.URL.Path):
		m.menuHandler.Get(w, r)
		return
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		m.menuHandler.Get(w, r)
		return
	default:
		notFound(w)
		return
	}
}
