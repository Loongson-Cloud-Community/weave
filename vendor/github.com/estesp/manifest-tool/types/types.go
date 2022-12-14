package types

import (
	"github.com/docker/distribution/manifest/manifestlist"
	containerTypes "github.com/docker/engine-api/types/container"
)

// ImageInspect holds information about an image in a registry
type ImageInspect struct {
	Size            int64
	MediaType       string
	Tag             string
	Digest          string
	RepoTags        []string
	Comment         string
	Created         string
	ContainerConfig *containerTypes.Config
	DockerVersion   string
	Author          string
	Config          *containerTypes.Config
	Architecture    string
	Os              string
	Layers          []string
	Platform        manifestlist.PlatformSpec
	CanonicalJSON   []byte
}

// YAMLInput represents the YAML format input to the pushml
// command.
type YAMLInput struct {
	Image     string
	Manifests []ManifestEntry
}

// ManifestEntry represents an entry in the list of manifests to
// be combined into a manifest list, provided via the YAML input
type ManifestEntry struct {
	Image    string
	Platform manifestlist.PlatformSpec
}

// AuthInfo holds information about how manifest-tool should connect and authenticate to the docker registry
type AuthInfo struct {
	Username  string
	Password  string
	DockerCfg string
}
