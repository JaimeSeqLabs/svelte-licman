// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"license-manager/pkg/repositories/ent-fw/ent/migrate"

	"license-manager/pkg/repositories/ent-fw/ent/contact"
	"license-manager/pkg/repositories/ent-fw/ent/organization"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Contact is the client for interacting with the Contact builders.
	Contact *ContactClient
	// Organization is the client for interacting with the Organization builders.
	Organization *OrganizationClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Contact = NewContactClient(c.config)
	c.Organization = NewOrganizationClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Contact:      NewContactClient(cfg),
		Organization: NewOrganizationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Contact:      NewContactClient(cfg),
		Organization: NewOrganizationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Contact.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Contact.Use(hooks...)
	c.Organization.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Contact.Intercept(interceptors...)
	c.Organization.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ContactMutation:
		return c.Contact.mutate(ctx, m)
	case *OrganizationMutation:
		return c.Organization.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ContactClient is a client for the Contact schema.
type ContactClient struct {
	config
}

// NewContactClient returns a client for the Contact from the given config.
func NewContactClient(c config) *ContactClient {
	return &ContactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contact.Hooks(f(g(h())))`.
func (c *ContactClient) Use(hooks ...Hook) {
	c.hooks.Contact = append(c.hooks.Contact, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `contact.Intercept(f(g(h())))`.
func (c *ContactClient) Intercept(interceptors ...Interceptor) {
	c.inters.Contact = append(c.inters.Contact, interceptors...)
}

// Create returns a builder for creating a Contact entity.
func (c *ContactClient) Create() *ContactCreate {
	mutation := newContactMutation(c.config, OpCreate)
	return &ContactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contact entities.
func (c *ContactClient) CreateBulk(builders ...*ContactCreate) *ContactCreateBulk {
	return &ContactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contact.
func (c *ContactClient) Update() *ContactUpdate {
	mutation := newContactMutation(c.config, OpUpdate)
	return &ContactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContactClient) UpdateOne(co *Contact) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContact(co))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContactClient) UpdateOneID(id int) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContactID(id))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contact.
func (c *ContactClient) Delete() *ContactDelete {
	mutation := newContactMutation(c.config, OpDelete)
	return &ContactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContactClient) DeleteOne(co *Contact) *ContactDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ContactClient) DeleteOneID(id int) *ContactDeleteOne {
	builder := c.Delete().Where(contact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContactDeleteOne{builder}
}

// Query returns a query builder for Contact.
func (c *ContactClient) Query() *ContactQuery {
	return &ContactQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeContact},
		inters: c.Interceptors(),
	}
}

// Get returns a Contact entity by its id.
func (c *ContactClient) Get(ctx context.Context, id int) (*Contact, error) {
	return c.Query().Where(contact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContactClient) GetX(ctx context.Context, id int) *Contact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ContactClient) Hooks() []Hook {
	return c.hooks.Contact
}

// Interceptors returns the client interceptors.
func (c *ContactClient) Interceptors() []Interceptor {
	return c.inters.Contact
}

func (c *ContactClient) mutate(ctx context.Context, m *ContactMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ContactCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ContactUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ContactDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Contact mutation op: %q", m.Op())
	}
}

// OrganizationClient is a client for the Organization schema.
type OrganizationClient struct {
	config
}

// NewOrganizationClient returns a client for the Organization from the given config.
func NewOrganizationClient(c config) *OrganizationClient {
	return &OrganizationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `organization.Hooks(f(g(h())))`.
func (c *OrganizationClient) Use(hooks ...Hook) {
	c.hooks.Organization = append(c.hooks.Organization, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `organization.Intercept(f(g(h())))`.
func (c *OrganizationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Organization = append(c.inters.Organization, interceptors...)
}

// Create returns a builder for creating a Organization entity.
func (c *OrganizationClient) Create() *OrganizationCreate {
	mutation := newOrganizationMutation(c.config, OpCreate)
	return &OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Organization entities.
func (c *OrganizationClient) CreateBulk(builders ...*OrganizationCreate) *OrganizationCreateBulk {
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Organization.
func (c *OrganizationClient) Update() *OrganizationUpdate {
	mutation := newOrganizationMutation(c.config, OpUpdate)
	return &OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrganizationClient) UpdateOne(o *Organization) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganization(o))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrganizationClient) UpdateOneID(id int) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganizationID(id))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Organization.
func (c *OrganizationClient) Delete() *OrganizationDelete {
	mutation := newOrganizationMutation(c.config, OpDelete)
	return &OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrganizationClient) DeleteOne(o *Organization) *OrganizationDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrganizationClient) DeleteOneID(id int) *OrganizationDeleteOne {
	builder := c.Delete().Where(organization.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrganizationDeleteOne{builder}
}

// Query returns a query builder for Organization.
func (c *OrganizationClient) Query() *OrganizationQuery {
	return &OrganizationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrganization},
		inters: c.Interceptors(),
	}
}

// Get returns a Organization entity by its id.
func (c *OrganizationClient) Get(ctx context.Context, id int) (*Organization, error) {
	return c.Query().Where(organization.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrganizationClient) GetX(ctx context.Context, id int) *Organization {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryContact queries the contact edge of a Organization.
func (c *OrganizationClient) QueryContact(o *Organization) *ContactQuery {
	query := (&ContactClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(contact.Table, contact.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, organization.ContactTable, organization.ContactColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrganizationClient) Hooks() []Hook {
	return c.hooks.Organization
}

// Interceptors returns the client interceptors.
func (c *OrganizationClient) Interceptors() []Interceptor {
	return c.inters.Organization
}

func (c *OrganizationClient) mutate(ctx context.Context, m *OrganizationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Organization mutation op: %q", m.Op())
	}
}
