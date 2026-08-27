package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.10-vip"
	gitCommit    = "ae6fece6fea92ecfcccd13af7d5f2ea8cb9b33fe"
	gitTreeState = "clean"
	buildDate = "2026-10-27T12:02:52Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
