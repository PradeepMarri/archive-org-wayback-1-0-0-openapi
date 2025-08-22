package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/wayback-api/mcp-server/config"
	"github.com/wayback-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_wayback_v1_availableHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("url=%v", val))
		}
		if val, ok := args["timestamp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timestamp=%v", val))
		}
		if val, ok := args["callback"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("callback=%v", val))
		}
		if val, ok := args["timeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeout=%v", val))
		}
		if val, ok := args["closest"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("closest=%v", val))
		}
		if val, ok := args["status_code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status_code=%v", val))
		}
		if val, ok := args["tag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/wayback/v1/available%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.AvailabilityResults
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_wayback_v1_availableTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_wayback_v1_available",
		mcp.WithDescription(""),
		mcp.WithString("url", mcp.Required(), mcp.Description("A single URL to query.")),
		mcp.WithString("timestamp", mcp.Description("Timestamp requested in ISO 8601 format. The following formats are acceptable:\n - YYYY\n - YYYY-MM\n - YYYY-MM-DD\n - YYYY-MM-DDTHH:mm:SSz\n - YYYY-MM-DD:HH:mm+00:00\n")),
		mcp.WithString("callback", mcp.Description("Specifies a JavaScript function func, for a JSON-P response. When provided, results are wrapped as `callback(data)`, and the returned MIME type is application/javascript. This causes the caller to automatically run the func with the JSON results as its argument.\n")),
		mcp.WithString("timeout", mcp.Description("Timeout is the maximum number of seconds to wait for the availability API to get its underlying results from the CDX server. The default value is 5.0.\n")),
		mcp.WithString("closest", mcp.Description("The direction specifies whether to match archived timestamps that are before the provided one, after, or the default either (closest in either direction). Must be before, after, or either. May be overidden by individual requests.\n")),
		mcp.WithNumber("status_code", mcp.Description("HTTP status codes to filter by. Only results with these codes will be returned\n")),
		mcp.WithString("tag", mcp.Description("The optional tag can have any value, and is returned with the results; it can be used to help collate input and output values.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_wayback_v1_availableHandler(cfg),
	}
}
