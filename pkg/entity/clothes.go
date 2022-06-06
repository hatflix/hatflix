package entity

type Clothes struct {
	Id         int     `db:"id"`
	Name       string  `db:"nome"`
	StoreID    int     `db:"id_loja"`
	CategoryID int     `db:"id_categoria"`
	Size       string  `db:"tamanho"`
	Price      float64 `db:"preco"`
	Quantity   int     `db:"quantidade"`
}

func (d *Clothes) IsEmpty() bool {
	return d == nil || (d.Id == 0 && d.CategoryID == 0)
}
