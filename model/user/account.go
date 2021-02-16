package user


type Account struct {
	Username		string	`json:"username" form:"username" binding:"required"`
	Password		string	`json:"password" form:"password" binding:"required"`
}

func NewAccount() *Account {
	return new(Account)
}
