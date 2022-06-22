package entities

import (
	"time"
)

type Token struct {
	UserId           string `bson:"userId"`
	AccessId         string `bson:"accessId"`
	Refresh          string `bson:"refresh"`
	RefreshExpiresAt int64  `bson:"refreshExpiresAt"`
	Expired          bool   `bson:"expired"`
}

func NewToken(userId string, accessId string, refresh string) *Token {
	return &Token{
		UserId:           userId,
		AccessId:         accessId,
		Refresh:          refresh,
		RefreshExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Expired:          false,
	}
}

type TokenIn struct {
	UserId    string
	AccessId  string
	Refresh   string
	ExpiredAt int64
}

func NewTokenIn(userId string, accessId string, refresh string, expiredAt int64) *TokenIn {
	return &TokenIn{
		UserId:    userId,
		AccessId:  accessId,
		Refresh:   refresh,
		ExpiredAt: expiredAt,
	}
}
