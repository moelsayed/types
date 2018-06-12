package client

const (
	KubeAPIServiceType                       = "kubeAPIService"
	KubeAPIServiceFieldExtraArgs             = "extraArgs"
	KubeAPIServiceFieldExtraBinds            = "extraBinds"
	KubeAPIServiceFieldExtraEnv              = "extraEnv"
	KubeAPIServiceFieldImage                 = "image"
	KubeAPIServiceFieldPodSecurityPolicy     = "podSecurityPolicy"
	KubeAPIServiceFieldServiceClusterIPRange = "serviceClusterIpRange"
	KubeAPIServiceFieldServiceNodePortRange  = "serviceNodePortRange"
)

type KubeAPIService struct {
	ExtraArgs             map[string]string `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
	ExtraBinds            []string          `json:"extraBinds,omitempty" yaml:"extraBinds,omitempty"`
	ExtraEnv              map[string]string `json:"extraEnv,omitempty" yaml:"extraEnv,omitempty"`
	Image                 string            `json:"image,omitempty" yaml:"image,omitempty"`
	PodSecurityPolicy     bool              `json:"podSecurityPolicy,omitempty" yaml:"podSecurityPolicy,omitempty"`
	ServiceClusterIPRange string            `json:"serviceClusterIpRange,omitempty" yaml:"serviceClusterIpRange,omitempty"`
	ServiceNodePortRange  string            `json:"serviceNodePortRange,omitempty" yaml:"serviceNodePortRange,omitempty"`
}
