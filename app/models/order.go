package models

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/saadrupai/order-assignment/app/consts"
	"regexp"
)

type OrderReqBody struct {
	StoreID            uint    `json:"store_id"`
	MerchantOrderID    string  `json:"merchant_order_id"`
	RecipientName      string  `json:"recipient_name"`    //required
	RecipientPhone     string  `json:"recipient_phone"`   //required
	RecipientAddress   string  `json:"recipient_address"` //required
	RecipientCity      uint    `json:"recipient_city"`
	RecipientZone      uint    `json:"recipient_zone"`
	RecipientArea      uint    `json:"recipient_area"`
	DeliveryType       uint    `json:"delivery_type"` //required
	ItemType           uint    `json:"item_type"`     //required
	SpecialInstruction string  `json:"special_instruction"`
	ItemQuantity       uint    `json:"item_quantity"`     //required
	ItemWeight         float64 `json:"item_weight"`       //required
	AmountToCollect    string  `json:"amount_to_collect"` //required
	ItemDescription    string  `json:"item_description"`
}

func (orderReq *OrderReqBody) Validate() error {
	return validation.ValidateStruct(orderReq, validation.Field(&orderReq.RecipientName, validation.Required.Error("The recipient name field is required.")),
		validation.Field(&orderReq.RecipientPhone, validation.Required.Error("The recipient phone field is required."),
			validation.By(func(value interface{}) error {
				strValue, ok := value.(string)
				if !ok {
					return fmt.Errorf("%v must be a string", value)
				}

				reg := regexp.MustCompile(consts.MobileNumberRegex)
				if !reg.MatchString(strValue) {
					return fmt.Errorf("invalid phone number format")
				}

				return nil
			})),
		validation.Field(&orderReq.RecipientAddress, validation.Required.Error("The recipient address field is required.")),
		validation.Field(&orderReq.DeliveryType, validation.Required.Error("The delivery type field is required.")),
		validation.Field(&orderReq.AmountToCollect, validation.Required.Error("The amount to collect field is required.")),
		validation.Field(&orderReq.ItemQuantity, validation.Required.Error("The item quantity field is required.")),
		validation.Field(&orderReq.ItemWeight, validation.Required.Error("The item weight field is required.")),
		validation.Field(&orderReq.ItemType, validation.Required.Error("The item type field is required.")),
	)
}

type Response struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"` // Use interface{} for dynamic data
	Errors  interface{} `json:"errors,omitempty"`
}

type OrderCreateResponse struct {
	ConsignmentID   uint    `json:"consignment_id"`
	MerchantOrderID string  `json:"merchant_order_id"`
	OrderStatus     string  `json:"order_status"`
	DeliveryFee     float64 `json:"delivery_fee"`
}
