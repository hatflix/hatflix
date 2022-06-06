package entity

type User struct {
	Id          int    `db:"id"`
	FirstName   string `db:"primeiro_nome"`
	LastName    string `db:"ultimo_nome"`
	Email       string `db:"email"`
	PhoneNumber string `db:"telefone"`
	Password    string `db:"senha_hash"`
}

func (u *User) IsEmpty() bool {
	return u == nil || u.Id == 0
}
