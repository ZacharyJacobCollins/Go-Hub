package main

import (
  "fmt"
   "labix.org/v2/mgo"
   "labix.org/v2/mgo/bson"
)

type User struct {
  ID       bson.ObjectId `bson:"_id,omitempty"`
  Username string
  Password []byte
  Posts    int
}
