package entity

type Category struct {
	Id   int    `db:"id"`
	Name string `db:"nome"`
}

func (c *Category) IsEmpty() bool {
	return c == nil || (c.Id == 0 && c.Name == "")
}
