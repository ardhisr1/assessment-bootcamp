package user

type UserFormat struct {
	ID       int
	FullName string
	Address  string
	Email    string
}

func FormatUser(id int, fullname string, address string, email string) UserFormat {
	var formatUser = UserFormat{
		ID:       id,
		FullName: fullname,
		Address:  address,
		Email:    email,
	}
	return formatUser
}
