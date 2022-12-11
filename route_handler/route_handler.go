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
	book := models.Book{}
	rows, err := h.database.Queryx("SELECT * FROM books")
	if err != nil {
		h.logger.Fatalf("Failed to query Select: %v", err)
	}
	for rows.Next() {
		err := rows.StructScan(&book)
		if err != nil {
			h.logger.Fatalf("Failed to get select result: %v", err)
		}
		resp.Books = append(resp.Books, book)
	}
	json.NewEncoder(w).Encode(resp)
}
