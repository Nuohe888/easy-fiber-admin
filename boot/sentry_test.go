package boot

import (
	"easy-fiber-admin/pkg/config"
	"github.com/getsentry/sentry-go"
	"testing"
	"time"
)

func TestInitSentry(t *testing.T) {
	// Test case 1: DSN is provided
	// Note: Using a syntactically valid but non-functional DSN for testing Init.
	// Sentry SDK's Init will attempt to parse DSN and set up a client.
	// It won't actually send data unless it's a real DSN and network is available.
	cfgWithDsn := &config.Config{
		Sentry: config.SentryConfig{Dsn: "https://examplePublicKey@o0.ingest.sentry.io/0"},
	}
	InitSentry(cfgWithDsn)
	if sentry.CurrentHub().Client() == nil {
		t.Error("Sentry client should be initialized when DSN is provided")
	}
	// It's good practice to flush and unbind client if tests might interfere with each other
	// or if other parts of the test suite expect a clean Sentry state.
	sentry.Flush(2 * time.Second)       // Ensure events are flushed (though none are sent here)
	sentry.CurrentHub().BindClient(nil) // Unbind client for subsequent tests or clean state

	// Test case 2: DSN is not provided
	cfgWithoutDsn := &config.Config{
		Sentry: config.SentryConfig{Dsn: ""}, // Empty DSN
	}
	InitSentry(cfgWithoutDsn)
	if sentry.CurrentHub().Client() != nil {
		// If DSN is empty, InitSentry should log and not initialize the client.
		// Thus, the client on the current hub should remain nil (or whatever it was before if not reset).
		t.Error("Sentry client should not be initialized when DSN is empty")
	}
	// Clean up, though in this case, no client was bound by InitSentry.
	sentry.CurrentHub().BindClient(nil)
}
