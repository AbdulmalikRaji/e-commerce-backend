package storeDto

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// ValidateCreateStore validates the CreateStoreRequest fields.
func ValidateCreateStore(req CreateStoreRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if _, err := uuid.Parse(req.OwnerID); err != nil {
		return errors.New("owner_id must be a valid UUID")
	}

	if strings.TrimSpace(req.Settings.CurrencyID) == "" {
		return errors.New("settings.currency_id is required")
	}
	if _, err := uuid.Parse(req.Settings.CurrencyID); err != nil {
		return errors.New("settings.currency_id must be a valid UUID")
	}

	if strings.TrimSpace(req.Settings.LanguageID) == "" {
		return errors.New("settings.language_id is required")
	}
	if _, err := uuid.Parse(req.Settings.LanguageID); err != nil {
		return errors.New("settings.language_id must be a valid UUID")
	}

	return nil
}
