// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/casbinrule"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CasbinRuleCreate is the builder for creating a CasbinRule entity.
type CasbinRuleCreate struct {
	config
	mutation *CasbinRuleMutation
	hooks    []Hook
}

// SetPtype sets the "Ptype" field.
func (crc *CasbinRuleCreate) SetPtype(s string) *CasbinRuleCreate {
	crc.mutation.SetPtype(s)
	return crc
}

// SetNillablePtype sets the "Ptype" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillablePtype(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetPtype(*s)
	}
	return crc
}

// SetV0 sets the "V0" field.
func (crc *CasbinRuleCreate) SetV0(s string) *CasbinRuleCreate {
	crc.mutation.SetV0(s)
	return crc
}

// SetNillableV0 sets the "V0" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillableV0(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetV0(*s)
	}
	return crc
}

// SetV1 sets the "V1" field.
func (crc *CasbinRuleCreate) SetV1(s string) *CasbinRuleCreate {
	crc.mutation.SetV1(s)
	return crc
}

// SetNillableV1 sets the "V1" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillableV1(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetV1(*s)
	}
	return crc
}

// SetV2 sets the "V2" field.
func (crc *CasbinRuleCreate) SetV2(s string) *CasbinRuleCreate {
	crc.mutation.SetV2(s)
	return crc
}

// SetNillableV2 sets the "V2" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillableV2(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetV2(*s)
	}
	return crc
}

// SetV3 sets the "V3" field.
func (crc *CasbinRuleCreate) SetV3(s string) *CasbinRuleCreate {
	crc.mutation.SetV3(s)
	return crc
}

// SetNillableV3 sets the "V3" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillableV3(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetV3(*s)
	}
	return crc
}

// SetV4 sets the "V4" field.
func (crc *CasbinRuleCreate) SetV4(s string) *CasbinRuleCreate {
	crc.mutation.SetV4(s)
	return crc
}

// SetNillableV4 sets the "V4" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillableV4(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetV4(*s)
	}
	return crc
}

// SetV5 sets the "V5" field.
func (crc *CasbinRuleCreate) SetV5(s string) *CasbinRuleCreate {
	crc.mutation.SetV5(s)
	return crc
}

