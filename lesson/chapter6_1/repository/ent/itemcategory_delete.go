// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-web/lesson/chapter6_1/repository/ent/itemcategory"
	"go-web/lesson/chapter6_1/repository/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ItemCategoryDelete is the builder for deleting a ItemCategory entity.
type ItemCategoryDelete struct {
	config
	hooks    []Hook
	mutation *ItemCategoryMutation
}

// Where appends a list predicates to the ItemCategoryDelete builder.
func (icd *ItemCategoryDelete) Where(ps ...predicate.ItemCategory) *ItemCategoryDelete {
	icd.mutation.Where(ps...)
	return icd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (icd *ItemCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(icd.hooks) == 0 {
		affected, err = icd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			icd.mutation = mutation
			affected, err = icd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(icd.hooks) - 1; i >= 0; i-- {
			if icd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, icd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (icd *ItemCategoryDelete) ExecX(ctx context.Context) int {
	n, err := icd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (icd *ItemCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: itemcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: itemcategory.FieldID,
			},
		},
	}
	if ps := icd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, icd.driver, _spec)
}

// ItemCategoryDeleteOne is the builder for deleting a single ItemCategory entity.
type ItemCategoryDeleteOne struct {
	icd *ItemCategoryDelete
}

// Exec executes the deletion query.
func (icdo *ItemCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := icdo.icd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{itemcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (icdo *ItemCategoryDeleteOne) ExecX(ctx context.Context) {
	icdo.icd.ExecX(ctx)
}
