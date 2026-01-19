package storeDto

type CreateStoreRequest struct {
	Name        string        `json:"name"`
	OwnerID     string        `json:"owner_id"`
	Description string        `json:"description"`
	Settings    StoreSettings `json:"settings"`
}

type StoreSettings struct {
	CurrencyID     string `json:"currency_id"`
	LanguageID     string `json:"language_id"`
	InventoryAlert bool   `json:"inventory_alert"`
}

type GetStoreByIDRequest struct {
	StoreID string `json:"store_id"`
}

type GetStoreByIDResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	Image       string `json:"image,omitempty"`
	Rating      string `json:"rating,omitempty"`
	Settings    string `json:"settings"`
}

type FindStoreRequest struct {
	Name string `json:"name"`
}

type FindStoreResponse struct {
	Stores []StoreSummary `json:"stores"`
}

type StoreSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image,omitempty"`
	Description string `json:"description"`
	Rating      string `json:"rating,omitempty"`
	OwnerID     string `json:"owner_id"`
}

type GetStoreProductsRequest struct {
	StoreID string `json:"store_id"`
}

type GetStoreProductsResponse struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	OwnerID     string          `json:"owner_id"`
	Image       string          `json:"image,omitempty"`
	Rating      string          `json:"rating,omitempty"`
	Products    []StoreProducts `json:"products"`
}

type StoreProducts struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image,omitempty"`
	Stock       int     `json:"stock"`
}

type GetStoreByOwnerIDRequest struct {
	OwnerID string `json:"owner_id"`
}

type GetStoreByOwnerIDResponse struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	OwnerID     string          `json:"owner_id"`
	Image       string          `json:"image,omitempty"`
	Rating      string          `json:"rating,omitempty"`
	Products    []StoreProducts `json:"products"`
}
