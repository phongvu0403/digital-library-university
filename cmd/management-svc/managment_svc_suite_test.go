package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/library/envConfig"
	"github.com/library/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestManagement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Management-Svc Handler Tests")
}

func setupAuthInfo(env *envConfig.Env, db *gorm.DB) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   1010,
		"role": models.AdminAccount,
	})
	adminToken, err := token.SignedString([]byte(env.JwtSigningKey))
	if err != nil {
		return "", "", err
	}

	user := &models.Account{
		BaseModel:   *&models.BaseModel{ID: 1010},
		Email:       "unit@user.com",
		AccountRole: "user",
	}
	err = db.Create(user).Error
	if err != nil {
		return "", "", err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   1010,
		"role": models.UserAccount,
	})
	userToken, err := token.SignedString([]byte(env.JwtSigningKey))
	if err != nil {
		return "", "", err
	}

	return adminToken, userToken, err
}

func setupTestData(db *gorm.DB) error {
	book := models.Book{
		BaseModel: *&models.BaseModel{ID: 1010},
		Name:      "testBook",
		Author:    "testAuthor",
	}
	err := db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func cleanTestData(db *gorm.DB) error {
	if err := db.Exec(`delete from account where id = ?`, "1010").Error; err != nil {
		return err
	}
	if err := db.Exec(`delete from author where id = ?`, "1010").Error; err != nil {
		return err
	}
	if err := db.Exec(`delete from subject where id = ?`, "1010").Error; err != nil {
		return err
	}
	if err := db.Exec(`delete from book where id = ?`, "1010").Error; err != nil {
		return err
	}
	return nil
}
