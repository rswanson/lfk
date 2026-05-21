// Package fakeapi serves a minimal subset of the Kubernetes API for
// deterministic energy benchmarks. It is NOT a general-purpose K8s mock —
// it implements only the endpoints lfk's informers and resource list calls
// hit during the harness scenarios.
package fakeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

// Fixture holds the in-memory K8s objects served by the fake.
type Fixture struct {
	Pods        []map[string]any
	Services    []map[string]any
	Deployments []map[string]any
	Namespaces  []map[string]any
}

// Server is a running httptest server backed by a Fixture.
type Server struct {
	*httptest.Server
	fx Fixture
}

// New returns a Server not yet started.
func New(fx Fixture) *Server { return &Server{fx: fx} }

// Start launches the underlying httptest server.
func (s *Server) Start() *Server {
	mux := http.NewServeMux()
	s.installRoutes(mux)
	s.Server = httptest.NewServer(mux)
	return s
}

func (s *Server) installRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", okHandler)
	mux.HandleFunc("/readyz", okHandler)
	mux.HandleFunc("/version", versionHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/apis", apisHandler)
	mux.HandleFunc("/api/v1", coreV1DiscoveryHandler)
	mux.HandleFunc("/apis/apps/v1", appsV1DiscoveryHandler)
	mux.HandleFunc("/api/v1/namespaces/default/pods", s.listOrWatch("pods", s.fx.Pods))
	mux.HandleFunc("/api/v1/pods", s.listOrWatch("pods", s.fx.Pods))
	mux.HandleFunc("/api/v1/namespaces/default/services", s.listOrWatch("services", s.fx.Services))
	mux.HandleFunc("/api/v1/namespaces", s.listOrWatch("namespaces", s.fx.Namespaces))
	mux.HandleFunc("/apis/apps/v1/namespaces/default/deployments", s.listOrWatch("deployments", s.fx.Deployments))
}

func okHandler(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("ok")) }

func versionHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{"gitVersion": "v1.30.0-fake", "major": "1", "minor": "30"})
}

func apiHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]any{
		"kind":     "APIVersions",
		"versions": []string{"v1"},
	})
}

func apisHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]any{
		"kind": "APIGroupList",
		"groups": []map[string]any{
			{
				"name":             "apps",
				"versions":         []map[string]string{{"groupVersion": "apps/v1", "version": "v1"}},
				"preferredVersion": map[string]string{"groupVersion": "apps/v1", "version": "v1"},
			},
		},
	})
}

func coreV1DiscoveryHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]any{
		"kind":         "APIResourceList",
		"groupVersion": "v1",
		"resources": []map[string]any{
			{"name": "pods", "namespaced": true, "kind": "Pod", "verbs": []string{"list", "watch"}},
			{"name": "services", "namespaced": true, "kind": "Service", "verbs": []string{"list", "watch"}},
			{"name": "namespaces", "namespaced": false, "kind": "Namespace", "verbs": []string{"list", "watch"}},
		},
	})
}

func appsV1DiscoveryHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]any{
		"kind":         "APIResourceList",
		"groupVersion": "apps/v1",
		"resources": []map[string]any{
			{"name": "deployments", "namespaced": true, "kind": "Deployment", "verbs": []string{"list", "watch"}},
		},
	})
}

// listOrWatch returns a handler that either lists items or streams a
// long-lived watch (bookmark events every 30s) depending on the ?watch query.
func (s *Server) listOrWatch(kind string, items []map[string]any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.RawQuery, "watch=true") {
			s.streamWatch(w, r, kind)
			return
		}
		writeJSON(w, map[string]any{
			"kind":     listKind(kind),
			"metadata": map[string]any{"resourceVersion": "1"},
			"items":    items,
		})
	}
}

// listKind capitalises the first rune of kind and appends "List".
func listKind(kind string) string {
	if kind == "" {
		return "List"
	}
	return strings.ToUpper(kind[:1]) + kind[1:] + "List"
}

func (s *Server) streamWatch(w http.ResponseWriter, r *http.Request, kind string) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "no flusher", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	enc := json.NewEncoder(w)
	for {
		_ = enc.Encode(map[string]any{
			"type":   "BOOKMARK",
			"object": map[string]any{"kind": kind, "metadata": map[string]any{"resourceVersion": "1"}},
		})
		flusher.Flush()
		select {
		case <-ticker.C:
		case <-r.Context().Done():
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
