package route_handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Handler struct for HTTP requests
type Handler struct {
	logger *zap.SugaredLogger
	router *mux.Router
}

// New creates a Handler struct
func New(logger *zap.SugaredLogger, router *mux.Router) *Handler {
	h := Handler{logger, router}
	h.registerRoutes()

	return &h
}

// RegisterRoutes registers all the routes for the route handler
func (h *Handler) registerRoutes() {
	h.router.HandleFunc("/books", h.getBooks).Methods("GET")
}

// Book represents product in our API
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BooksResp is the response for the GET /books endpoint
type BooksResp struct {
	Books []Book `json:"books"`
}

// getBeans is the route handler for the GET /beans endpoint
func (h *Handler) getBooks(w http.ResponseWriter, r *http.Request) {
	var books = make([]Book, 0, 1)
	var b = Book{
		Title:  "Head First Go",
		Author: "McGavren",
	}
	books = append(books, b)

	json.NewEncoder(w).Encode(&BooksResp{books})
}
