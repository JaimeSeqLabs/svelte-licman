// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/license"
	"license-manager/pkg/repositories/ent-fw/ent/organization"
	"license-manager/pkg/repositories/ent-fw/ent/product"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LicenseCreate is the builder for creating a License entity.
type LicenseCreate struct {
	config
	mutation *LicenseMutation
	hooks    []Hook
}

// SetFeatures sets the "features" field.
func (lc *LicenseCreate) SetFeatures(s string) *LicenseCreate {
	lc.mutation.SetFeatures(s)
	return lc
}

// SetStatus sets the "status" field.
func (lc *LicenseCreate) SetStatus(s string) *LicenseCreate {
	lc.mutation.SetStatus(s)
	return lc
}

// SetVersion sets the "version" field.
func (lc *LicenseCreate) SetVersion(s string) *LicenseCreate {
	lc.mutation.SetVersion(s)
	return lc
}

// SetNote sets the "note" field.
func (lc *LicenseCreate) SetNote(s string) *LicenseCreate {
	lc.mutation.SetNote(s)
	return lc
}

// SetContact sets the "contact" field.
func (lc *LicenseCreate) SetContact(s string) *LicenseCreate {
	lc.mutation.SetContact(s)
	return lc
}

// SetMail sets the "mail" field.
func (lc *LicenseCreate) SetMail(s string) *LicenseCreate {
	lc.mutation.SetMail(s)
	return lc
}

// SetSecret sets the "secret" field.
func (lc *LicenseCreate) SetSecret(s string) *LicenseCreate {
	lc.mutation.SetSecret(s)
	return lc
}

// SetExpirationDate sets the "expiration_date" field.
func (lc *LicenseCreate) SetExpirationDate(t time.Time) *LicenseCreate {
	lc.mutation.SetExpirationDate(t)
	return lc
}

// SetActivationDate sets the "activation_date" field.
func (lc *LicenseCreate) SetActivationDate(t time.Time) *LicenseCreate {
	lc.mutation.SetActivationDate(t)
	return lc
}

// SetLastAccessed sets the "last_accessed" field.
func (lc *LicenseCreate) SetLastAccessed(t time.Time) *LicenseCreate {
	lc.mutation.SetLastAccessed(t)
	return lc
}

// SetNillableLastAccessed sets the "last_accessed" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableLastAccessed(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetLastAccessed(*t)
	}
	return lc
}

// SetLastAccessIP sets the "last_access_IP" field.
func (lc *LicenseCreate) SetLastAccessIP(s string) *LicenseCreate {
	lc.mutation.SetLastAccessIP(s)
	return lc
}

// SetNillableLastAccessIP sets the "last_access_IP" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableLastAccessIP(s *string) *LicenseCreate {
	if s != nil {
		lc.SetLastAccessIP(*s)
	}
	return lc
}

// SetAccessCount sets the "access_count" field.
func (lc *LicenseCreate) SetAccessCount(i int) *LicenseCreate {
	lc.mutation.SetAccessCount(i)
	return lc
}

// SetNillableAccessCount sets the "access_count" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableAccessCount(i *int) *LicenseCreate {
	if i != nil {
		lc.SetAccessCount(*i)
	}
	return lc
}

// SetDateCreated sets the "date_created" field.
func (lc *LicenseCreate) SetDateCreated(t time.Time) *LicenseCreate {
	lc.mutation.SetDateCreated(t)
	return lc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableDateCreated(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetDateCreated(*t)
	}
	return lc
}

