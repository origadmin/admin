package dto

//go:generate go run github.com/origadmin/abgen/cmd/abgen -debug go run ./cmd/abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/features/system/data/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"
