package entities

type Users struct {
	Name string `json:"name" gorm:"column:username"`
}
