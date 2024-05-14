package database




type User struct {
	ID  string `gorm:"id ;primaryKey" json:"id"`
	Name string `gorm:"name ; not null" json:"name"`
	Email string `gorm:"email ;unique;not null" json:"email"`
	Password string `gorm:"password;not null" json:"password"`
}

type Session struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Token string `json:"token"`
	
	ExpiresAt string `json:"expires_at"`
}


type Otp struct {
	ID string `json:"id"`
	Email string `json:"email"`
	OTP string `json:"otp"`
	ExpiresAt string `json:"expires_at"`
}

