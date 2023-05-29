package entity

import "time"

type ContentCategory struct {
	ID         int32      `gorm:"primaryKey" json:"id"`
	ContentID  int32      `gorm:"type:int" json:"content_id"`
	CategoryID int32      `gorm:"type:int" json:"category_id"`
	CreatedAt  time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"type:timestamp" json:"deleted_at"`
	Deleted    bool       `gorm:"type:bool;default:false" json:"deleted"`
	DeletedBy  int32      `gorm:"type:int" json:"deleted_by"`
}
