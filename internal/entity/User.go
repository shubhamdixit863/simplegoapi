package entity

type User struct {
	Id   string `bson:"_id,omitempty"` // omit empty says to db that ignore the zero value of string and create the id from your own end
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
