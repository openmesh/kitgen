package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/kitgen/domain"
)

type StoreEndpoints struct {
	DeleteOrderEndpoint  endpoint.Endpoint
	GetInventoryEndpoint endpoint.Endpoint
	GetOrderByIDEndpoint endpoint.Endpoint
	PlaceOrderEndpoint   endpoint.Endpoint
}

func MakeStoreEndpoints(s domain.StoreService) StoreEndpoints {
	return StoreEndpoints{
		DeleteOrderEndpoint:  MakeDeleteOrderEndpoint(s),
		GetInventoryEndpoint: MakeGetInventoryEndpoint(s),
		GetOrderByIDEndpoint: MakeGetOrderByIDEndpoint(s),
		PlaceOrderEndpoint:   MakePlaceOrderEndpoint(s),
	}
}

type deleteOrderRequest struct {
	OrderID
}

type getInventoryRequest struct {
}

type getOrderByIdRequest struct {
	OrderID
}

type placeOrderRequest struct {
	Body
}
