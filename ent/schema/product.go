package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.Float("price").GoType(decimal.Decimal{}),
		field.Int("quantity").Default(0),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
