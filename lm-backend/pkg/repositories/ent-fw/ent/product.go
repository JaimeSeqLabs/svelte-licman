// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/product"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// InstallInstr holds the value of the "install_instr" field.
	InstallInstr string `json:"install_instr,omitempty"`
	// LicenseCount holds the value of the "license_count" field.
	LicenseCount int `json:"license_count,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// LastUpdated holds the value of the "last_updated" field.
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges ProductEdges `json:"edges"`
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// License holds the value of the license edge.
	License []*License `json:"license,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LicenseOrErr returns the License value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) LicenseOrErr() ([]*License, error) {
	if e.loadedTypes[0] {
		return e.License, nil
	}
	return nil, &NotLoadedError{edge: "license"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldLicenseCount:
			values[i] = new(sql.NullInt64)
		case product.FieldID, product.FieldSku, product.FieldName, product.FieldInstallInstr:
			values[i] = new(sql.NullString)
		case product.FieldDateCreated, product.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case product.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				pr.Sku = value.String
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldInstallInstr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field install_instr", values[i])
			} else if value.Valid {
				pr.InstallInstr = value.String
			}
		case product.FieldLicenseCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field license_count", values[i])
			} else if value.Valid {
				pr.LicenseCount = int(value.Int64)
			}
		case product.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				pr.DateCreated = value.Time
			}
		case product.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated", values[i])
			} else if value.Valid {
				pr.LastUpdated = value.Time
			}
		}
	}
	return nil
}

// QueryLicense queries the "license" edge of the Product entity.
func (pr *Product) QueryLicense() *LicenseQuery {
	return NewProductClient(pr.config).QueryLicense(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("sku=")
	builder.WriteString(pr.Sku)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("install_instr=")
	builder.WriteString(pr.InstallInstr)
	builder.WriteString(", ")
	builder.WriteString("license_count=")
	builder.WriteString(fmt.Sprintf("%v", pr.LicenseCount))
	builder.WriteString(", ")
	builder.WriteString("date_created=")
	builder.WriteString(pr.DateCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_updated=")
	builder.WriteString(pr.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
