package compose

// Returns config
func MakeK0sConfig() string {
	return `apiVersion: k0s.k0sproject.io/v1beta1
kind: ClusterConfig
metadata:
  name: k0s
spec:
  api:
    address: 172.17.0.2
    port: 6443
    sans:
      - 172.17.0.2
      - 172.17.0.3
      - 172.17.0.4
      - 172.17.0.5
      - 10.96.0.1
  network:
    provider: custom
    kubeProxy:
      disabled: true
  storage:
    type: kine
`
}
