package entity

import "time"

type ContentComment struct {
	ID        int32     `gorm:"primaryKey" json:"id"`
	ContentId int32     `gorm:"type:int" json:"content_id"`
	UserId    int32     `gorm:"type:int" json:"user_id"`
	LikeCount int32     `gorm:"type:int" json:"like_count"`
	Comment   string    `gorm:"type:varchar" json:"comment"`
	ReplyTo   int32     `gorm:"type:int; default:null" json:"reply_to"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
