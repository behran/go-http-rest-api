package apiserver

import (
	"net/http"

	"github.com/behran/go-http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// newServer функция конструктор
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.HandleUsersCreate()).Methods(http.MethodPost)
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