// SetLastUpdated sets the "last_updated" field.
func (lc *LicenseCreate) SetLastUpdated(t time.Time) *LicenseCreate {
	lc.mutation.SetLastUpdated(t)
	return lc
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableLastUpdated(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetLastUpdated(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LicenseCreate) SetID(s string) *LicenseCreate {
	lc.mutation.SetID(s)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableID(s *string) *LicenseCreate {
	if s != nil {
		lc.SetID(*s)
	}
	return lc
}

// AddLicenseProductIDs adds the "license_products" edge to the Product entity by IDs.
func (lc *LicenseCreate) AddLicenseProductIDs(ids ...string) *LicenseCreate {
	lc.mutation.AddLicenseProductIDs(ids...)
	return lc
}

// AddLicenseProducts adds the "license_products" edges to the Product entity.
func (lc *LicenseCreate) AddLicenseProducts(p ...*Product) *LicenseCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lc.AddLicenseProductIDs(ids...)
}

// SetOwnerOrgID sets the "owner_org" edge to the Organization entity by ID.
func (lc *LicenseCreate) SetOwnerOrgID(id string) *LicenseCreate {
	lc.mutation.SetOwnerOrgID(id)
	return lc
}

// SetNillableOwnerOrgID sets the "owner_org" edge to the Organization entity by ID if the given value is not nil.
func (lc *LicenseCreate) SetNillableOwnerOrgID(id *string) *LicenseCreate {
	if id != nil {
		lc = lc.SetOwnerOrgID(*id)
	}
	return lc
}

// SetOwnerOrg sets the "owner_org" edge to the Organization entity.
func (lc *LicenseCreate) SetOwnerOrg(o *Organization) *LicenseCreate {
	return lc.SetOwnerOrgID(o.ID)
}

// Mutation returns the LicenseMutation object of the builder.
func (lc *LicenseCreate) Mutation() *LicenseMutation {
	return lc.mutation
}

// Save creates the License in the database.
func (lc *LicenseCreate) Save(ctx context.Context) (*License, error) {
	lc.defaults()
	return withHooks[*License, LicenseMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LicenseCreate) SaveX(ctx context.Context) *License {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LicenseCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LicenseCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LicenseCreate) defaults() {
	if _, ok := lc.mutation.AccessCount(); !ok {
		v := license.DefaultAccessCount
		lc.mutation.SetAccessCount(v)
	}
	if _, ok := lc.mutation.DateCreated(); !ok {
		v := license.DefaultDateCreated()
		lc.mutation.SetDateCreated(v)
	}
	if _, ok := lc.mutation.LastUpdated(); !ok {
		v := license.DefaultLastUpdated()
		lc.mutation.SetLastUpdated(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := license.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LicenseCreate) check() error {
	if _, ok := lc.mutation.Features(); !ok {
		return &ValidationError{Name: "features", err: errors.New(`ent: missing required field "License.features"`)}
	}
	if _, ok := lc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "License.status"`)}
	}
	if _, ok := lc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "License.version"`)}
	}
	if _, ok := lc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "License.note"`)}
	}
	if _, ok := lc.mutation.Contact(); !ok {
		return &ValidationError{Name: "contact", err: errors.New(`ent: missing required field "License.contact"`)}
	}
	if _, ok := lc.mutation.Mail(); !ok {
		return &ValidationError{Name: "mail", err: errors.New(`ent: missing required field "License.mail"`)}
	}
	if _, ok := lc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`ent: missing required field "License.secret"`)}
	}
	if _, ok := lc.mutation.ExpirationDate(); !ok {
		return &ValidationError{Name: "expiration_date", err: errors.New(`ent: missing required field "License.expiration_date"`)}
	}
	if _, ok := lc.mutation.ActivationDate(); !ok {
		return &ValidationError{Name: "activation_date", err: errors.New(`ent: missing required field "License.activation_date"`)}
	}
	if _, ok := lc.mutation.AccessCount(); !ok {
		return &ValidationError{Name: "access_count", err: errors.New(`ent: missing required field "License.access_count"`)}
	}
	if _, ok := lc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "License.date_created"`)}
	}
	if _, ok := lc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New(`ent: missing required field "License.last_updated"`)}
	}
	if len(lc.mutation.LicenseProductsIDs()) == 0 {
		return &ValidationError{Name: "license_products", err: errors.New(`ent: missing required edge "License.license_products"`)}
	}
	return nil
}

func (lc *LicenseCreate) sqlSave(ctx context.Context) (*License, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected License.ID type: %T", _spec.ID.Value)
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LicenseCreate) createSpec() (*License, *sqlgraph.CreateSpec) {
	var (
		_node = &License{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: license.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: license.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Features(); ok {
		_spec.SetField(license.FieldFeatures, field.TypeString, value)
		_node.Features = value
	}
	if value, ok := lc.mutation.Status(); ok {
		_spec.SetField(license.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := lc.mutation.Version(); ok {
		_spec.SetField(license.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := lc.mutation.Note(); ok {
		_spec.SetField(license.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := lc.mutation.Contact(); ok {
		_spec.SetField(license.FieldContact, field.TypeString, value)
		_node.Contact = value
	}
	if value, ok := lc.mutation.Mail(); ok {
		_spec.SetField(license.FieldMail, field.TypeString, value)
		_node.Mail = value
	}
	if value, ok := lc.mutation.Secret(); ok {
		_spec.SetField(license.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := lc.mutation.ExpirationDate(); ok {
		_spec.SetField(license.FieldExpirationDate, field.TypeTime, value)
		_node.ExpirationDate = value
	}
	if value, ok := lc.mutation.ActivationDate(); ok {
		_spec.SetField(license.FieldActivationDate, field.TypeTime, value)
		_node.ActivationDate = value
	}
	if value, ok := lc.mutation.LastAccessed(); ok {
		_spec.SetField(license.FieldLastAccessed, field.TypeTime, value)
		_node.LastAccessed = value
	}
	if value, ok := lc.mutation.LastAccessIP(); ok {
		_spec.SetField(license.FieldLastAccessIP, field.TypeString, value)
		_node.LastAccessIP = value
	}
	if value, ok := lc.mutation.AccessCount(); ok {
		_spec.SetField(license.FieldAccessCount, field.TypeInt, value)
		_node.AccessCount = value
	}
	if value, ok := lc.mutation.DateCreated(); ok {
		_spec.SetField(license.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := lc.mutation.LastUpdated(); ok {
		_spec.SetField(license.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := lc.mutation.LicenseProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   license.LicenseProductsTable,
			Columns: license.LicenseProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.OwnerOrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   license.OwnerOrgTable,
			Columns: []string{license.OwnerOrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_licenses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LicenseCreateBulk is the builder for creating many License entities in bulk.
type LicenseCreateBulk struct {
	config
	builders []*LicenseCreate
}

// Save creates the License entities in the database.
func (lcb *LicenseCreateBulk) Save(ctx context.Context) ([]*License, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*License, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LicenseMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LicenseCreateBulk) SaveX(ctx context.Context) []*License {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LicenseCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LicenseCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
