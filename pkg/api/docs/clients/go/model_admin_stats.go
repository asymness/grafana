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

type AdminStats struct {
	ActiveAdmins        int64 `json:"activeAdmins,omitempty"`
	ActiveEditors       int64 `json:"activeEditors,omitempty"`
	ActiveSessions      int64 `json:"activeSessions,omitempty"`
	ActiveUsers         int64 `json:"activeUsers,omitempty"`
	ActiveViewers       int64 `json:"activeViewers,omitempty"`
	Admins              int64 `json:"admins,omitempty"`
	Alerts              int64 `json:"alerts,omitempty"`
	DailyActiveAdmins   int64 `json:"dailyActiveAdmins,omitempty"`
	DailyActiveEditors  int64 `json:"dailyActiveEditors,omitempty"`
	DailyActiveSessions int64 `json:"dailyActiveSessions,omitempty"`
	DailyActiveUsers    int64 `json:"dailyActiveUsers,omitempty"`
	DailyActiveViewers  int64 `json:"dailyActiveViewers,omitempty"`
	Dashboards          int64 `json:"dashboards,omitempty"`
	Datasources         int64 `json:"datasources,omitempty"`
	Editors             int64 `json:"editors,omitempty"`
	MonthlyActiveUsers  int64 `json:"monthlyActiveUsers,omitempty"`
	Orgs                int64 `json:"orgs,omitempty"`
	Playlists           int64 `json:"playlists,omitempty"`
	Snapshots           int64 `json:"snapshots,omitempty"`
	Stars               int64 `json:"stars,omitempty"`
	Tags                int64 `json:"tags,omitempty"`
	Users               int64 `json:"users,omitempty"`
	Viewers             int64 `json:"viewers,omitempty"`
}