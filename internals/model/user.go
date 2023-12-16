package model

type User struct {
	Username string   `gorm:"column:username;primary_key;type:varchar(100);not null" json:"username" validate:"required,min=5,max=100,alphanum"`
	Password string   `gorm:"column:password;type:text;not null" json:"password" validate:"required,min=6,max=100"`
	Fullname string   `gorm:"column:fullname;type:varchar(100);not null" json:"fullname" validate:"required,min=5,max=100"`
	Token    string   `gorm:"column:token;type:text;default:null" json:"token"`
	Threads  []Thread `gorm:"foreignKey:owner;references:username"`
}

func (user *User) TableName() string {
	return "users"
}
