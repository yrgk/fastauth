package storage

type (
	User struct {
		ID       uint   `json:"id" gorm:"primarykey"`
		Username string `json:"username" gorm:"unique; not null"`
		Email    string `json:"email" gorm:"unique; not null"`
		Password string `json:"password" gorm:"unique; not null"`
	}
)
