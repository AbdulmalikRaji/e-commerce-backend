package productDto

type ProductFilter struct {
	MinPrice     *float64 `json:"min_price,omitempty"`
	MaxPrice     *float64 `json:"max_price,omitempty"`
	CategoryID   *string  `json:"category_id,omitempty"`
	SubCatIDs    []string `json:"sub_cat_ids,omitempty"`
	MinRating    *float64 `json:"min_rating,omitempty"`
	DiscountOnly bool     `json:"discount_only,omitempty"`
	IsPopular    *bool    `json:"is_popular,omitempty"`
	Page         int      `json:"page,omitempty"`
	PageSize     int      `json:"page_size,omitempty"`
}
