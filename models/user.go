package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the model that governs all notes objects retrived or inserted into the DB
//User ID is a unique string that identifies users.
//token is the signed jwt token with the user details
//refresh token is empty token 
type User struct {
    ID            primitive.ObjectID `bson:"_id"`
    First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
    Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
    Password      *string            `json:"password" validate:"required,min=6""`
    Email         *string            `json:"email" validate:"email,required"`
    Phone         *string            `json:"phone" validate:"required"`
    Token         *string            `json:"token"`
    Refresh_token *string            `json:"refresh_token"`
    User_id       string             `json:"user_id"`
}