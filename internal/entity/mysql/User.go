package mysql

// Particularly for relational Dbs

type User struct {
	Id   int    `gorm:"primaryKey,omitempty"` // omit empty says to db that ignore the zero value of string and create the id from your own end
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
