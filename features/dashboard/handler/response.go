package handler

import "immersive_project/klp3/features/dashboard"

type DashboardResponse struct {
	CountActive    int64 `json:"count_active"`
	CountPlacement int64 `json:"count_placement"`
	CountLog       int64 `json:"count_log"`
}

func DataToResponse(input dashboard.DashboardEntity) DashboardResponse {
	var resp = DashboardResponse{
		CountActive:    input.CountActive,
		CountPlacement: input.CountPlacement,
		CountLog:       input.CountLog,
	}
	return resp
}
