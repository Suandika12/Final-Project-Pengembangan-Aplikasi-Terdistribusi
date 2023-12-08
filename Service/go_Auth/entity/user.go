package entity

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increament" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
	Nik      string `json:"nik" form:"nik" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}
