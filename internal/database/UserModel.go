package database




type User struct {
	ID  string `gorm:"id ;primaryKey" json:"id"`
	Name string `gorm:"name ; not null" json:"name"`
	Email string `gorm:"email ;unique;not null" json:"email"`
	Password string `gorm:"password;not null" json:"password"`
}
