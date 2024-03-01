package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("car id"),
		field.String("name").Comment("car name"),
		field.Float32("price").Comment("car price"),
		field.Time("create_at").Comment("create time").UpdateDefault(time.Now),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host", User.Type).Ref("cars").Unique(),
	}
}
