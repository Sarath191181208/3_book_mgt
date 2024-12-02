package book_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sarath/3_book_mgt/internal/data"
	"sarath/3_book_mgt/internal/logger"

	"github.com/gorilla/mux"
)

type handler struct {
	Logger logger.Logger
	Db     *data.Models
}

func New(logger logger.Logger, db *data.Models) handler {
	return handler{
		Logger: logger,
		Db:     db,
	}
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Db.Book.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    h.Logger.Error(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *handler) RegisterBook(w http.ResponseWriter, r *http.Request) {
	book := data.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

  err = h.Db.Book.Insert(&book)
  if err != nil{
    w.WriteHeader(http.StatusInternalServerError)
    h.Logger.Error(err)
    return 
  }

	w.WriteHeader(http.StatusCreated)
  h.Logger.Info("Got Register request")
}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id_str, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
    h.Logger.Error(nil, "can't find id in params")
		return
	}
	_, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
    h.Logger.Error(err)
		return
	}

	// parsing the book data from the body
	book := data.Book{}
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// reading the previous book
	id_num := json.Number(params["id"])
	prev_book, err := h.Db.Book.GetById(id_num)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
    h.Logger.Error(err)
		return
	}

	// updating the missing fields
	if book.Author == "" {
		book.Author = prev_book.Author
	}
	if book.Name == "" {
		book.Name = prev_book.Name
	}
  if book.Id == ""{
    book.Id = prev_book.Id
  }

	// performing the db update
	err = h.Db.Book.Update(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    h.Logger.Error(err)
		return
	}

	// returning the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id_str, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	_, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	h.Db.Book.Delete(json.Number(params["id"]))
}
