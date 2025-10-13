package productDao

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"github.com/abdulmalikraji/e-commerce/dto"
	"gorm.io/gorm"
)

type DataAccess interface {
	// Postgres Data Access Object Methods
	FindAll() ([]models.Product, error)
	FindById(id string) (models.Product, error)
	FindByStoreId(storeId string) ([]models.Product, error)
	FindPopular() ([]models.Product, error)
	FindPopularByStoreId(storeId string) ([]models.Product, error)
	FindByCategoryId(categoryId string) ([]models.Product, error)
	FindByName(name string) (models.Product, error)
	FindProductReviews(id string) (models.Product, error)
	Insert(item models.Product) (models.Product, error)
	Update(item models.Product) error
	SoftDelete(id string) error
	Delete(id string) error
	FindByFilter(filter dto.ProductFilter) ([]models.Product, int64, error)
}

type dataAccess struct {
	db *gorm.DB
}

func New(client connection.Client) DataAccess {
	return dataAccess{
		db: client.PostgresConnection,
	}
}

func (d dataAccess) FindAll() ([]models.Product, error) {
	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("del_flg = ?", false).
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) FindByStoreId(storeId string) ([]models.Product, error) {
	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("store_id = ? AND del_flg = ?", storeId, true, false).
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) FindById(id string) (models.Product, error) {
	var product models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Preload("Reviews.User").
		First(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (d dataAccess) FindPopular() ([]models.Product, error) {
	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("is_popular = ? AND del_flg = ?", true, false).
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Order("rating_count DESC, rating_average DESC").
		Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) FindPopularByStoreId(storeId string) ([]models.Product, error) {
	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("store_id = ? AND is_popular = ? AND del_flg = ?", storeId, true, false).
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Order("rating_count DESC, rating_average DESC").
		Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) FindByCategoryId(categoryId string) ([]models.Product, error) {
	var products []models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("category_id = ? AND del_flg = ?", categoryId, false).
		Preload("Category").
		Preload("Images").
		Preload("Reviews.User").
		Preload("Reviews.Variant").
		Preload("SubCategories").
		Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return products, nil
}

func (d dataAccess) FindByName(name string) (models.Product, error) {

	var product models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("name = ? AND del_flg = ?", name, false).
		Preload("Category").
		First(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (d dataAccess) FindProductReviews(id string) (models.Product, error) {

	var product models.Product
	result := d.db.Table(models.Product{}.TableName()).
		Where("id = ? AND del_flg = ?", id, false).
		Preload("Category").
		Preload("Reviews.User").
		First(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (d dataAccess) Insert(item models.Product) (models.Product, error) {

	result := d.db.Table(item.TableName()).Create(&item)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return item, nil
}

func (d dataAccess) Update(item models.Product) error {

	result := d.db.Table(item.TableName()).
		Where("id = ? ", item.ID).
		Updates(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) SoftDelete(id string) error {

	var item models.Product

	result := d.db.Table(item.TableName()).
		Where("id = ? ", id).
		Update("del_flg", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d dataAccess) Delete(id string) error {

	var item models.Product

	result := d.db.Table(item.TableName()).
		Where("id = ? ", id).
		Delete(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ProductFilter allows flexible filtering of products
// Add more fields as needed
// Pagination: Page (1-based), PageSize

func (d dataAccess) FindByFilter(filter dto.ProductFilter) ([]models.Product, int64, error) {
	query := d.db.Table(models.Product{}.TableName()).Where("del_flg = ?", false)

	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}
	if filter.CategoryID != nil {
		query = query.Where("category_id = ?", *filter.CategoryID)
	}
	if len(filter.SubCatIDs) > 0 {
		query = query.Joins("JOIN product_subcategories ON products.id = product_subcategories.product_id").
			Where("product_subcategories.subcategory_id IN ?", filter.SubCatIDs)
	}
	if filter.MinRating != nil {
		query = query.Where("rating_average >= ?", *filter.MinRating)
	}
	if filter.DiscountOnly {
		query = query.Where("is_discounted = ?", true)
	}
	if filter.IsPopular != nil {
		query = query.Where("is_popular = ?", *filter.IsPopular)
	}

	// Count total for pagination
	var total int64
	query.Count(&total)

	// Pagination
	page := filter.Page
	if page < 1 {
		page = 1
	}
	pageSize := filter.PageSize
	if pageSize <= 0 {
		pageSize = 20 // default page size
	}
	offset := (page - 1) * pageSize

	query = query.
		Preload("Category").
		Preload("Images").
		Preload("Variants").
		Preload("WarehouseStock").
		Preload("Tags").
		Preload("SubCategories").
		Preload("Reviews.User").
		Offset(offset).
		Limit(pageSize)

	var products []models.Product
	result := query.Find(&products)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return products, total, nil
}

// Returns the category hierarchy for a product as an array of names, from root to leaf
func (d dataAccess) GetCategoryHierarchy(productId string) ([]string, error) {
	var product models.Product
	err := d.db.Table(models.Product{}.TableName()).
		Where("id = ? AND del_flg = ?", productId, false).
		Preload("Category.Parent.Parent"). // Preload up to 3 levels
		Preload("SubCategories.Parent.Parent").
		First(&product).Error
	if err != nil {
		return nil, err
	}

	// Main category chain
	catChain := []string{}
	cat := &product.Category
	for cat != nil {
		if cat.Name != "" {
			catChain = append([]string{cat.Name}, catChain...)
		}
		cat = cat.Parent
	}

	// If you want to include subcategory chains, you can loop through product.SubCategories
	// and build similar arrays for each subcategory

	return catChain, nil
}
