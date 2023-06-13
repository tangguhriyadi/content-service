package dto

import (
	"github.com/tangguhriyadi/content-service/internals/entity"
)

type Content struct {
	ID           int32  `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	LikeCount    int32  `gorm:"type:int" json:"like_count"`
	CommentCount int32  `gorm:"type:int" json:"comment_count"`
	OwnerID      int32  `gorm:"type:int" json:"owner_id"`
	TypeID       int32  `gorm:"type:int" json:"type_id"`
	IsPremium    bool   `gorm:"type:bool; default:false" json:"is_premium"`
}

type ContentPaginate struct {
	Data       *[]entity.Content
	Page       int
	Limit      int
	TotalItems int64
}

type ContentPayload struct {
	Name      string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	IsPremium bool   `gorm:"type:bool; default:false" json:"is_premium"`
}

type ContentCreate struct {
	Name      string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	IsPremium bool   `gorm:"type:bool; default:false" json:"is_premium"`
	OwnerID   int32  `gorm:"type:int" json:"owner_id"`
}

type User struct {
	FullName string `gorm:"type:varchar(300)" json:"full_name" `
	Email    string `gorm:"type:varchar(300)" json:"email" `
	Age      int32  `gorm:"type:int" json:"age"`
}

type NewContent struct {
	ID           int32  `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	LikeCount    int32  `gorm:"type:int" json:"like_count"`
	CommentCount int32  `gorm:"type:int" json:"comment_count"`
	Owner        *User  `json:"owner"`
	TypeID       int32  `gorm:"type:int" json:"type_id"`
	IsPremium    bool   `gorm:"type:bool; default:false" json:"is_premium"`
}
