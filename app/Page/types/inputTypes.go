package types

// PageInput represents the input type for creating/updating a page

type PageInput struct {
    gorm.Model
    Title string `gorm:"column:title;json:\"title\""`
    Content string `gorm:"column:content;json:\"content\""`
}

// UpdatePageInput represents the input type for updating a page

type UpdatePageInput struct {
    ID      int     `json:"id"`
    Title string `json:"title"`
    Content string `json:"content"`
}
