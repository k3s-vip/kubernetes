package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.8-vip"
	gitCommit    = "1c2e10a409eb1b03f2f28f401ce935312e20d9fb"
	gitTreeState = "clean"
	buildDate = "2026-08-20T15:16:14Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
