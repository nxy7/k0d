package compose

import (
	"fmt"
	"k0d/utils"

	"github.com/spf13/cobra"
)

func MakeComposeFile(cmd *cobra.Command) string {
	k0sConfigPath := utils.SaveToFile(MakeK0sConfig(), "k0sconfig.yaml")
	containerdConfigPath := utils.SaveToFile(MakeK0sConfig(), "containerd.toml")

	config := fmt.Sprintf(`services:
  k0s:
    container_name: k0s
    image: docker.io/k0sproject/k0s:v1.27.4-k0s.0
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
