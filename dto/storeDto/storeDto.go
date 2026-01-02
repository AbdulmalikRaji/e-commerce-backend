package storeDto

type CreateStoreRequest struct {
	Name        string        `json:"name"`
	OwnerID     string        `json:"owner_id"`
	Description string        `json:"description"`
	Settings    StoreSettings `json:"settings"`
}

type StoreSettings struct {
	Currency       string `json:"currency"`
	Language       string `json:"language"`
	InventoryAlert bool   `json:"inventory_alert"`
}
