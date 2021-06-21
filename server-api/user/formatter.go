package user

type UserFormat struct {
	FullName string
	Address  string
	Email    string
}

func FormatUser(fullname string, address string, email string) UserFormat {
	var formatUser = UserFormat{
		FullName: fullname,
		Address:  address,
		Email:    email,
	}
	return formatUser
}
