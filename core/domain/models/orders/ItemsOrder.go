package orders

import (
	"CaliYa/core/domain/dto/order"
	"CaliYa/core/domain/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ItemsOrder struct {
	bun.BaseModel `bun:"table:orders.order_items"`

	OrderID     uuid.UUID `bun:"order_id"`
	ItemID      uuid.UUID `bun:"item_id"`
	Amount      int       `bun:"amount"`
	UnitPrice   int       `bun:"unit_price"`
	TotalPrice  int       `bun:"total_price"`
	Observation string    `bun:"observation"`
}

func (io *ItemsOrder) BuildDtoToModel(dto order.ItemsOrder, orderID uuid.UUID) {
	io.OrderID = orderID
	io.ItemID = dto.ItemID
	io.Amount = dto.Amount
	io.Observation = dto.Observation
}

func (o *Order) SetPrices(items []models.Items) {

	for _, item := range items {
		for i, io := range o.ItemsOrder {
			if io.ItemID == item.ID {

				var subTotal int

				subTotal += item.Price * io.Amount

				o.ItemsOrder[i].TotalPrice = subTotal
				o.ItemsOrder[i].UnitPrice = item.Price

				o.TotalPrice += subTotal

			}
		}
	}

}