// SetNillableV5 sets the "V5" field if the given value is not nil.
func (crc *CasbinRuleCreate) SetNillableV5(s *string) *CasbinRuleCreate {
	if s != nil {
		crc.SetV5(*s)
	}
	return crc
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (crc *CasbinRuleCreate) Mutation() *CasbinRuleMutation {
	return crc.mutation
}

// Save creates the CasbinRule in the database.
func (crc *CasbinRuleCreate) Save(ctx context.Context) (*CasbinRule, error) {
	crc.defaults()
	return withHooks(ctx, crc.sqlSave, crc.mutation, crc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (crc *CasbinRuleCreate) SaveX(ctx context.Context) *CasbinRule {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crc *CasbinRuleCreate) Exec(ctx context.Context) error {
	_, err := crc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crc *CasbinRuleCreate) ExecX(ctx context.Context) {
	if err := crc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crc *CasbinRuleCreate) defaults() {
	if _, ok := crc.mutation.Ptype(); !ok {
		v := casbinrule.DefaultPtype
		crc.mutation.SetPtype(v)
	}
	if _, ok := crc.mutation.V0(); !ok {
		v := casbinrule.DefaultV0
		crc.mutation.SetV0(v)
	}
	if _, ok := crc.mutation.V1(); !ok {
		v := casbinrule.DefaultV1
		crc.mutation.SetV1(v)
	}
	if _, ok := crc.mutation.V2(); !ok {
		v := casbinrule.DefaultV2
		crc.mutation.SetV2(v)
	}
	if _, ok := crc.mutation.V3(); !ok {
		v := casbinrule.DefaultV3
		crc.mutation.SetV3(v)
	}
	if _, ok := crc.mutation.V4(); !ok {
		v := casbinrule.DefaultV4
		crc.mutation.SetV4(v)
	}
	if _, ok := crc.mutation.V5(); !ok {
		v := casbinrule.DefaultV5
		crc.mutation.SetV5(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crc *CasbinRuleCreate) check() error {
	if _, ok := crc.mutation.Ptype(); !ok {
		return &ValidationError{Name: "Ptype", err: errors.New(`ent: missing required field "CasbinRule.Ptype"`)}
	}
	if _, ok := crc.mutation.V0(); !ok {
		return &ValidationError{Name: "V0", err: errors.New(`ent: missing required field "CasbinRule.V0"`)}
	}
	if _, ok := crc.mutation.V1(); !ok {
		return &ValidationError{Name: "V1", err: errors.New(`ent: missing required field "CasbinRule.V1"`)}
	}
	if _, ok := crc.mutation.V2(); !ok {
		return &ValidationError{Name: "V2", err: errors.New(`ent: missing required field "CasbinRule.V2"`)}
	}
	if _, ok := crc.mutation.V3(); !ok {
		return &ValidationError{Name: "V3", err: errors.New(`ent: missing required field "CasbinRule.V3"`)}
	}
	if _, ok := crc.mutation.V4(); !ok {
		return &ValidationError{Name: "V4", err: errors.New(`ent: missing required field "CasbinRule.V4"`)}
	}
	if _, ok := crc.mutation.V5(); !ok {
		return &ValidationError{Name: "V5", err: errors.New(`ent: missing required field "CasbinRule.V5"`)}
	}
	return nil
}

func (crc *CasbinRuleCreate) sqlSave(ctx context.Context) (*CasbinRule, error) {
	if err := crc.check(); err != nil {
		return nil, err
	}
	_node, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	crc.mutation.id = &_node.ID
	crc.mutation.done = true
	return _node, nil
}

func (crc *CasbinRuleCreate) createSpec() (*CasbinRule, *sqlgraph.CreateSpec) {
	var (
		_node = &CasbinRule{config: crc.config}
		_spec = sqlgraph.NewCreateSpec(casbinrule.Table, sqlgraph.NewFieldSpec(casbinrule.FieldID, field.TypeInt))
	)
	if value, ok := crc.mutation.Ptype(); ok {
		_spec.SetField(casbinrule.FieldPtype, field.TypeString, value)
		_node.Ptype = value
	}
	if value, ok := crc.mutation.V0(); ok {
		_spec.SetField(casbinrule.FieldV0, field.TypeString, value)
		_node.V0 = value
	}
	if value, ok := crc.mutation.V1(); ok {
		_spec.SetField(casbinrule.FieldV1, field.TypeString, value)
		_node.V1 = value
	}
	if value, ok := crc.mutation.V2(); ok {
		_spec.SetField(casbinrule.FieldV2, field.TypeString, value)
		_node.V2 = value
	}
	if value, ok := crc.mutation.V3(); ok {
		_spec.SetField(casbinrule.FieldV3, field.TypeString, value)
		_node.V3 = value
	}
	if value, ok := crc.mutation.V4(); ok {
		_spec.SetField(casbinrule.FieldV4, field.TypeString, value)
		_node.V4 = value
	}
	if value, ok := crc.mutation.V5(); ok {
		_spec.SetField(casbinrule.FieldV5, field.TypeString, value)
		_node.V5 = value
	}
	return _node, _spec
}

// SetCasbinRule set the CasbinRule
func (crc *CasbinRuleCreate) SetCasbinRule(input *CasbinRule, fields ...string) *CasbinRuleCreate {
	m := crc.mutation
	if len(fields) == 0 {
		fields = casbinrule.Columns
	}
	_ = m.SetFields(input, fields...)
	return crc
}

// SetCasbinRuleWithZero set the CasbinRule
func (crc *CasbinRuleCreate) SetCasbinRuleWithZero(input *CasbinRule, fields ...string) *CasbinRuleCreate {
	m := crc.mutation
	if len(fields) == 0 {
		fields = casbinrule.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return crc
}

// CasbinRuleCreateBulk is the builder for creating many CasbinRule entities in bulk.
type CasbinRuleCreateBulk struct {
	config
	err      error
	builders []*CasbinRuleCreate
}

// Save creates the CasbinRule entities in the database.
func (crcb *CasbinRuleCreateBulk) Save(ctx context.Context) ([]*CasbinRule, error) {
	if crcb.err != nil {
		return nil, crcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(crcb.builders))
	nodes := make([]*CasbinRule, len(crcb.builders))
	mutators := make([]Mutator, len(crcb.builders))
	for i := range crcb.builders {
		func(i int, root context.Context) {
			builder := crcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CasbinRuleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, crcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, crcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crcb *CasbinRuleCreateBulk) SaveX(ctx context.Context) []*CasbinRule {
	v, err := crcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crcb *CasbinRuleCreateBulk) Exec(ctx context.Context) error {
	_, err := crcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcb *CasbinRuleCreateBulk) ExecX(ctx context.Context) {
	if err := crcb.Exec(ctx); err != nil {
		panic(err)
	}
}
