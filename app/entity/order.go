package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID                 uint      `gorm:"primaryKey" json:"id"`
	ConsignmentID      uuid.UUID `gorm:"type:varchar(255)" json:"consignment_id"`
	StoreID            uint      `json:"store_id"`
	MerchantOrderID    string    `json:"merchant_order_id"`
	RecipientName      string    `json:"recipient_name"`    //required
	RecipientPhone     string    `json:"recipient_phone"`   //required
	RecipientAddress   string    `json:"recipient_address"` //required
	RecipientCity      uint      `json:"recipient_city"`
	RecipientZone      uint      `json:"recipient_zone"`
	RecipientArea      uint      `json:"recipient_area"`
	DeliveryType       uint      `json:"delivery_type"` //required
	ItemType           uint      `json:"item_type"`     //required
	SpecialInstruction string    `json:"special_instruction"`
	ItemQuantity       uint      `json:"item_quantity"`     //required
	ItemWeight         float64   `json:"item_weight"`       //required
	AmountToCollect    string    `json:"amount_to_collect"` //required
	ItemDescription    string    `json:"item_description"`
	OrderStatus        string    `json:"order_status" gorm:"default:'Pending'"`
	DeliveryFee        float64   `json:"delivery_fee"`
	CODFee             float64   `json:"cod_fee"`
	PromoDiscount      float64   `json:"promo_discount"`
	Discount           float64   `json:"discont"`
	OrderType          string    `json:"order_type"`
	TransferStatus     bool      `json:"transfer_status"`
	Archive            bool      `json:"archive"`
}
