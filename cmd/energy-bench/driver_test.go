package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteKubeconfig_ThreeContextsPointingAtURL(t *testing.T) {
	dir := t.TempDir()
	kc := filepath.Join(dir, "kubeconfig")
	require.NoError(t, writeKubeconfig(kc, "http://127.0.0.1:9999"))

	data, err := os.ReadFile(kc)
	require.NoError(t, err)
	s := string(data)
	assert.Contains(t, s, "ctx-a")
	assert.Contains(t, s, "ctx-b")
	assert.Contains(t, s, "ctx-c")
	assert.Contains(t, s, "http://127.0.0.1:9999")
}
