package utils

import "fmt"

func AddInsecureRegistry() {
	config := MakeRegistryConfig("noxy.ddns.net:5000")

	fmt.Println(config)
}

func MakeRegistryConfig(registryUrl string) string {
	// "/run/k0s/containerd-cri.toml",
	return fmt.Sprintf(`version = 2
imports = [
]
[plugins]
[plugins."io.containerd.grpc.v1.cri"]
[plugins."io.containerd.grpc.v1.cri".registry]
[plugins."io.containerd.grpc.v1.cri".registry.mirrors]
[plugins."io.containerd.grpc.v1.cri".registry.mirrors."%s"]
  endpoint = ["http://%s"]
[plugins."io.containerd.grpc.v1.cri".registry.configs]
[plugins."io.containerd.grpc.v1.cri".registry.configs."%s"]
[plugins."io.containerd.grpc.v1.cri".registry.configs."%s".tls]
  insecure_skip_verify = true`, registryUrl, registryUrl, registryUrl, registryUrl)
}
