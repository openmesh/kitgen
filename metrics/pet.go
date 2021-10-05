package metrics

import (
	"context"
	"io"
	"time"

	"github.com/go-kit/kit/metrics"
)

func PetMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) domain.PetServiceMiddleware {
	return func(next domain.PetService) domain.PetService {
		return petMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type petMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	next            domain.PetService
}

func (mw *petMetricsMiddleware) AddPet(ctx context.Context, body domain.Pet) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "add_pet"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.AddPet(ctx, body)
}

func (mw *petMetricsMiddleware) DeletePet(ctx context.Context, apiKey string, petId int64) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete_pet"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.DeletePet(ctx, apiKey, petId)
}

func (mw *petMetricsMiddleware) FindPetsByStatus(ctx context.Context, status []string) (_ domain.ListPetsResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_pets_by_status"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.FindPetsByStatus(ctx, status)
}

func (mw *petMetricsMiddleware) FindPetsByTags(ctx context.Context, tags []string) (_ []*domain.Pet, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find_pets_by_tags"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.FindPetsByTags(ctx, tags)
}

func (mw *petMetricsMiddleware) GetPetByID(ctx context.Context, petId int64) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_pet_by_id"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.GetPetByID(ctx, petId)
}

func (mw *petMetricsMiddleware) UpdatePet(ctx context.Context, body domain.Pet) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update_pet"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.UpdatePet(ctx, body)
}

func (mw *petMetricsMiddleware) UpdatePetWithForm(ctx context.Context, name string, petId int64, status string) (_ domain.Pet, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update_pet_with_form"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.UpdatePetWithForm(ctx, name, petId, status)
}

func (mw *petMetricsMiddleware) UploadFile(ctx context.Context, additionalMetadata string, file io.ReadCloser, petId int64) (_ domain.APIResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "upload_file"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.UploadFile(ctx, additionalMetadata, file, petId)
}
