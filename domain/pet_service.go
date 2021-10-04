package domain

import (
	"context"
)

type PetService interface {
	AddPet(ctx context.Context, req AddPetRequest) AddPetResponse
	DeletePet(ctx context.Context, req DeletePetRequest) DeletePetResponse
	FindPetsByStatus(ctx context.Context, req FindPetsByStatusRequest) FindPetsByStatusResponse
	FindPetsByTags(ctx context.Context, req FindPetsByTagsRequest) FindPetsByTagsResponse
	GetPetByID(ctx context.Context, req GetPetByIDRequest) GetPetByIDResponse
	UpdatePet(ctx context.Context, req UpdatePetRequest) UpdatePetResponse
	UpdatePetWithForm(ctx context.Context, req UpdatePetWithFormRequest) UpdatePetWithFormResponse
	UploadFile(ctx context.Context, req UploadFileRequest) UploadFileResponse
	
}

type PetService interface {
	AddPet(ctx context.Context, body models.Pet)  models.Pet, error
	DeletePet(ctx context.Context, api_key string, petId int64) , error
	FindPetsByStatus(ctx context.Context, status []string)  []*models.Pet, error
	FindPetsByTags(ctx context.Context, tags []string)  []*models.Pet, error
	GetPetByID(ctx context.Context, petId int64)  models.Pet, error
	UpdatePet(ctx context.Context, body models.Pet)  models.Pet, error
	UpdatePetWithForm(ctx context.Context, name string, petId int64, status string)  models.Pet, error
	UploadFile(ctx context.Context, additionalMetadata string, file io.ReadCloser, petId int64)  models.APIResponse, error
	
}
