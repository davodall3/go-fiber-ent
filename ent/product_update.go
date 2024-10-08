// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"projectSwagger/ent/predicate"
	"projectSwagger/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableName(s *string) *ProductUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetPrice sets the "price" field.
func (pu *ProductUpdate) SetPrice(d decimal.Decimal) *ProductUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(d)
	return pu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pu *ProductUpdate) SetNillablePrice(d *decimal.Decimal) *ProductUpdate {
	if d != nil {
		pu.SetPrice(*d)
	}
	return pu
}

// AddPrice adds d to the "price" field.
func (pu *ProductUpdate) AddPrice(d decimal.Decimal) *ProductUpdate {
	pu.mutation.AddPrice(d)
	return pu
}

// SetQuantity sets the "quantity" field.
func (pu *ProductUpdate) SetQuantity(i int) *ProductUpdate {
	pu.mutation.ResetQuantity()
	pu.mutation.SetQuantity(i)
	return pu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableQuantity(i *int) *ProductUpdate {
	if i != nil {
		pu.SetQuantity(*i)
	}
	return pu
}

// AddQuantity adds i to the "quantity" field.
func (pu *ProductUpdate) AddQuantity(i int) *ProductUpdate {
	pu.mutation.AddQuantity(i)
	return pu
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Quantity(); ok {
		_spec.SetField(product.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedQuantity(); ok {
		_spec.AddField(product.FieldQuantity, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableName(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetPrice sets the "price" field.
func (puo *ProductUpdateOne) SetPrice(d decimal.Decimal) *ProductUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(d)
	return puo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillablePrice(d *decimal.Decimal) *ProductUpdateOne {
	if d != nil {
		puo.SetPrice(*d)
	}
	return puo
}

// AddPrice adds d to the "price" field.
func (puo *ProductUpdateOne) AddPrice(d decimal.Decimal) *ProductUpdateOne {
	puo.mutation.AddPrice(d)
	return puo
}

// SetQuantity sets the "quantity" field.
func (puo *ProductUpdateOne) SetQuantity(i int) *ProductUpdateOne {
	puo.mutation.ResetQuantity()
	puo.mutation.SetQuantity(i)
	return puo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableQuantity(i *int) *ProductUpdateOne {
	if i != nil {
		puo.SetQuantity(*i)
	}
	return puo
}

// AddQuantity adds i to the "quantity" field.
func (puo *ProductUpdateOne) AddQuantity(i int) *ProductUpdateOne {
	puo.mutation.AddQuantity(i)
	return puo
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (puo *ProductUpdateOne) Where(ps ...predicate.Product) *ProductUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Quantity(); ok {
		_spec.SetField(product.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedQuantity(); ok {
		_spec.AddField(product.FieldQuantity, field.TypeInt, value)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
