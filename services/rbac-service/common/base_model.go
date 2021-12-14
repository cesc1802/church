package common

import "time"

type BaseModel struct {
	ID        uint64     `gorm:"column:id"`
	Status    int64      `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}
