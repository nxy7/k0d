CLI tools for creating k0s in docker (similarly to what k3d is doing for k3s).
It needs the following tools to be already installed on the system: docker, docker compose, helm

# Current status
This tool currently starts up single-node cluster with cilium, gateway API, OpenEBS and CertManager enabled.
Configuration mechanism for setting up multi-node cluster and available utilities is not yet implemented.
