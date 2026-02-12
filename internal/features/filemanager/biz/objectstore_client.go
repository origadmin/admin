/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"

	pbtypes "origadmin/application/admin/api/v1/services/types"
)

// ObjectStoreClient defines the interface for interacting with the ObjectStore service.
type ObjectStoreClient interface {
	CreateObject(ctx context.Context, name string, data []byte) (*pbtypes.Object, error)
	GetObject(ctx context.Context, id string) (*pbtypes.Object, error)
	DeleteObject(ctx context.Context, id string) error
}
