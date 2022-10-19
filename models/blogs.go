package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:100;unique"`
	Approved  bool      `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Blogs     []Blog    `gorm:"foreignKey:CategoryID"` // Has Many Relations
}

// Blog implement for blog model
type Blog struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:100;unique;index"`
	AuthorID    uint      // Author ID foreign key
	CategoryID  uint      // Category ID foreign key
	Logo        string    `gorm:"size:255;null"`
	Description string    `gorm:"size:255"`
	Content     string    `gorm:"type:longtext;null"`
	Views       int       `gorm:"default:0"`
	Status      bool      `gorm:"default:0"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:nano"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

// CreateNewBlogCategory create new blog category
// If data is exists return error
func CreateNewBlogCategory(category *Category) error {
	err = db.Create(category).Error
	return err
}

// GetBlogCategoryByID return category by passing the ID
func GetBlogCategoryByID(id uint) (Category, error) {
	var category Category
	err = db.Model(&Category{}).Where("id = ?", id).First(&category).Error
	return category, err
}

// CreateNewBlog create new blog if data is exists return error
func CreateNewBlog(blog *Blog) error {
	err = db.Create(blog).Error
	return err
}

// GetBlogByID return blog by passing the ID
func GetBlogByID(id uint) (Blog, error) {
	var blog Blog
	err := db.Model(&Blog{}).Where("id = ?", id).First(&blog).Error
	return blog, err
}

// GetBlogByTitle return blog  by passing by Title
func GetBlogByTitle(title string) (Blog, error) {
	var blog Blog
	err = db.Model(&Blog{}).Where("title = ?", title).First(&blog).Error
	return blog, err
}
