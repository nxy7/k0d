package compose

import (
	"fmt"
	"k0d/utils"

	"github.com/spf13/cobra"
)

func MakeComposeFile(cmd *cobra.Command) string {
	k0sConfigPath, err := utils.SaveToFile(MakeK0sConfig(), "k0sconfig.yaml")
	if err != nil {
		panic(err)
	}

	containerdConfigPath, err := utils.SaveToFile(utils.MakeRegistryConfig("noxy.ddns.net:5000"), "containerd.toml")
	if err != nil {
		panic(err)
	}

	config := fmt.Sprintf(`services:
  k0s:
    container_name: k0s
    image: docker.io/k0sproject/k0s:v1.29.1-k0s.1
    command: k0s controller --config=/etc/k0s/config.yaml --single
    hostname: k0s
    privileged: true
    cgroup: host
    volumes:
      - %s:/etc/k0s/config.yaml
      - %s:/etc/k0s/containerd.toml
      - /var/lib/k0s
      - /sys/fs/cgroup:/run/cilium/cgroupv2:shared
      - /run/udev:/run/udev:slave
      - /sys/fs/bpf:/sys/fs/bpf:shared
    ports:
      - "6443:6443"
      - "80:80"
      - "443:443"
    network_mode: "bridge"`, k0sConfigPath, containerdConfigPath)

	return config
}
