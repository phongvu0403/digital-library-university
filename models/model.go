package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AdminAccount = "admin"
	UserAccount  = "user"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Account struct {
	BaseModel
	Email        string `json:"email"`
	AccountRole  string `json:"accountRole"`
	Password     string `gorm:"-" json:"password"`
	PasswordHash string `json:"-"`
}

func (Account) TableName() string {
	return "account"
}

type Book struct {
	BaseModel
	Name          string    `json:"name"`
	ISBN          string    `json:"isbn"`
	Stock         uint      `json:"stock"`
	Author        string    `json:"author"`
	Year          string    `json:"year"`
	Edition       uint      `json:"edition"`
	Cover         string    `json:"cover"`
	Abstract      string    `json:"abstract"`
	Category      string    `json:"category"`
	Available     bool      `json:"available"`
	AvailableDate time.Time `json:"availableDate"`
}

func (Book) TableName() string {
	return "book"
}

type BookHistory struct {
	UserID     uint       `json:"userId"`
	BookID     uint       `json:"bookId"`
	IssueDate  *time.Time `json:"issueDate"`
	ReturnDate *time.Time `json:"returnDate"`
	Returned   bool       `json:"returned"`
}

func (BookHistory) TableName() string {
	return "book_history"
}

type Response struct {
	AccountRole string `json:"accountRole"`
	Token       string `json:"token"`
}

type LoginDetails struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccountRole string `json:"accountRole"`
}

type AuthInfo struct {
	Role string
	jwt.StandardClaims
}