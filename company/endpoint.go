package company

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satori/go.uuid"
)

// Handler handles incoming http requests
type Handler struct {
	companies UseCase
}

// NewHandler returns a new api/endpoint handler
func NewHandler(companies UseCase) *Handler {
	return &Handler{companies}
}

// Create endpoint stores a company in the database
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewV4()
	company := &Company{
		UUID:      UUID(id.String()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	json.NewDecoder(r.Body).Decode(&company)
	err := h.companies.Create(company)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, company)
}

// List endpoint displays all companies in the database
func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	companies, err := h.companies.FindAll()
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, companies)
}

// Get endpoint retrieves a company with given id from the database
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := UUID(chi.URLParam(r, "id"))
	company, err := h.companies.Find(id)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, company)
}

// Update endpoint edits a company in the database
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	var company *Company
	id := UUID(chi.URLParam(r, "id"))
	json.NewDecoder(r.Body).Decode(&company)
	err := h.companies.Update(id, company)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, company)
}
