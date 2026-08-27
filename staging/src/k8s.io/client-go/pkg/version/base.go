package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.9-vip"
	gitCommit    = "329000efd0b26bd78155e51e4df449130f0fca00"
	gitTreeState = "clean"
	buildDate = "2026-09-23T17:07:48Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
