/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package generic implements the functions, types, and interfaces for the module.
package generic

func ConvertSliceToPB[GoModel any, PBModel any](gos []GoModel, converter func(GoModel) PBModel) []PBModel {
	pbs := make([]PBModel, 0, len(gos))
	for _, m := range gos {
		pbs = append(pbs, converter(m))
	}
	return pbs
}

func ConvertSliceToGo[GoModel any, PBModel any](pbs []PBModel, converter func(PBModel) GoModel) []GoModel {
	gos := make([]GoModel, 0, len(pbs))
	for _, m := range pbs {
		gos = append(gos, converter(m))
	}
	return gos
}
