package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.32"
)

var (
	gitMajor = "1"
	gitMinor = "32"
	gitVersion   = "v1.32.13-vip"
	gitCommit    = "1214ccfca074d7677ec959e327bc03148ee436f4"
	gitTreeState = "clean"
	buildDate = "2026-02-26T20:20:22Z"
)
