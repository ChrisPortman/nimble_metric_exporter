package version

// Version is set in CI using an LDFLAG.  E.g.:
//
//	go build -v -ldflags "-X 'github.com/ChrisPortman/nimble_metric_exporter/internal/version.Version=<my version>"
var (
	Version   = "develop"
	BuildDate = ""
	GoVersion = ""
)
