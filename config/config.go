package config

import "os"

// SesExtMonth represents the user token expire length
const SesExtMonth = 1

// TokenSecret use to create user login token
var TokenSecret = os.Getenv("TOKEN_SECRET")

// MongoURI ex: MONGO_URI=mongodb://localhost:27017/pos_server
var MongoURI = os.Getenv("MONGO_URI")
