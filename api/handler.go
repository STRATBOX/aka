package api

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/STRATBOX/aka"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satori/go.uuid"
)

// Handler handles incoming http requests
type Handler struct {
	service aka.Service
}

// Company represents a company document in mongo
type Company struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	UUID string        `json:"uuid" bson:"uuid"`
	*aka.Company
}

// NewHandler returns a new api/endpoint handler
func NewHandler(service aka.Service) *Handler {
	return &Handler{service}
}

// Create endpoint stores a company in the database
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.NewV4()
	company := &Company{
		ID:        bson.NewObjectId(),
		UUID:      id.String(),
		CreatedAt: time.Now(),
	}

	json.NewDecoder(r.Body).Decode(&company)
	err := h.service.Create(company)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, company)
}

// List endpoint displays all companies in the database
func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	companies, err := h.service.FindAll()
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, companies)
}

// Get endpoint retrieves a company with given id from the database
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	company, err := h.service.Find(id)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, company)
}
