package storage

import (
	"errors"
	"gin-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserGormStorage struct {
	db *gorm.DB
}

func NewUserGormStorage(db *gorm.DB) *UserGormStorage {
	return &UserGormStorage{db: db}
}

func (g *UserGormStorage) IsEmailAvailable(email string) bool {
	var user models.User
	g.db.Where("email=?", email).First(&user)

	return user.ID == 0
}

func (g *UserGormStorage) MatchUserPassword(email, password string) (*models.User, error) {
	var user models.User
	err := g.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	user.Password = ""

	return &user, nil
}

func (g *UserGormStorage) AddUser(name, email, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// c.Error(models.NewBadRequest("Password is too long"))
		return models.User{}, err
	}
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = g.db.Create(&user).Error
	return user, err
}

func (g *UserGormStorage) GetUsers(limit, offset int, filters map[string]string) ([]models.UserResponse, error) {
	var users []models.UserResponse
	query := g.db.Model(&models.User{})

	if email, ok := filters["email"]; ok {
		query = query.Where("email=?", email)
	}
	if name, ok := filters["name"]; ok {
		query = query.Where("name=?", name)
	}

	err := query.
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	return users, err
}

func (g *UserGormStorage) UpdateUser(id int, name, email string) (models.User, error) {
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

func (g *UserGormStorage) DeleteUser(id int) error {
	return g.db.Delete(&models.User{}, id).Error
}

func (g *UserGormStorage) CountUsers(filters map[string]string) (int64, error) {
	var count int64
	query := g.db.Model(&models.User{})

	if email, ok := filters["email"]; ok {
		query = query.Where("email=?", email)
	}
	if name, ok := filters["name"]; ok {
		query = query.Where("name=?", name)
	}

	err := query.Count(&count).Error
	return count, err
}

func (g *UserGormStorage) GetUser(id int) (models.User, error) {
	var user models.User

	err := g.db.First(&user, id).Error
	return user, err
}
