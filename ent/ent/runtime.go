// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent-demo/ent/car"
	"ent-demo/ent/schema"
	"ent-demo/ent/todo"
	"ent-demo/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescCreateAt is the schema descriptor for create_at field.
	carDescCreateAt := carFields[3].Descriptor()
	// car.UpdateDefaultCreateAt holds the default value on update for the create_at field.
	car.UpdateDefaultCreateAt = carDescCreateAt.UpdateDefault.(func() time.Time)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescTitle is the schema descriptor for title field.
	todoDescTitle := todoFields[0].Descriptor()
	// todo.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	todo.TitleValidator = todoDescTitle.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[2].Descriptor()
	// user.DefaultAge holds the default value on creation for the age field.
	user.DefaultAge = userDescAge.Default.(uint8)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int64) error)
}
