package store

// binding:"required" из библиотеки gin проверяет наличие поля
type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
