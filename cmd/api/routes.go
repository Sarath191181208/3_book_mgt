package api

import (
	"net/http"

	book_handler "sarath/3_book_mgt/cmd/services"

	"github.com/gorilla/mux"
)

func (app *Application) Routes() *mux.Router {
	router := mux.NewRouter()
	api_router := router.PathPrefix("/api/v1").Subrouter()

	books_handler := book_handler.New(app.Logger, app.Db)

	api_router.HandleFunc("/books", books_handler.GetAllBooks).Methods(http.MethodGet)
	api_router.HandleFunc("/books", books_handler.RegisterBook).Methods(http.MethodPost)
	api_router.HandleFunc("/books/{id}", books_handler.UpdateBook).Methods(http.MethodPatch)
	api_router.HandleFunc("/books/{id}", books_handler.DeleteBook).Methods(http.MethodDelete)

	return router
}
