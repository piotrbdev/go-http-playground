package storage

import (
	"fmt"
	"gin-app/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to postgres (gorm)")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type GormStorage struct {
	db *gorm.DB
}

func NewGormStorage(db *gorm.DB) *GormStorage {
	return &GormStorage{db: db}
}

func (g *GormStorage) AddUser(name, email string) (models.User, error) {
	user := models.User{
		Name:  name,
		Email: email,
	}

	err := g.db.Create(&user).Error
	return user, err
}

func (g *GormStorage) GetUser(id int) (models.User, error) {
	var user models.User

	err := g.db.First(&user, id).Error
	return user, err
}

func (g *GormStorage) GetUsers(limit, offset int) ([]models.User, error) {
	var users []models.User

	err := g.db.
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	return users, err
}

func (g *GormStorage) UpdateUser(id int, name, email string) (models.User, error) {
	var user models.User

	err := g.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	user.Name = name
	user.Email = email
	err = g.db.Save(&user).Error
	return user, err
}

func (g *GormStorage) DeleteUser(id int) error {
	return g.db.Delete(&models.User{}, id).Error
}
