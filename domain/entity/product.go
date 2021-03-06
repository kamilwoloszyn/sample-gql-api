package entity

type Product struct {
	ID                string
	Name              string
	Price             float32
	Color             string
	Gtin              string
	Model             string
	CountryOfAssembly string
	Category          Category
	SKU               string
	CreatedAt         int64
	UpdatedAt         int64
	DeletedAt         int64
}

func (p *Product) IsValid() bool {
	if p.Name != "" &&
		p.Gtin != "" &&
		p.SKU != "" {
		return true
	}
	return false
}

func (p *Product) GetID() string {
	return p.ID
}
