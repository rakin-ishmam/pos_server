package main

import "os"

// mongoURI ex: MONGO_URI=mongodb://localhost:27017/pos_server
var mongoURI = os.Getenv("MONGO_URI")
