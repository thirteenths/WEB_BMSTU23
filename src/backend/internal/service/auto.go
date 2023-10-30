package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelDB"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
	"time"
)

const (
	salt       = "kjhsfkshfihesifhash"
	tokenTTL   = 12 * time.Hour
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int `json:"id_user"`
	UserRole int `json:"user_role"`
}

type AuthService struct {
	storage storage.IStorage
}

func NewAuthService(storage storage.IStorage) *AuthService {
	return &AuthService{storage: storage}
}

func (s *AuthService) CreateUser(person modelUI.Person) (int, error) {
	person.Hash = generationPasswordHash(person.Hash)
	return s.storage.Repositories()[storage.PERSON].Insert(person)
}

func (s *AuthService) GenerationToken(login string, password string) (string, error) {
	persons, err := s.storage.Repositories()[storage.PERSON].SelectByField("login", login)
	//s.repository.GetUser(login, generationPasswordHash(password))
	if err != nil {
		return "", err
	}

	if len(persons) == 0 {
		return "", err
	}

	var person modelDB.Person
	person.Id = int(persons[0].([]interface{})[0].(int32))
	person.Email = persons[0].([]interface{})[1].(string)
	person.Login = persons[0].([]interface{})[2].(string)
	person.Hash = persons[0].([]interface{})[4].(string)
	person.Role = int(persons[0].([]interface{})[5].(int32))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		person.Id,
		person.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tockenClaims")
	}

	return claims.UserId, claims.UserRole, nil
}

func generationPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
