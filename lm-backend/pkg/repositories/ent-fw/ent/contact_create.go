// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/contact"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContactCreate is the builder for creating a Contact entity.
type ContactCreate struct {
	config
	mutation *ContactMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ContactCreate) SetName(s string) *ContactCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetMail sets the "mail" field.
func (cc *ContactCreate) SetMail(s string) *ContactCreate {
	cc.mutation.SetMail(s)
	return cc
}

// Mutation returns the ContactMutation object of the builder.
func (cc *ContactCreate) Mutation() *ContactMutation {
	return cc.mutation
}

// Save creates the Contact in the database.
func (cc *ContactCreate) Save(ctx context.Context) (*Contact, error) {
	return withHooks[*Contact, ContactMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContactCreate) SaveX(ctx context.Context) *Contact {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContactCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContactCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContactCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Contact.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := contact.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Contact.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Mail(); !ok {
		return &ValidationError{Name: "mail", err: errors.New(`ent: missing required field "Contact.mail"`)}
	}
	if v, ok := cc.mutation.Mail(); ok {
		if err := contact.MailValidator(v); err != nil {
			return &ValidationError{Name: "mail", err: fmt.Errorf(`ent: validator failed for field "Contact.mail": %w`, err)}
		}
	}
	return nil
}

func (cc *ContactCreate) sqlSave(ctx context.Context) (*Contact, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ContactCreate) createSpec() (*Contact, *sqlgraph.CreateSpec) {
	var (
		_node = &Contact{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: contact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contact.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(contact.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Mail(); ok {
		_spec.SetField(contact.FieldMail, field.TypeString, value)
		_node.Mail = value
	}
	return _node, _spec
}

// ContactCreateBulk is the builder for creating many Contact entities in bulk.
type ContactCreateBulk struct {
	config
	builders []*ContactCreate
}

// Save creates the Contact entities in the database.
func (ccb *ContactCreateBulk) Save(ctx context.Context) ([]*Contact, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contact, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContactCreateBulk) SaveX(ctx context.Context) []*Contact {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContactCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContactCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
