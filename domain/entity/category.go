package entity

type Category struct {
	ID           string
	CategoryName string
	CreatedAt    int64
	UpdatedAt    int64
	DeletedAt    int64
}

func (c *Category) IsValid() bool {
	return c.CategoryName != ""
}
func (c *Category) GetID() string {
	return c.ID
}
