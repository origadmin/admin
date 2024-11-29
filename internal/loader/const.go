/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

const (
	Application = `OrigAdmin`
	WebSite     = `https://origadmin.com`
	Description = `A distributed backend management system with a focus on scalability, security, and flexibility.`
)

// UI generated by https://patorjk.com/software/taag/#p=display&f=Graffiti&t=origadmin
const UI = `
________        .__          _____       .___      .__        
\_____  \_______|__| ____   /  _  \    __| _/_____ |__| ____  
 /   |   \_  __ \  |/ ___\ /  /_\  \  / __ |/     \|  |/    \ 
/    |    \  | \/  / /_/  >    |    \/ /_/ |  Y Y  \  |   |  \
\_______  /__|  |__\___  /\____|__  /\____ |__|_|  /__|___|  /
        \/        /_____/         \/      \/     \/        \/ 
`

const (
	CacheNSForUser   = "user"
	CacheNSForRole   = "role"
	CacheNSForMenu   = "menu"
	CacheNSForRes    = "resource"
	CacheNSForAuth   = "auth"
	CacheNSForCasbin = "casbin"
)

const (
	CacheKeyForSyncToCasbin = "sync:casbin:update"
	CacheKeyForSyncedCasbin = "sync:casbin:success"
)

const (
	// InvalidTokenID is the error code for invalid token ID
	InvalidTokenID = "com.invalid.token"
	// InvalidCaptchaID is the error code for invalid captcha ID
	InvalidCaptchaID = "com.invalid.captcha"
	// InvalidUsernameOrPasswordID is the error code for invalid username or password
	InvalidUsernameOrPasswordID = "com.invalid.username.or.password"
)