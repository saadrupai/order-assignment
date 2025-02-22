package consts

const OrderStatusPending = "Pending"
const OrderStatusCancelled = "Cancelled"

const DeliveryFee = 60
const CODFee = 12
const MobileNumberRegex = `^(01)[3-9]{1}[0-9]{8}$`
const OrderTypeDelivery = "Delivery"

const (
	ItemTypeDocument = iota + 1
	ItemTypeParcel
)

const ItemTypeStatusParcel = "Parcel"
const ItemTypeStatusDocument = "Document"
