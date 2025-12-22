package dto

import (
	"origadmin/application/admin/internal/data/entity/ent/user"
)

//go:abgen:package:path=origadmin/application/admin/internal/data/entity/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="origadmin/application/admin/internal/data/entity/ent,origadmin/application/admin/api/v1/services/types"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"
//go:abgen:convert:direction="both"
//go:abgen:convert="source=ent.Resource,target=types.Menu"

type (
	Gender = user.Gender
)
