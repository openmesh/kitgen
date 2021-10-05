package domain

import (
	"context"
)

type PetService interface {
	AddPet(ctx context.Context, body Pet) (Pet, error)
	DeletePet(ctx context.Context, aPIKey string, petID int64) error
	FindPetsByStatus(ctx context.Context, status []string) (ListPetsResponse, error)
	FindPetsByTags(ctx context.Context, tags []string) (Pet, error)
	GetPetByID(ctx context.Context, petID int64) (Pet, error)
	UpdatePet(ctx context.Context, body Pet) (Pet, error)
	UpdatePetWithForm(ctx context.Context, name string, petID int64, status string) (Pet, error)
	UploadFile(ctx context.Context, additionalMetadata string, file ReadCloser, petID int64) (APIResponse, error)
}
