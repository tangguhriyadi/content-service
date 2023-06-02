package entity

import "time"

type Content struct {
	ID           int32      `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"type:varchar(300)" json:"name"`
	LikeCount    int32      `gorm:"type:int" json:"like_count"`
	CommentCount int32      `gorm:"type:int" json:"comment_count"`
	OwnerID      int32      `gorm:"type:int" json:"owner_id"`
	TypeID       int32      `gorm:"type:int" json:"type_id"`
	IsPremium    bool       `gorm:"type:bool;default:false" json:"is_premium"`
	CreatedAt    time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"type:timestamp" json:"deleted_at"`
	Deleted      bool       `gorm:"type:bool;default:false" json:"deleted"`
	DeletedBy    int32      `gorm:"type:int" json:"deleted_by"`
}
