package users

import "go.mongodb.org/mongo-driver/mongo"

type UsersService struct {
	Client *mongo.Client
}
