package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string
	Description string
	ImageURL    string
}

type ProductCategory struct {
	gorm.Model
	Name string `json:"Name"`
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Category    string
	ImageURL    string
	Quality     string
	Size        string
	Finishing   string
}

type Portfolio struct {
	gorm.Model
	Title       string `json:"Title"`
	Description string `json:"Description"`
	ImageURL    string `json:"ImageURL"`
}

type Message struct {
	gorm.Model
	Name    string
	Phone   string
	Email   string
	Message string
	IsRead  bool
}

type Visitor struct {
	ID    uint `gorm:"primarykey"`
	Date  string
	Count int
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}


type Testimoni struct { // New struct for "Testimoni"
	gorm.Model
	ClientName      string    `gorm:"not null" json:"ClientName"`
	TestimonialText string    `gorm:"type:text;not null" json:"TestimonialText"`
	ImagePath       string    `gorm:"default:''" json:"ImagePath"` // Path to the client's image/avatar
	VideoPath       string    `gorm:"default:''" json:"VideoPath"` // Path to the client's video
	PortfolioID     *uint      `json:"PortfolioID"`                        // Foreign key to Portfolio
	Portfolio       Portfolio `gorm:"foreignKey:PortfolioID" json:"Portfolio"` // Belongs to relationship
}
