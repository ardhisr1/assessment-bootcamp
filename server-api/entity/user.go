package entity

type User struct {
	ID       int    `gorm:"primaryKey" json:"user_id"`
	FullName string `json:"fullname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//PasswordItems []PasswordItem `gorm:"foreignKey:UserID`
}

type UserInput struct {
	FullName string `json:"fullname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateInput struct {
	FullName string `json:"fullname"`
	Address  string `json:"address"`
}
