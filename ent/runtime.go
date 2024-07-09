// Code generated by ent, DO NOT EDIT.

package ent

import (
	"projectSwagger/ent/product"
	"projectSwagger/ent/schema"
	"projectSwagger/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[0].Descriptor()
	// product.DefaultName holds the default value on creation for the name field.
	product.DefaultName = productDescName.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescSurname is the schema descriptor for surname field.
	userDescSurname := userFields[1].Descriptor()
	// user.DefaultSurname holds the default value on creation for the surname field.
	user.DefaultSurname = userDescSurname.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[4].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[5].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
}
