package dto

type ContentTypePaginate struct {
	Data       *[]ContentType
	Page       int
	Limit      int
	TotalItems int64
}

type ContentType struct {
	ID   int32  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:int" json:"name"`
}
