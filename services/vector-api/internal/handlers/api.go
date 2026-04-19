package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	pb "github.com/radia78/orpheus/services/vector-api/internal/embedding"
)

func Handler(r *chi.Mux, mc *milvusclient.Client, oc *pb.EmbedderClient) {
	// Text cleaning
	r.Use(chimiddle.StripSlashes)

	// Home page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Orpheus API. Please see the docs."))
	})

	r.Route("/collections", func(router chi.Router) {
		ch := NewCollectionHandler(mc)
		eh := NewEntriesHandler(mc, oc)

		router.Post("/", ch.CreateCollection)

		router.Route("/{collection_name}", func(router chi.Router) {
			router.Delete("/", ch.DeleteCollection)
			router.Patch("/", ch.RenameCollection)

			r.Route("/songs", func(router chi.Router) {
				r.Post("/", eh.AddNewEntry)
				r.Get("/", eh.QueryEntries)
				r.Patch("/{id}", eh.UpdateEntry)
				r.Delete("/{id}", eh.DeleteEntry)
			})
		})
	})

}
