package entity

import "time"

type ContentLikeHistory struct {
	ID        int32     `gorm:"primaryKey" json:"id"`
	UserId    int32     `gorm:"type:int" json:"user_id"`
	ContentId int32     `gorm:"type:int" json:"content_id"`
	Type      string    `gorm:"type:varchar(300)" json:"type"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
