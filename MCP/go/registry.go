package main

import (
	"github.com/wayback-api/mcp-server/config"
	"github.com/wayback-api/mcp-server/models"
	tools_wayback "github.com/wayback-api/mcp-server/tools/wayback"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_wayback.CreatePost_wayback_v1_availableTool(cfg),
		tools_wayback.CreateGet_wayback_v1_availableTool(cfg),
	}
}
