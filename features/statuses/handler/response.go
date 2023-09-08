package handler

import (
	"immersive_project/klp3/features/statuses"
)

type StatusResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func StatusEntityToResponse(status statuses.StatusEntity) StatusResponse {
	return StatusResponse{
		Id:   status.ID,
		Name: status.Name,
	}
}
