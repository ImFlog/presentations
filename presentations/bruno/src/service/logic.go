package service

import (
	"catalog/domain"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	productRepo domain.Repository
}

// NewProductService a service struct that attaches to a repository via the Repository interface
func NewProductService(productRepo domain.Repository) *service {
	return &service{productRepo: productRepo}
}

func (s *service) Find(code string) (*domain.Product, error) {
	return s.productRepo.Find(code)
}
func (s *service) Store(product *domain.Product) error {
	return s.productRepo.Store(product)
}
func (s *service) Update(product *domain.Product) error {
	return s.productRepo.Update(product)
}

func (s *service) FindAll() ([]*domain.Product, error) {
	return s.productRepo.FindAll()
}

func (s *service) Delete(code string) error {
	return s.productRepo.Delete(code)
}

type loginService struct{}

func NewLoginService() *loginService {
	return &loginService{}
}

func (ls *loginService) Login(login domain.Login) (string, error) {
	// Hardcoded allowed password
	if login.Password != "password" {
		return "", domain.ErrInvalidLogin
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": login,
	})
	return token.SignedString([]byte("secret"))
}
