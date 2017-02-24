package config

import "os"

// SesExtMonth represents the user token expire length
const SesExtMonth = 1

// TokenSecret use to create user login token, ex: TOKEN_SECRET=hagu
var TokenSecret string

// MongoURI ex: MONGO_URI=mongodb://localhost:27017/pos_server
var MongoURI string

func init() {
	TokenSecret = os.Getenv("TOKEN_SECRET")
	if TokenSecret == "" {
		TokenSecret = "hagu"
	}

	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		MongoURI = "mongodb://localhost:27017/pos_server"
	}
}
