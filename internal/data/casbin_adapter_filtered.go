/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"context"
	"fmt"
	"reflect"

	"github.com/casbin/casbin/v3/model"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
)

// Filter defines the filtering rules for a FilteredAdapter's policy. Empty values
// are ignored, but all others must match the filter.
type Filter struct {
	P  []string
	G  []string
	G1 []string
	G2 []string
	G3 []string
	G4 []string
	G5 []string
}

// FilteredAdapter is the filtered adapter for Casbin. It can load policy
// from database or save policy to database and supports loading of filtered policies.
type FilteredAdapter struct {
	*Adapter
	filtered bool
}

// NewFilteredAdapter creates a new filtered casbin adapter.
func NewFilteredAdapter(ctx context.Context, db *ent.Database, logger log.Logger) (*FilteredAdapter, error) {
	adapter, err := NewAdapter(ctx, db, logger)
	if err != nil {
		return nil, err
	}
	return &FilteredAdapter{
		Adapter:  adapter,
		filtered: true,
	}, nil
}

// NewFilteredAdapterFromApp creates a new filtered casbin adapter.
func NewFilteredAdapterFromApp(app *runtime.App, db *ent.Database) (*FilteredAdapter, error) {
	adapter, err := NewAdapterFromApp(app, db)
	if err != nil {
		return nil, err
	}
	return &FilteredAdapter{
		Adapter:  adapter,
		filtered: true,
	}, nil
}

// LoadPolicy loads all policy rules from the storage.
func (a *FilteredAdapter) LoadPolicy(model model.Model) error {
	a.filtered = false
	return a.Adapter.LoadPolicy(model)
}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *FilteredAdapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	if filter == nil {
		return a.LoadPolicy(model)
	}

	filterValue, ok := filter.(*Filter)
	if !ok {
		return fmt.Errorf("invalid filter type: %v", reflect.TypeOf(filter))
	}

	session := a.db.CasbinRule(a.ctx).Query()

	if len(filterValue.P) != 0 {
		session.Where(casbinrule.PtypeIn("p"))
		if len(filterValue.P) > 1 {
			session.Where(casbinrule.V0In(filterValue.P...))
		} else {
			session.Where(casbinrule.V0EQ(filterValue.P[0]))
		}
	}
	if len(filterValue.G) != 0 {
		session.Where(casbinrule.PtypeIn("g"))
		if len(filterValue.G) > 1 {
			session.Where(casbinrule.V0In(filterValue.G...))
		} else {
			session.Where(casbinrule.V0EQ(filterValue.G[0]))
		}
	}
	if len(filterValue.G1) != 0 {
		session.Where(casbinrule.PtypeIn("g1"))
		if len(filterValue.G1) > 1 {
			session.Where(casbinrule.V1In(filterValue.G1...))
		} else {
			session.Where(casbinrule.V1EQ(filterValue.G1[0]))
		}
	}
	if len(filterValue.G2) != 0 {
		session.Where(casbinrule.PtypeIn("g2"))
		if len(filterValue.G2) > 1 {
			session.Where(casbinrule.V2In(filterValue.G2...))
		} else {
			session.Where(casbinrule.V2EQ(filterValue.G2[0]))
		}
	}
	if len(filterValue.G3) != 0 {
		session.Where(casbinrule.PtypeIn("g3"))
		if len(filterValue.G3) > 1 {
			session.Where(casbinrule.V3In(filterValue.G3...))
		} else {
			session.Where(casbinrule.V3EQ(filterValue.G3[0]))
		}
	}
	if len(filterValue.G4) != 0 {
		session.Where(casbinrule.PtypeIn("g4"))
		if len(filterValue.G4) > 1 {
			session.Where(casbinrule.V4In(filterValue.G4...))
		} else {
			session.Where(casbinrule.V4EQ(filterValue.G4[0]))
		}
	}
	if len(filterValue.G5) != 0 {
		session.Where(casbinrule.PtypeIn("g5"))
		if len(filterValue.G5) > 1 {
			session.Where(casbinrule.V5In(filterValue.G5...))
		} else {
			session.Where(casbinrule.V5EQ(filterValue.G5[0]))
		}
	}

	lines, err := session.All(a.ctx)
	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	a.filtered = true

	return nil
}

// IsFiltered returns true if the loaded policy has been filtered.
func (a *FilteredAdapter) IsFiltered() bool {
	return a.filtered
}

// SavePolicy saves all policy rules to the storage.
func (a *FilteredAdapter) SavePolicy(model model.Model) error {
	if a.filtered {
		return fmt.Errorf("cannot save a filtered policy")
	}
	return a.Adapter.SavePolicy(model)
}
