package messages

var (
	RequiredField = "required_field"
	InvalidFormat = "invalid_format"
	InvalidToken  = "invalid_token"
)

// Templates holds message templates keyed by message id.
// Keep templates here so they can be localized/overridden in one place.
var Templates = map[string]string{
	RequiredField: "{{.Field}} not provided",
	InvalidFormat: "{{.Field}} has invalid format",
	InvalidToken:  "Invalid or expired authentication token",
}
