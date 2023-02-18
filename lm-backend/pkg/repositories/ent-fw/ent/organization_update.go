// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/contact"
	"license-manager/pkg/repositories/ent-fw/ent/organization"
	"license-manager/pkg/repositories/ent-fw/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OrganizationUpdate) SetName(s string) *OrganizationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetLocation sets the "location" field.
func (ou *OrganizationUpdate) SetLocation(s string) *OrganizationUpdate {
	ou.mutation.SetLocation(s)
	return ou
}

// SetContactID sets the "contact_id" field.
func (ou *OrganizationUpdate) SetContactID(s string) *OrganizationUpdate {
	ou.mutation.SetContactID(s)
	return ou
}

// SetNillableContactID sets the "contact_id" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableContactID(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetContactID(*s)
	}
	return ou
}

// ClearContactID clears the value of the "contact_id" field.
func (ou *OrganizationUpdate) ClearContactID() *OrganizationUpdate {
	ou.mutation.ClearContactID()
	return ou
}

// SetContact sets the "contact" edge to the Contact entity.
func (ou *OrganizationUpdate) SetContact(c *Contact) *OrganizationUpdate {
	return ou.SetContactID(c.ID)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// ClearContact clears the "contact" edge to the Contact entity.
func (ou *OrganizationUpdate) ClearContact() *OrganizationUpdate {
	ou.mutation.ClearContact()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrganizationUpdate) check() error {
	if v, ok := ou.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Location(); ok {
		if err := organization.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Organization.location": %w`, err)}
		}
	}
	return nil
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organization.Table,
			Columns: organization.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: organization.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Location(); ok {
		_spec.SetField(organization.FieldLocation, field.TypeString, value)
	}
	if ou.mutation.ContactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.ContactTable,
			Columns: []string{organization.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.ContactTable,
			Columns: []string{organization.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetName sets the "name" field.
func (ouo *OrganizationUpdateOne) SetName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetLocation sets the "location" field.
func (ouo *OrganizationUpdateOne) SetLocation(s string) *OrganizationUpdateOne {
	ouo.mutation.SetLocation(s)
	return ouo
}

// SetContactID sets the "contact_id" field.
func (ouo *OrganizationUpdateOne) SetContactID(s string) *OrganizationUpdateOne {
	ouo.mutation.SetContactID(s)
	return ouo
}

// SetNillableContactID sets the "contact_id" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableContactID(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetContactID(*s)
	}
	return ouo
}

// ClearContactID clears the value of the "contact_id" field.
func (ouo *OrganizationUpdateOne) ClearContactID() *OrganizationUpdateOne {
	ouo.mutation.ClearContactID()
	return ouo
}

// SetContact sets the "contact" edge to the Contact entity.
func (ouo *OrganizationUpdateOne) SetContact(c *Contact) *OrganizationUpdateOne {
	return ouo.SetContactID(c.ID)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// ClearContact clears the "contact" edge to the Contact entity.
func (ouo *OrganizationUpdateOne) ClearContact() *OrganizationUpdateOne {
	ouo.mutation.ClearContact()
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	return withHooks[*Organization, OrganizationMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrganizationUpdateOne) check() error {
	if v, ok := ouo.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Location(); ok {
		if err := organization.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Organization.location": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organization.Table,
			Columns: organization.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: organization.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Location(); ok {
		_spec.SetField(organization.FieldLocation, field.TypeString, value)
	}
	if ouo.mutation.ContactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.ContactTable,
			Columns: []string{organization.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.ContactTable,
			Columns: []string{organization.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
