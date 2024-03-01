package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("用户ID").Positive(),
		field.String("username").Unique().Comment("用户名"),
		field.Uint8("age").Default(18).Comment("年龄"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
		edge.To("cars", Car.Type),
		edge.From("groups", Group.Type).Ref("user"),
	}
}
