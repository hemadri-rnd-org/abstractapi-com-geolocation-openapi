package main

import (
	"github.com/ip-geolocation-api/mcp-server/config"
	"github.com/ip-geolocation-api/mcp-server/models"
	tools_v1 "github.com/ip-geolocation-api/mcp-server/tools/v1"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v1.CreateGet_v1Tool(cfg),
	}
}
