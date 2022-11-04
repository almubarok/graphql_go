package model

import (
	"graphql_go/src/modules/user/domain"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `gorm:"type:bigserial"`
	Name      string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m User) ToDomain() domain.User {
	return domain.User{
		ID:        m.ID,
		Name:      m.Name,
		Username:  m.Username,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
	}
}
