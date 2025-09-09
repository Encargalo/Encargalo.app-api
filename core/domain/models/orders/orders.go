package orders

import (
	"CaliYa/core/domain/dto/order"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel `bun:"table:orders.orders"`

	ID            uuid.UUID `bun:"id,pk,type:uuid"`
	ShopID        uuid.UUID `bun:"shop_id"`
	CustomerID    uuid.UUID `bun:"customer_id"`
	Address       string    `bun:"address"`
	Latitude      float64   `bun:"latitude"`
	Longitude     float64   `bun:"longitude"`
	MethodPayment string    `bun:"method_payment"`
	DeliveryFee   int       `bun:"delivery_fee"`
	TotalPrice    int       `bun:"total_price"`

	ItemsOrder []ItemsOrder `bun:"rel:has-many,join:id=order_id"`
}

func (o *Order) BuildDtoToModel(dto order.CreateOrder) {
	o.ID = dto.ID
	o.ShopID = dto.ShopID
	o.CustomerID = dto.CustomerID
	o.Address = dto.Address
	o.Latitude = dto.Coords.Latitude
	o.Longitude = dto.Coords.Longitude
	o.MethodPayment = dto.MethodPayment
	o.DeliveryFee = 0
	o.ItemsOrder = make([]ItemsOrder, len(dto.Items))

	for i := range dto.Items {
		o.ItemsOrder[i].BuildDtoToModel(dto.Items[i], o.ID)
	}
}

func (o *Order) GetItemsID() []uuid.UUID {

	ret := make([]uuid.UUID, len(o.ItemsOrder))

	for i := range ret {
		ret[i] = o.ItemsOrder[i].ItemID
	}

	return ret

}
