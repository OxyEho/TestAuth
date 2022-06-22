package services

import (
	"awesomeProject/entities"
	"awesomeProject/repositpries"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	repo repositpries.IAuthRepo
}

func NewAuthService(repo repositpries.AuthRepo) *AuthService {
	return &AuthService{repo: &repo}
}

func (service *AuthService) Login(userId string) (*entities.Token, error) {
	if !isValidUUID(userId) {
		return nil, errors.New("bad user id")
	}
	token := entities.NewToken(userId, uuid.NewString(), GetRefresh())
	hashedToken := getHashToken(token)
	service.repo.SetAllExpired(userId)
	service.repo.CreateToken(hashedToken)
	return token, nil
}

func (service *AuthService) Refresh(tokenIn *entities.TokenIn) (*entities.Token, error) {
	token, err := service.repo.GetActiveToken(tokenIn.UserId, tokenIn.AccessId)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(token.Refresh), []byte(tokenIn.Refresh))
	if err != nil {
		return nil, errors.New("bad token")
	}
	if token.Expired || token.RefreshExpiresAt < time.Now().Unix() {
		return nil, errors.New("refresh expired")
	}
	service.repo.SetExpired(token.UserId, tokenIn.AccessId)
	token = entities.NewToken(token.UserId, uuid.NewString(), tokenIn.Refresh)
	service.repo.CreateToken(getHashToken(token))
	return token, nil
}

func getHashToken(token *entities.Token) *entities.Token {
	return &entities.Token{
		UserId:           token.UserId,
		AccessId:         token.AccessId,
		Refresh:          hash(token.Refresh),
		RefreshExpiresAt: token.RefreshExpiresAt,
		Expired:          false,
	}
}

func hash(t string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(t), 6)
	return string(bytes)
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func GetRefresh() string {
	return base64.StdEncoding.EncodeToString(GenerateBytes(24))
}

func GenerateBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
