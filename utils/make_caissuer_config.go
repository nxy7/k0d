package utils

func MakeCaIssuerConfig() string {
	return `apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: ca-issuer
  namespace: cert-manager
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: ca-key
    solvers:
    - http01:
        gatewayHTTPRoute:
          parentRefs:
            - name: gateway
              namespace: default`
}
