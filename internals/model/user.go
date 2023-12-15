package model

type User struct {
	Username string   `gorm:"column:username;primary_key;type:varchar(100);not null" json:"username"`
	Password string   `gorm:"column:password;type:text;not null" json:"password"`
	Fullname string   `gorm:"column:fullname;type:varchar(100);not null" json:"fullname"`
	Token    string   `gorm:"column:token;type:text" json:"token"`
	Threads  []Thread `gorm:"foreignKey:owner;references:username"`
}

func (user *User) TableName() string {
	return "users"
}
