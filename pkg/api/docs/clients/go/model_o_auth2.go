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

type OAuth2 struct {
	TLSConfig        *TlsConfig        `json:"TLSConfig,omitempty"`
	ClientId         string            `json:"client_id,omitempty"`
	ClientSecret     *Secret           `json:"client_secret,omitempty"`
	ClientSecretFile string            `json:"client_secret_file,omitempty"`
	EndpointParams   map[string]string `json:"endpoint_params,omitempty"`
	Scopes           []string          `json:"scopes,omitempty"`
	TokenUrl         string            `json:"token_url,omitempty"`
}