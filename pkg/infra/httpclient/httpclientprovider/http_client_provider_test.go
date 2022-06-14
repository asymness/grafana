package httpclientprovider

import (
	"testing"

	"github.com/grafana/grafana/pkg/services/validations"

	sdkhttpclient "github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana/pkg/infra/tracing"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/stretchr/testify/require"
)

func TestHTTPClientProvider(t *testing.T) {
	t.Run("When creating new provider and SigV4 is disabled should apply expected middleware", func(t *testing.T) {
		origNewProviderFunc := newProviderFunc
		providerOpts := []sdkhttpclient.ProviderOptions{}
		newProviderFunc = func(opts ...sdkhttpclient.ProviderOptions) *sdkhttpclient.Provider {
			providerOpts = opts
			return nil
		}
		t.Cleanup(func() {
			newProviderFunc = origNewProviderFunc
		})
		tracer := tracing.InitializeTracerForTest()
		_ = New(&setting.Cfg{SigV4AuthEnabled: false}, &validations.OSSPluginRequestValidator{}, tracer)
		require.Len(t, providerOpts, 1)
		o := providerOpts[0]
		require.Len(t, o.Middlewares, 8)
		require.Equal(t, TracingMiddlewareName, o.Middlewares[0].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, DataSourceMetricsMiddlewareName, o.Middlewares[1].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, SetUserAgentMiddlewareName, o.Middlewares[2].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, sdkhttpclient.BasicAuthenticationMiddlewareName, o.Middlewares[3].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, sdkhttpclient.CustomHeadersMiddlewareName, o.Middlewares[4].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, sdkhttpclient.ContextualMiddlewareName, o.Middlewares[5].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, ResponseLimitMiddlewareName, o.Middlewares[6].(sdkhttpclient.MiddlewareName).MiddlewareName())
	})

	t.Run("When creating new provider and SigV4 is enabled should apply expected middleware", func(t *testing.T) {
		origNewProviderFunc := newProviderFunc
		providerOpts := []sdkhttpclient.ProviderOptions{}
		newProviderFunc = func(opts ...sdkhttpclient.ProviderOptions) *sdkhttpclient.Provider {
			providerOpts = opts
			return nil
		}
		t.Cleanup(func() {
			newProviderFunc = origNewProviderFunc
		})
		tracer := tracing.InitializeTracerForTest()
		_ = New(&setting.Cfg{SigV4AuthEnabled: true}, &validations.OSSPluginRequestValidator{}, tracer)
		require.Len(t, providerOpts, 1)
		o := providerOpts[0]
		require.Len(t, o.Middlewares, 9)
		require.Equal(t, TracingMiddlewareName, o.Middlewares[0].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, DataSourceMetricsMiddlewareName, o.Middlewares[1].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, SetUserAgentMiddlewareName, o.Middlewares[2].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, sdkhttpclient.BasicAuthenticationMiddlewareName, o.Middlewares[3].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, sdkhttpclient.CustomHeadersMiddlewareName, o.Middlewares[4].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, sdkhttpclient.ContextualMiddlewareName, o.Middlewares[5].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, ResponseLimitMiddlewareName, o.Middlewares[6].(sdkhttpclient.MiddlewareName).MiddlewareName())
		require.Equal(t, SigV4MiddlewareName, o.Middlewares[8].(sdkhttpclient.MiddlewareName).MiddlewareName())
	})
}
