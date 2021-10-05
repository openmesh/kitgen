package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func PetLoggingMiddleware(logger log.Logger) domain.PetServiceMiddleware {
	return func(next domain.PetService) domain.PetService {
		return petLoggingMiddleware{logger, next}
	}
}

type petLoggingMiddleware struct {
	logger log.Logger
	next   domain.PetService
}

func (mw *addPetLoggingMiddleware) AddPet(ctx context.Context, body Pet) (addPetOK domain.Pet, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "add_pet",
			"body", body,
		)
	}()
}

func (mw *deletePetLoggingMiddleware) DeletePet(ctx context.Context, aPIKey string, petID int64) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "delete_pet",
			"apiKey", apiKey,
			"petId", petId,
		)
	}()
}

func (mw *findPetsByStatusLoggingMiddleware) FindPetsByStatus(ctx context.Context, status []string) (findPetsByStatusOK domain.ListPetsResponse, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "find_pets_by_status",
			"status", status,
		)
	}()
}

func (mw *findPetsByTagsLoggingMiddleware) FindPetsByTags(ctx context.Context, tags []string) (findPetsByTagsOK []*domain.Pet, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "find_pets_by_tags",
			"tags", tags,
		)
	}()
}

func (mw *getPetByIdLoggingMiddleware) GetPetByID(ctx context.Context, petID int64) (getPetByIdOK domain.Pet, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "get_pet_by_id",
			"petId", petId,
		)
	}()
}

func (mw *updatePetLoggingMiddleware) UpdatePet(ctx context.Context, body Pet) (updatePetOK domain.Pet, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "update_pet",
			"body", body,
		)
	}()
}

func (mw *updatePetWithFormLoggingMiddleware) UpdatePetWithForm(ctx context.Context, name string, petID int64, status string) (updatePetWithFormOK domain.Pet, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "update_pet_with_form",
			"name", name,
			"petId", petId,
			"status", status,
		)
	}()
}

func (mw *uploadFileLoggingMiddleware) UploadFile(ctx context.Context, additionalMetadata string, file ReadCloser, petID int64) (uploadFileOK domain.APIResponse, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "upload_file",
			"additionalMetadata", additionalMetadata,
			"file", file,
			"petId", petId,
		)
	}()
}
