package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Thread struct {
	ID        uuid.UUID `gorm:"column:id;type:varchar(200);primary_key" json:"id"`
	Content   string    `gorm:"column:content;type:varchar(200);not null" json:"content" validate:"required,max=200"`
	Owner     string    `gorm:"column:owner;type:text;not null" json:"owner"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	User      User      `gorm:"foreignKey:owner;references:username" validate:"-"`
}

func (thread *Thread) TableName() string {
	return "threads"
}

func (thread *Thread) BeforeCreate(tx *gorm.DB) (err error) {
	thread.ID = uuid.New()
	return
}
