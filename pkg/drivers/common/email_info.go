package common

type EmailInfo struct {
	Smtp     string
	Port     string
	Email    string
	Password string
}

func NewEmailInfo(smpt string, port string, email string, password string) *EmailInfo {

	return &EmailInfo{
		Smtp:     smpt,
		Port:     port,
		Email:    email,
		Password: password,
	}
}
