package company

import "gopkg.in/mgo.v2/bson"

// ID type
type ID bson.ObjectId

// NewID creates a new object id
func NewID() ID {
	return ID(bson.NewObjectId())
}

// StringToID converts a string to type ID
func StringToID(s string) ID {
	return ID(bson.ObjectIdHex(s))
}

// String convert an ID in a string
func (i ID) String() string {
	return bson.ObjectId(i).Hex()
}
