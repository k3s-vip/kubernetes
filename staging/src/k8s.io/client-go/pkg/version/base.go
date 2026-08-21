package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.11-vip"
	gitCommit    = "3a634765b787dd069f7f714fa77d767cb7d43795"
	gitTreeState = "clean"
	buildDate = "2026-08-20T15:15:36Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)
