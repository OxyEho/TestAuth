package handlers

import (
	"awesomeProject/entities"
	"awesomeProject/services"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

var secret = services.GenerateBytes(64)

type JsonToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type CustomClaims struct {
	UserId   string
	AccessId string
	jwt.StandardClaims
}

type Handler struct {
	authService *services.AuthService
}

func NewHandler(authService *services.AuthService) *Handler {
	return &Handler{authService: authService}
}

func (handler *Handler) Start() {
	e := echo.New()
	e.GET("/login", handler.Login)
	e.POST("/refresh", handler.Refresh)
	log.Fatal(e.Start(":8080"))
}

func (handler *Handler) Login(c echo.Context) error {
	id := c.QueryParam("id")
	token, err := handler.authService.Login(id)
	if err != nil {
		return c.JSON(404, errorResp(err))
	}
	jsonToken, err := newJsonToken(token)
	if err != nil {
		return c.JSON(400, errorResp(err))
	}
	return c.JSON(201, jsonToken)
}

func (handler *Handler) Refresh(c echo.Context) error {
	var jsonTokenIn JsonToken
	err := json.NewDecoder(c.Request().Body).Decode(&jsonTokenIn)
	if err != nil {
		return c.JSON(400, errorResp(err))
	}
	jwtToken, err := jwt.ParseWithClaims(jsonTokenIn.Access, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	claims := jwtToken.Claims.(*CustomClaims)
	tokenIn := entities.NewTokenIn(claims.UserId, claims.AccessId, jsonTokenIn.Refresh, claims.StandardClaims.ExpiresAt)
	refreshedToken, err := handler.authService.Refresh(tokenIn)
	if err != nil {
		return c.JSON(400, errorResp(err))
	}
	jsonTokenOut, err := newJsonToken(refreshedToken)
	if err != nil {
		return c.JSON(400, errorResp(err))
	}
	return c.JSON(201, jsonTokenOut)
}

func newJsonToken(token *entities.Token) (*JsonToken, error) {
	claims := CustomClaims{
		UserId:   token.UserId,
		AccessId: token.AccessId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	access, err := jwtToken.SignedString(secret)
	return &JsonToken{Access: access, Refresh: token.Refresh}, err
}

func errorResp(err error) interface{} {
	msg := make(map[string]string)
	msg["msg"] = err.Error()
	return msg
}
