/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

import "time"

// FileMetadata is the data transfer object for file metadata.
// It represents the data structure stored in the data access layer.
type FileMetadata struct {
	ID        string
	Name      string
	ObjectID  string
	OwnerID   string
	CreatedAt time.Time
}
