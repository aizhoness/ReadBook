package route_handler

import (
	"ReadBook/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Handler struct for HTTP requests
type Handler struct {
	logger   *zap.SugaredLogger
	router   *mux.Router
	database *sqlx.DB
}

// New creates a Handler struct
func New(logger *zap.SugaredLogger, router *mux.Router, database *sqlx.DB) *Handler {
	h := Handler{logger, router, database}
	h.registerRoutes()

	return &h
}

// RegisterRoutes registers all the routes for the route handler
func (h *Handler) registerRoutes() {
	h.router.HandleFunc("/books", h.getBooks).Methods("GET")
}

// getBeans is the route handler for the GET /beans endpoint
func (h *Handler) getBooks(w http.ResponseWriter, r *http.Request) {
	var resp = &models.BooksResp{}
	var b = models.Book{
		ID:       1,
		Title:    "Head First Go",
		AuthorID: 1,
	}
	resp.Books = append(resp.Books, b)

	json.NewEncoder(w).Encode(resp)
}
