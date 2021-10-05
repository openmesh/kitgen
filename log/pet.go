package log

import (
	"context"
	"io"
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

func (mw *petLoggingMiddleware) AddPet(ctx context.Context, body domain.Pet) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "add_pet",
			"body", body,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.AddPet(ctx, body)
}

func (mw *petLoggingMiddleware) DeletePet(ctx context.Context, apiKey string, petId int64) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_pet",
			"apiKey", apiKey,
			"petId", petId,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.DeletePet(ctx, apiKey, petId)
}

func (mw *petLoggingMiddleware) FindPetsByStatus(ctx context.Context, status []string) (_ domain.ListPetsResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_pets_by_status",
			"status", status,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.FindPetsByStatus(ctx, status)
}

func (mw *petLoggingMiddleware) FindPetsByTags(ctx context.Context, tags []string) (_ []*domain.Pet, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find_pets_by_tags",
			"tags", tags,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.FindPetsByTags(ctx, tags)
}

func (mw *petLoggingMiddleware) GetPetByID(ctx context.Context, petId int64) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_pet_by_id",
			"petId", petId,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.GetPetByID(ctx, petId)
}

func (mw *petLoggingMiddleware) UpdatePet(ctx context.Context, body domain.Pet) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_pet",
			"body", body,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.UpdatePet(ctx, body)
}

func (mw *petLoggingMiddleware) UpdatePetWithForm(ctx context.Context, name string, petId int64, status string) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_pet_with_form",
			"name", name,
			"petId", petId,
			"status", status,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.UpdatePetWithForm(ctx, name, petId, status)
}

func (mw *petLoggingMiddleware) UploadFile(ctx context.Context, additionalMetadata string, file io.ReadCloser, petId int64) (_ domain.APIResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "upload_file",
			"additionalMetadata", additionalMetadata,
			"file", file,
			"petId", petId,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.UploadFile(ctx, additionalMetadata, file, petId)
}
