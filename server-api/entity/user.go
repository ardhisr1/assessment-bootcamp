package entity

type User struct {
	ID            int            `gorm:"primaryKey" json:"user_id"`
	FullName      string         `json:"fullname"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	PasswordItems []PasswordItem `gorm:"foreignKey:UserID`
}
