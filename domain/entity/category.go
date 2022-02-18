package entity

type Category struct {
	ID           string
	CategoryName string
	CreatedAt    int64
	UpdatedAt    int64
	DeletedAt    int64
}

func (c *Category) Validate() error {
	return nil
}
func (c *Category) GetID() string {
	return c.ID
}
