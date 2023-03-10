// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/jwttoken"
	"license-manager/pkg/repositories/ent-fw/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JwtTokenCreate is the builder for creating a JwtToken entity.
type JwtTokenCreate struct {
	config
	mutation *JwtTokenMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (jtc *JwtTokenCreate) SetToken(s string) *JwtTokenCreate {
	jtc.mutation.SetToken(s)
	return jtc
}

// SetRevoked sets the "revoked" field.
func (jtc *JwtTokenCreate) SetRevoked(b bool) *JwtTokenCreate {
	jtc.mutation.SetRevoked(b)
	return jtc
}

// SetNillableRevoked sets the "revoked" field if the given value is not nil.
func (jtc *JwtTokenCreate) SetNillableRevoked(b *bool) *JwtTokenCreate {
	if b != nil {
		jtc.SetRevoked(*b)
	}
	return jtc
}

// SetClaims sets the "claims" field.
func (jtc *JwtTokenCreate) SetClaims(m map[string]interface{}) *JwtTokenCreate {
	jtc.mutation.SetClaims(m)
	return jtc
}

// SetIssuerID sets the "issuer_id" field.
func (jtc *JwtTokenCreate) SetIssuerID(s string) *JwtTokenCreate {
	jtc.mutation.SetIssuerID(s)
	return jtc
}

// SetID sets the "id" field.
func (jtc *JwtTokenCreate) SetID(s string) *JwtTokenCreate {
	jtc.mutation.SetID(s)
	return jtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (jtc *JwtTokenCreate) SetNillableID(s *string) *JwtTokenCreate {
	if s != nil {
		jtc.SetID(*s)
	}
	return jtc
}

// SetIssuer sets the "issuer" edge to the User entity.
func (jtc *JwtTokenCreate) SetIssuer(u *User) *JwtTokenCreate {
	return jtc.SetIssuerID(u.ID)
}

// Mutation returns the JwtTokenMutation object of the builder.
func (jtc *JwtTokenCreate) Mutation() *JwtTokenMutation {
	return jtc.mutation
}

// Save creates the JwtToken in the database.
func (jtc *JwtTokenCreate) Save(ctx context.Context) (*JwtToken, error) {
	jtc.defaults()
	return withHooks[*JwtToken, JwtTokenMutation](ctx, jtc.sqlSave, jtc.mutation, jtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jtc *JwtTokenCreate) SaveX(ctx context.Context) *JwtToken {
	v, err := jtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jtc *JwtTokenCreate) Exec(ctx context.Context) error {
	_, err := jtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jtc *JwtTokenCreate) ExecX(ctx context.Context) {
	if err := jtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jtc *JwtTokenCreate) defaults() {
	if _, ok := jtc.mutation.Revoked(); !ok {
		v := jwttoken.DefaultRevoked
		jtc.mutation.SetRevoked(v)
	}
	if _, ok := jtc.mutation.ID(); !ok {
		v := jwttoken.DefaultID()
		jtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jtc *JwtTokenCreate) check() error {
	if _, ok := jtc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "JwtToken.token"`)}
	}
	if v, ok := jtc.mutation.Token(); ok {
		if err := jwttoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "JwtToken.token": %w`, err)}
		}
	}
	if _, ok := jtc.mutation.Revoked(); !ok {
		return &ValidationError{Name: "revoked", err: errors.New(`ent: missing required field "JwtToken.revoked"`)}
	}
	if _, ok := jtc.mutation.Claims(); !ok {
		return &ValidationError{Name: "claims", err: errors.New(`ent: missing required field "JwtToken.claims"`)}
	}
	if _, ok := jtc.mutation.IssuerID(); !ok {
		return &ValidationError{Name: "issuer_id", err: errors.New(`ent: missing required field "JwtToken.issuer_id"`)}
	}
	if _, ok := jtc.mutation.IssuerID(); !ok {
		return &ValidationError{Name: "issuer", err: errors.New(`ent: missing required edge "JwtToken.issuer"`)}
	}
	return nil
}

func (jtc *JwtTokenCreate) sqlSave(ctx context.Context) (*JwtToken, error) {
	if err := jtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected JwtToken.ID type: %T", _spec.ID.Value)
		}
	}
	jtc.mutation.id = &_node.ID
	jtc.mutation.done = true
	return _node, nil
}

func (jtc *JwtTokenCreate) createSpec() (*JwtToken, *sqlgraph.CreateSpec) {
	var (
		_node = &JwtToken{config: jtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: jwttoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: jwttoken.FieldID,
			},
		}
	)
	if id, ok := jtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jtc.mutation.Token(); ok {
		_spec.SetField(jwttoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := jtc.mutation.Revoked(); ok {
		_spec.SetField(jwttoken.FieldRevoked, field.TypeBool, value)
		_node.Revoked = value
	}
	if value, ok := jtc.mutation.Claims(); ok {
		_spec.SetField(jwttoken.FieldClaims, field.TypeJSON, value)
		_node.Claims = value
	}
	if nodes := jtc.mutation.IssuerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jwttoken.IssuerTable,
			Columns: []string{jwttoken.IssuerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.IssuerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JwtTokenCreateBulk is the builder for creating many JwtToken entities in bulk.
type JwtTokenCreateBulk struct {
	config
	builders []*JwtTokenCreate
}

// Save creates the JwtToken entities in the database.
func (jtcb *JwtTokenCreateBulk) Save(ctx context.Context) ([]*JwtToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jtcb.builders))
	nodes := make([]*JwtToken, len(jtcb.builders))
	mutators := make([]Mutator, len(jtcb.builders))
	for i := range jtcb.builders {
		func(i int, root context.Context) {
			builder := jtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JwtTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, jtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jtcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, jtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jtcb *JwtTokenCreateBulk) SaveX(ctx context.Context) []*JwtToken {
	v, err := jtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jtcb *JwtTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := jtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jtcb *JwtTokenCreateBulk) ExecX(ctx context.Context) {
	if err := jtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
