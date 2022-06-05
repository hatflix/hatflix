package entity

type Store struct {
	Id          int    `db:"id"`
	Name        string `db:"nome"`
	Cnpj        string `db:"cnpj"`
	PhoneNumber string `db:"telegone"`
	Address     string `db:"endereco"`
	CategoryID  int    `db:"id_categoria"`
}

func (d *Store) IsEmpty() bool {
	return d == nil || (d.Id == 0 && d.CategoryID == 0)
}
