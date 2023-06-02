package model

import (
	"github.com/goccy/go-json"
)

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(100)" json:"name"`
	Email     string `gorm:"type:varchar(100)" json:"email"`
	Avatar    string `gorm:"default:null" json:"avatar"`
	Password  string `json:"password"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
