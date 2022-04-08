/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// DashboardSnapshot model
type DashboardSnapshot struct {
	Created            time.Time `json:"Created,omitempty"`
	Dashboard          *Json     `json:"Dashboard,omitempty"`
	DashboardEncrypted []int32   `json:"DashboardEncrypted,omitempty"`
	DeleteKey          string    `json:"DeleteKey,omitempty"`
	Expires            time.Time `json:"Expires,omitempty"`
	External           bool      `json:"External,omitempty"`
	ExternalDeleteUrl  string    `json:"ExternalDeleteUrl,omitempty"`
	ExternalUrl        string    `json:"ExternalUrl,omitempty"`
	Id                 int64     `json:"Id,omitempty"`
	Key                string    `json:"Key,omitempty"`
	Name               string    `json:"Name,omitempty"`
	OrgId              int64     `json:"OrgId,omitempty"`
	Updated            time.Time `json:"Updated,omitempty"`
	UserId             int64     `json:"UserId,omitempty"`
}