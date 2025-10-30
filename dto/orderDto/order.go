package orderDto

// OrderFilter struct for flexible filtering
type OrderFilter struct {
 Status        *string
 DateFrom      *string // ISO8601 date string
 DateTo        *string // ISO8601 date string
 PaymentStatus *string
}