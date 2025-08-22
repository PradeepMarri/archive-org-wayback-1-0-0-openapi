package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Url string `json:"url,omitempty"` // The URL requested
	Status int `json:"status,omitempty"` // The HTTP status of the URL requested
	Timestamp string `json:"timestamp,omitempty"` // The timestamp of the snapshot in [RFC 3339](http://xml2rfc.ietf.org/public/rfc/html/rfc3339.html) format
}

// ArchivedResult represents the ArchivedResult schema from the OpenAPI specification
type ArchivedResult struct {
	Tag string `json:"tag,omitempty"` // The user-supplied tag for use in collation
	Timestamp string `json:"timestamp"` // The _intepreted_ timestamp requested, in [RFC 3339](http://xml2rfc.ietf.org/public/rfc/html/rfc3339.html) format
	Url string `json:"url"` // The URL requested
	Snapshot Snapshot `json:"snapshot"`
}

// AvailabilityRequest represents the AvailabilityRequest schema from the OpenAPI specification
type AvailabilityRequest struct {
	Closest string `json:"closest,omitempty"` // The direction to find the closest snapshot to the requested timestamp
	Tag string `json:"tag,omitempty"` // A user-supplied tag, used for collation
	Timestamp string `json:"timestamp,omitempty"` // Timestamp requested in ISO 8601 format. The following formats are acceptable: - YYYY - YYYY-MM - YYYY-MM-DD - YYYY-MM-DDTHH:mm:SSz - YYYY-MM-DD:HH:mm+00:00
	Url string `json:"url"` // The URL requested
}

// AvailabilityResults represents the AvailabilityResults schema from the OpenAPI specification
type AvailabilityResults struct {
	Results []ArchivedResult `json:"results"` // A list of results
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}
