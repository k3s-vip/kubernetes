package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.13-vip"
	gitCommit    = "5bafe0f04b114c29e7bda3bd2d08919a8457b8d7"
	gitTreeState = "clean"
	buildDate = "2026-10-27T12:03:02Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)
