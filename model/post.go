package model

type Post struct {
	ID         uint   `json:"id" gorm:"primary key"`
	UserId     uint   `json:"user_id" gorm:"not null"`
	CategoryId uint   `json:"category_id" gorm:"not null"`
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}
