// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-web/lesson/chapter6_1/repository/ent/order"
	"go-web/lesson/chapter6_1/repository/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (oc *OrderCreate) SetCreateTime(t time.Time) *OrderCreate {
	oc.mutation.SetCreateTime(t)
	return oc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreateTime(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreateTime(*t)
	}
	return oc
}

// SetUpdateTime sets the "update_time" field.
func (oc *OrderCreate) SetUpdateTime(t time.Time) *OrderCreate {
	oc.mutation.SetUpdateTime(t)
	return oc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdateTime(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdateTime(*t)
	}
	return oc
}

// SetUserID sets the "user_id" field.
func (oc *OrderCreate) SetUserID(i int) *OrderCreate {
	oc.mutation.SetUserID(i)
	return oc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUserID(i *int) *OrderCreate {
	if i != nil {
		oc.SetUserID(*i)
	}
	return oc
}

// SetSn sets the "sn" field.
func (oc *OrderCreate) SetSn(s string) *OrderCreate {
	oc.mutation.SetSn(s)
	return oc
}

// SetUser sets the "user" edge to the User entity.
func (oc *OrderCreate) SetUser(u *User) *OrderCreate {
	return oc.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.CreateTime(); !ok {
		v := order.DefaultCreateTime()
		oc.mutation.SetCreateTime(v)
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		v := order.DefaultUpdateTime()
		oc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := oc.mutation.Sn(); !ok {
		return &ValidationError{Name: "sn", err: errors.New(`ent: missing required field "sn"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := oc.mutation.Sn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldSn,
		})
		_node.Sn = value
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
