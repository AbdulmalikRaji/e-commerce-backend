package models

import (
	"time"

	"github.com/google/uuid"
)

// ProductView tracks each time a product is viewed by a user or guest (session).
// Used for:
// - Measuring product popularity and engagement
// - Calculating conversion rates (views to purchases)
// - Analyzing user browsing patterns
// - Building product recommendations
type ProductView struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`                     // nullable for guests
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	ViewedAt  time.Time  `gorm:"autoCreateTime;index" json:"viewed_at"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	User    *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (ProductView) TableName() string {
	return "ecom.product_views"
}

// AddToCartEvent records when a user adds a product to their shopping cart.
// Used for:
// - Abandoned cart analysis
// - Product interest tracking
// - Conversion funnel analysis
// - User behavior patterns
type AddToCartEvent struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`                     // nullable for guests
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	Quantity  int        `gorm:"not null" json:"quantity"`
	AddedAt   time.Time  `gorm:"autoCreateTime;index" json:"added_at"`
	CartID    *uuid.UUID `gorm:"type:uuid;index" json:"cart_id"` // nullable for guests

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	User    *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Cart    *Cart   `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cart,omitempty"`
}

func (AddToCartEvent) TableName() string {
	return "ecom.add_to_cart_events"
}

// SalesStat aggregates sales statistics for analysis.
// Used for:
// - Revenue tracking and forecasting
// - Sales performance analysis
// - Product success metrics
// - Seasonal trend analysis
// - Store performance comparison
type SalesStat struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProductID uuid.UUID  `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"`           // optional, for user-specific stats
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	TotalSold int        `gorm:"default:0" json:"total_sold"`
	Revenue   float64    `gorm:"type:numeric(10,2);default:0" json:"revenue"`
	LastSold  *time.Time `json:"last_sold"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}

func (SalesStat) TableName() string {
	return "ecom.sales_stats"
}

type AbandonedCart struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CartID        uuid.UUID  `gorm:"type:uuid;index;not null" json:"cart_id"`
	UserID        *uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	SessionID     *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"` // for guest tracking
	LastUpdatedAt time.Time  `gorm:"autoUpdateTime" json:"last_updated_at"`
	AbandonedAt   time.Time  `gorm:"index" json:"abandoned_at"`

	// Relations
	Cart Cart  `gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"`
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (AbandonedCart) TableName() string {
	return "ecom.abandoned_carts"
}

// StoreVisit tracks user visits to individual store pages.
// Used for:
// - Store popularity metrics
// - User engagement analysis
// - Store performance tracking
// - Traffic pattern analysis
// - Marketing effectiveness measurement
type StoreVisit struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StoreID   uuid.UUID  `gorm:"type:uuid;index;not null" json:"store_id"`
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"` // nullable for guests
	SessionID *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"`
	VisitedAt time.Time  `gorm:"autoCreateTime;index" json:"visited_at"`

	// Relations
	Store Store `gorm:"foreignKey:StoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"store,omitempty"`
	User  *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (StoreVisit) TableName() string {
	return "ecom.store_visits"
}

// SearchAnalytics captures search queries and their results.
// Used for:
// - Search effectiveness evaluation
// - Trending search terms analysis
// - Search result optimization
// - User intent understanding
// - Product discoverability improvement
type SearchAnalytics struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID     *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"` // nullable for guests
	SessionID  *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"`
	Query      string     `gorm:"type:text;not null" json:"query"`
	Results    int        `json:"results"`
	SearchedAt time.Time  `gorm:"autoCreateTime;index" json:"searched_at"`

	// Relations
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (SearchAnalytics) TableName() string {
	return "ecom.search_analytics"
}

// MarketingAnalytics tracks the effectiveness of marketing campaigns and promotions.
// Used for:
// - Campaign performance measurement
// - ROI analysis
// - Customer acquisition tracking
// - Promotional effectiveness
// - Marketing channel optimization
type MarketingAnalytics struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CampaignID    uuid.UUID  `gorm:"type:uuid;index;not null" json:"campaign_id"`
	StoreID       uuid.UUID  `gorm:"type:uuid;index;not null" json:"store_id"`
	UserID        *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"` // nullable for guests
	SessionID     *string    `gorm:"type:varchar(64);index" json:"session_id,omitempty"`
	Action        string     `gorm:"type:varchar(50);not null" json:"action"` // e.g., view, click, purchase
	ActionDetails string     `gorm:"type:text" json:"action_details,omitempty"`
	OccurredAt    time.Time  `gorm:"autoCreateTime;index" json:"occurred_at"`

	// Relations
	Store Store `gorm:"foreignKey:StoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"store,omitempty"`
	User  *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

func (MarketingAnalytics) TableName() string {
	return "ecom.marketing_campaign_analytics"
}
