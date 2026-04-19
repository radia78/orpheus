package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	api "github.com/radia78/orpheus/src/internal/models"

	log "github.com/sirupsen/logrus"
)

type CollectionHandler struct {
	vectorStore *milvusclient.Client
}

func NewCollectionHandler(s *milvusclient.Client) *CollectionHandler {
	return &CollectionHandler{vectorStore: s}
}

func (c *CollectionHandler) CreateCollection(w http.ResponseWriter, r *http.Request) {
	var params = api.CreateDeleteCollectionRequest{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	ctx := r.Context()

	// Create collection with the specified schema
	// Note: Ideally a duplicate is if the song has the same artist-album-song_name combination but we'll do that later

	schema := entity.NewSchema()
	// Primary key is a unique ID
	schema.WithField(
		entity.NewField().WithName("id").
			WithDataType(entity.FieldTypeInt64).
			WithIsPrimaryKey(true).
			WithIsAutoID(false),
	)

	// Add string fields
	schema.WithField(
		entity.NewField().WithName("name").
			WithDataType(entity.FieldTypeVarChar).
			WithMaxLength(512),
	)
	schema.WithField(
		entity.NewField().WithName("album").
			WithDataType(entity.FieldTypeVarChar).
			WithMaxLength(512),
	)
	schema.WithField(
		entity.NewField().WithName("artist").
			WithDataType(entity.FieldTypeVarChar).
			WithMaxLength(512),
	)

	// Add vector fields
	schema.WithField(
		entity.NewField().WithName("mood").
			WithDataType(entity.FieldTypeFloatVector).
			WithDim(128), // The mood should be an image vector embedding
	)

	err = c.vectorStore.CreateCollection(ctx, milvusclient.NewCreateCollectionOption(
		params.Name,
		schema,
	))

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CreateDeleteCollectionResponse{
		Code:             http.StatusOK,
		Name:             params.Name,
		CreatedDeletedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response) // Set a new encoder using the writer and then encode the response
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func (c *CollectionHandler) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	var params = api.CreateDeleteCollectionRequest{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	ctx := r.Context()

	// Create collection with the specified schema
	// Note: Ideally a duplicate is if the song has the same artist-album-song_name combination but we'll do that later

	err = c.vectorStore.DropCollection(ctx, milvusclient.NewDropCollectionOption(params.Name))

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CreateDeleteCollectionResponse{
		Code:             http.StatusOK,
		Name:             params.Name,
		CreatedDeletedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response) // Set a new encoder using the writer and then encode the response
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func (c *CollectionHandler) RenameCollection(w http.ResponseWriter, r *http.Request) {
	// Still don't know the old name
	var params = api.RenameCollectionRequest{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	ctx := r.Context()

	// Create collection with the specified schema
	// Note: Ideally a duplicate is if the song has the same artist-album-song_name combination but we'll do that later

	err = c.vectorStore.RenameCollection(ctx, milvusclient.NewRenameCollectionOption(
		params.Name,
		params.NewName,
	))

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.RenameCollectionResponse{
		Code:      http.StatusOK,
		NewName:   params.NewName,
		UpdatedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response) // Set a new encoder using the writer and then encode the response
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
