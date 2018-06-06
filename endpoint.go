package aka

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
	"net/http"
)

// Handler handles incoming http requests
type Handler struct {
	service Service
}

// // ServerHTTP processes incoming http requests
// func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// handle request
// 	return http.HandleFunc
// }

func (h Handler) Create(w http.ResponseWriter, r *http.Request)  {
	id, _ := uuid.NewV4()
	company := &Company{
		uuid: id.String()
	}

	json.NewDecoder(r.Body).Decode(&company)
	err := h.service.CreateCompany(company)
	if err != nil {
		render.json(w, r, err)
		return
	}
	render.JSON(w, r, company)
}

func (h Handler) List(w http.ResponseWriter, r *http.Request)  {
	companies, err := h.service.CreateCompany(company)
	if err != nil {
		render.json(w, r, err)
		return
	}
	render.JSON(w, r, companies)
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	oid := bson.ObjectIdHex(id)
	company, err := h.service.Company(oid)
	if err != nil {
		render.json(w, r, err)
		return
	}
	render.JSON(w, r, company)
}