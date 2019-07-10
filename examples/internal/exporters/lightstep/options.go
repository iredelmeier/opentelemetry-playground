package lightstep

import (
	"os"
	"strconv"

	"github.com/lightstep/lightstep-tracer-go"
	"github.com/opentracing/opentracing-go"
)

// By default, the exporter will attempt to connect to a local satellite
const (
	DefaultSatelliteHost = "localhost"
	DefaultSatellitePort = 8360
)

// Option provides configuration for the Exporter
type Option func(*config)

// WithAccessToken sets an access token for communicating with LightStep
func WithAccessToken(accessToken string) Option {
	return func(c *config) {
		c.tracerOptions.AccessToken = accessToken
	}
}

// WithSatelliteHost sets the satellite host to which spans will be sent
func WithSatelliteHost(satelliteHost string) Option {
	return func(c *config) {
		c.tracerOptions.Collector.Host = satelliteHost
	}
}

// WithSatellitePort sets the satellite port to which spans will be sent
func WithSatellitePort(satellitePort int) Option {
	return func(c *config) {
		c.tracerOptions.Collector.Port = satellitePort
	}
}

// WithInsecure prevents the Exporter from communicating over TLS with the satellite,
// i.e., the connection will run over HTTP instead of HTTPS
func WithInsecure(insecure bool) Option {
	return func(c *config) {
		c.tracerOptions.Collector.Plaintext = insecure
	}
}

// WithMetaEventReportingEnabled configures the tracer to send meta events,
// e.g., events for span creation
func WithMetaEventReportingEnabled(metaEventReportingEnabled bool) Option {
	return func(c *config) {
		c.tracerOptions.MetaEventReportingEnabled = metaEventReportingEnabled
	}
}

// WithComponentName overrides the component (service) name that will be used in LightStep
func WithComponentName(componentName string) Option {
	return func(c *config) {
		if componentName != "" {
			c.tracerOptions.Tags[lightstep.ComponentNameKey] = componentName
		}
	}
}

type config struct {
	tracerOptions lightstep.Options
}

func newConfig(opts ...Option) config {
	c := config{
		tracerOptions: lightstep.Options{
			UseHttp: true,
			Tags:    make(opentracing.Tags),
		},
	}
	defaultOpts := []Option{
		WithSatelliteHost(DefaultSatelliteHost),
		WithSatellitePort(DefaultSatellitePort),
	}

	if accessToken := os.Getenv("LIGHTSTEP_ACCESS_TOKEN"); accessToken != "" {
		defaultOpts = append(defaultOpts, WithAccessToken(accessToken))
	}

	if satelliteHost := os.Getenv("LIGHTSTEP_SATELLITE_HOST"); satelliteHost != "" {
		defaultOpts = append(defaultOpts, WithSatelliteHost(satelliteHost))
	}

	if satellitePort := os.Getenv("LIGHTSTEP_SATELLITE_PORT"); satellitePort != "" {
		if i, err := strconv.Atoi(satellitePort); err == nil {
			defaultOpts = append(defaultOpts, WithSatellitePort(i))
		}
	}

	if insecure := os.Getenv("LIGHTSTEP_INSECURE"); insecure != "" {
		if b, err := strconv.ParseBool(insecure); err == nil {
			defaultOpts = append(defaultOpts, WithInsecure(b))
		}
	}

	if metaEventReportingEnabled := os.Getenv("LIGHTSTEP_META_EVENT_REPORTING_ENABLED"); metaEventReportingEnabled != "" {
		if b, err := strconv.ParseBool(metaEventReportingEnabled); err == nil {
			defaultOpts = append(defaultOpts, WithMetaEventReportingEnabled(b))
		}
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(&c)
	}

	return c
}
