package fakeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_HealthzAndReadyz(t *testing.T) {
	s := New(Fixture{}).Start()
	t.Cleanup(s.Close)

	for _, path := range []string{"/healthz", "/readyz"} {
		resp, err := http.Get(s.URL + path)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected 200 on %s", path)
		_ = resp.Body.Close()
	}
}

func TestServer_ListsPodsFromFixture(t *testing.T) {
	fx := Fixture{
		Pods: []map[string]any{
			{"metadata": map[string]any{"name": "pod-1", "namespace": "default"}},
			{"metadata": map[string]any{"name": "pod-2", "namespace": "default"}},
		},
	}
	s := New(fx).Start()
	t.Cleanup(s.Close)

	resp, err := http.Get(s.URL + "/api/v1/namespaces/default/pods")
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var list struct {
		Items []json.RawMessage `json:"items"`
	}
	require.NoError(t, json.Unmarshal(body, &list))
	assert.Len(t, list.Items, 2)
}

func TestServer_CoreDiscoveryHasNamespacesAndPods(t *testing.T) {
	s := New(Fixture{}).Start()
	t.Cleanup(s.Close)

	resp, err := http.Get(s.URL + "/api/v1")
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), `"name":"namespaces"`)
	assert.Contains(t, string(body), `"name":"pods"`)
}
