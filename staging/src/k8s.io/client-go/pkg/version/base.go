package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.12-vip"
	gitCommit    = "26c89157669bdc7e3657302bc1abde758507095b"
	gitTreeState = "clean"
	buildDate = "2026-09-23T17:08:30Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)
