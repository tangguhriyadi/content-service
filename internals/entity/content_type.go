package entity

import "time"

type ContentType struct {
	ID        int32      `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(300)" json:"name"`
	CreatedAt time.Time  `gorm:"type:timestamp" json:"created_at"`
	CreatedBy int32      `gorm:"type:int" json:"created_by"`
	UpdatedAt time.Time  `gorm:"type:timestamp; default:null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deleted_at"`
	Deleted   bool       `gorm:"type:bool;default:false" json:"deleted"`
	DeletedBy int32      `gorm:"type:int" json:"deleted_by"`
}
