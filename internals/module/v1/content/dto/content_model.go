package dto

import "github.com/tangguhriyadi/content-service/internals/entity"

type Content struct {
	ID           int32  `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"type:varchar(300)" json:"name"`
	LikeCount    int32  `gorm:"type:int" json:"like_count"`
	CommentCount int32  `gorm:"type:int" json:"comment_count"`
	OwnerID      int32  `gorm:"type:int" json:"owner_id"`
	TypeID       int32  `gorm:"type:int" json:"type_id"`
	IsPremium    bool   `gorm:"type:bool" json:"is_premium"`
}

type ContentPaginate struct {
	Data       *[]entity.Content
	Page       int
	Limit      int
	TotalItems int32
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
