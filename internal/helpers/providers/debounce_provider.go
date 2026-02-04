/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"time"

	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/debounce"
)

// NewDebounceExecutor creates a debounce executor based on the application configuration.
// It extracts the policy sync delay from the auth configuration.
func NewDebounceExecutor(c *conf.Config) debounce.Executor {
	var delay time.Duration
	policyDelay := c.GetBootstrap().GetAuth().GetPolicySyncDelay()
	if policyDelay != nil {
		delay = policyDelay.AsDuration()
	}
	return debounce.NewExecutor(delay)
}
