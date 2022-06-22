package repositpries

import (
	"awesomeProject/db"
	"awesomeProject/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type IAuthRepo interface {
	GetActiveToken(userId string, accessId string) (*entities.Token, error)
	SetAllExpired(userId string)
	SetExpired(userId string, accessId string)
	CreateToken(token *entities.Token)
}

type AuthRepo struct{}

func (repo *AuthRepo) GetActiveToken(userId string, accessId string) (*entities.Token, error) {
	var t entities.Token
	filter := bson.M{
		"userId":   userId,
		"accessId": accessId,
		"expired":  false,
	}
	res := db.Get(filter)
	_ = res.Decode(&t)
	return &t, res.Err()
}

func (repo *AuthRepo) SetAllExpired(userId string) {
	filter := bson.D{
		{"userId", userId},
	}
	update := bson.D{{"$set", bson.D{{"expired", true}}}}
	_, _ = db.UpdateDocs(filter, update)
}

func (repo *AuthRepo) SetExpired(userId string, accessId string) {
	filter := bson.D{
		{"userId", userId},
		{"accessId", accessId},
	}
	update := bson.D{{"$set", bson.D{{"expired", true}}}}
	_, _ = db.UpdateDoc(filter, update)
}

func (repo *AuthRepo) CreateToken(token *entities.Token) {
	_, _ = db.InsertDoc(token)
}
