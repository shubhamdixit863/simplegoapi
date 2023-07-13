package entity

type User struct {
	Id   int    `bson:"_id,omitempty" gorm:"primaryKey,omitempty"` // omit empty says to db that ignore the zero value of string and create the id from your own end
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

// Mysql entities
