package common

type DbInfo struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewDbInfo(host string, port string, user string, password string, name string) *DbInfo {

	if host == "" {
		panic("Host is not configured")
	}

	return &DbInfo{Host: host, User: user, Password: password, Name: name, Port: port}
}
