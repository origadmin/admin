/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"time"

	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/debounce"
)

// NewDebounceExecutor creates a debounce executor based on the application configuration.
// It extracts the policy sync delay from the auth configuration.
func NewDebounceExecutor(b *confpb.Bootstrap) debounce.Executor {
	var delay time.Duration
	policyDelay := b.GetAuth().GetPolicySyncDelay()
	if policyDelay != nil {
		delay = policyDelay.AsDuration()
	}
	return debounce.NewExecutor(delay)
}
