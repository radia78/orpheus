package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/milvus-io/milvus/client/v2/column"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	ollama "github.com/ollama/ollama/api"
	api "github.com/radia78/orpheus/src/internal/models"
	log "github.com/sirupsen/logrus"
)

type EntryHandler struct {
	vectorStore     *milvusclient.Client
	embeddingEngine *ollama.Client
}

func NewEntryHandler(mc *milvusclient.Client, oc *ollama.Client) *EntryHandler {
	return &EntryHandler{
		vectorStore:     mc,
		embeddingEngine: oc,
	}
}

func (eh *EntryHandler) AddNewEntry(w http.ResponseWriter, r *http.Request) {
	// Populate the request
	var entry = api.AddNewEntryRequest{}
	var collectionName string

	collectionName = chi.URLParam(r, "collection_name")

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	ctx := r.Context()

	// Pass the list of strings through ollama and get embeddings back
	req := &ollama.EmbedRequest{
		Model: "dummy-model",
		Input: entry.Mood,
	}
	resp, err := eh.embeddingEngine.Embed(ctx, req)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Buffering the metadata inserts
	nameColumn := column.NewColumnString("name", entry.Name)
	albumColumn := column.NewColumnString("album", entry.Album)
	artistColumn := column.NewColumnString("artist", entry.Artist)
	moodColumn := column.NewColumnString("mood", entry.Mood)

	// Mass insert but users are allowed to only insert one song?
	// Might change this so that users are allowed to buffer a lot of songs in one place
	// Before mass send but who knows...
	_, err = eh.vectorStore.Insert(ctx,
		milvusclient.NewColumnBasedInsertOption(collectionName).
			WithInt64Column("id", entry.ID).
			WithFloatVectorColumn("mood", 128, resp.Embeddings).
			WithColumns(nameColumn, albumColumn, artistColumn, moodColumn),
	)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

func (eh *EntryHandler) QueryEntry(w http.ResponseWriter, r *http.Request) {
	// TODO: To be implemented
}

func (eh *EntryHandler) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	var entry = api.AddNewEntryRequest{}
	var collectionName string

	collectionName = chi.URLParam(r, "collection_name")

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	ctx := r.Context()

	// Pass the list of strings through ollama and get embeddings back
	req := &ollama.EmbedRequest{
		Model: "dummy-model",
		Input: entry.Mood,
	}
	resp, err := eh.embeddingEngine.Embed(ctx, req)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Buffering the metadata updates
	nameColumn := column.NewColumnString("name", entry.Name)
	albumColumn := column.NewColumnString("album", entry.Album)
	artistColumn := column.NewColumnString("artist", entry.Artist)
	moodColumn := column.NewColumnString("mood", entry.Mood)

	// Pass the moods through the embedding service to yield the mood embedding vectors

	// Mass insert but users are allowed to only insert one song?
	// Might change this so that users are allowed to buffer a lot of songs in one place
	// Before mass send but who knows...
	_, err = eh.vectorStore.Upsert(ctx,
		milvusclient.NewColumnBasedInsertOption(collectionName).
			WithInt64Column("id", entry.ID).
			WithFloatVectorColumn("mood", 128, resp.Embeddings).
			WithColumns(nameColumn, albumColumn, artistColumn, moodColumn),
	)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func (eh *EntryHandler) DeleteEntry(w http.ResponseWriter, r *http.Request) {
	var entry = api.DeleteEntryRequest{}
	var collectionName string

	collectionName = chi.URLParam(r, "collection_name")

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	ctx := r.Context()

	// Mass insert but users are allowed to only insert one song?
	// Might change this so that users are allowed to buffer a lot of songs in one place
	// Before mass send but who knows...
	_, err = eh.vectorStore.Delete(ctx,
		milvusclient.NewDeleteOption(collectionName).
			WithInt64IDs("id", entry.ID),
	)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
