package company

import "github.com/go-chi/chi"

// Routes creates a REST router
func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	// add middleware specific to user Routes
	r.Get("/", h.List)    // GET /companies - read list of company
	r.Post("/", h.Create) // POST /company - create new company
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get) // GET /company/{id} - read a single company
		// r.Put("/", h.Update)    // PUT /company/{id} - update a single company
		// r.Delete("/", h.Delete) // DELETE /company/{id} - delete a single company
	})

	return r
}
