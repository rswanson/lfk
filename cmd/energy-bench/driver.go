package main

import (
	"fmt"
	"os"
)

// writeKubeconfig writes a minimal kubeconfig with three contexts all
// pointing at server. Used by the harness so lfk exercises its
// multi-context code path against the fake apiserver.
func writeKubeconfig(path, server string) error {
	contents := fmt.Sprintf(`apiVersion: v1
kind: Config
clusters:
- name: fake
  cluster:
    server: %s
    insecure-skip-tls-verify: true
contexts:
- name: ctx-a
  context: {cluster: fake, user: fake-user, namespace: default}
- name: ctx-b
  context: {cluster: fake, user: fake-user, namespace: default}
- name: ctx-c
  context: {cluster: fake, user: fake-user, namespace: default}
current-context: ctx-a
users:
- name: fake-user
  user:
    token: fake-token
`, server)
	return os.WriteFile(path, []byte(contents), 0o600)
}
