package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.11-vip"
	gitCommit    = "42b327bab2956b1f7dc113805ac73065e6a7866e"
	gitTreeState = "clean"
	buildDate = "2026-11-25T10:41:12Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
