package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

type PetEndpoints struct {
	AddPetEndpoint            endpoint.Endpoint
	DeletePetEndpoint         endpoint.Endpoint
	FindPetsByStatusEndpoint  endpoint.Endpoint
	FindPetsByTagsEndpoint    endpoint.Endpoint
	GetPetByIDEndpoint        endpoint.Endpoint
	UpdatePetEndpoint         endpoint.Endpoint
	UpdatePetWithFormEndpoint endpoint.Endpoint
	UploadFileEndpoint        endpoint.Endpoint
}

func MakePetEndpoints(s domain.PetService) PetEndpoints {
	return PetEndpoints{
		AddPetEndpoint:            MakeAddPetEndpoint(s),
		DeletePetEndpoint:         MakeDeletePetEndpoint(s),
		FindPetsByStatusEndpoint:  MakeFindPetsByStatusEndpoint(s),
		FindPetsByTagsEndpoint:    MakeFindPetsByTagsEndpoint(s),
		GetPetByIDEndpoint:        MakeGetPetByIDEndpoint(s),
		UpdatePetEndpoint:         MakeUpdatePetEndpoint(s),
		UpdatePetWithFormEndpoint: MakeUpdatePetWithFormEndpoint(s),
		UploadFileEndpoint:        MakeUploadFileEndpoint(s),
	}
}

type addPetRequest struct {
	Body
}

type deletePetRequest struct {
	APIKey

	PetID
}

type findPetsByStatusRequest struct {
	Status
}

type findPetsByTagsRequest struct {
	Tags
}

type getPetByIdRequest struct {
	PetID
}

type updatePetRequest struct {
	Body
}

type updatePetWithFormRequest struct {
	Name

	PetID

	Status
}

type uploadFileRequest struct {
	AdditionalMetadata

	File

	PetID
}
