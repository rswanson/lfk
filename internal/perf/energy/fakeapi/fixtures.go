package fakeapi

import "fmt"

func SmallFixture() Fixture  { return generate(10) }
func MediumFixture() Fixture { return generate(200) }
func LargeFixture() Fixture  { return generate(2000) }

func generate(podCount int) Fixture {
	ns := []map[string]any{
		{"metadata": map[string]any{"name": "default"}},
		{"metadata": map[string]any{"name": "kube-system"}},
	}
	deployments := make([]map[string]any, 0, podCount/5)
	services := make([]map[string]any, 0, podCount/5)
	for i := 0; i < podCount/5; i++ {
		deployments = append(deployments, map[string]any{
			"metadata": map[string]any{"name": fmt.Sprintf("app-%03d", i), "namespace": "default"},
			"spec":     map[string]any{"replicas": 5},
		})
		services = append(services, map[string]any{
			"metadata": map[string]any{"name": fmt.Sprintf("svc-%03d", i), "namespace": "default"},
			"spec":     map[string]any{"clusterIP": fmt.Sprintf("10.0.%d.%d", i/256, i%256)},
		})
	}
	pods := make([]map[string]any, 0, podCount)
	for i := 0; i < podCount; i++ {
		pods = append(pods, map[string]any{
			"metadata": map[string]any{
				"name":      fmt.Sprintf("pod-%04d", i),
				"namespace": "default",
				"labels":    map[string]any{"app": fmt.Sprintf("app-%03d", i/5)},
			},
			"spec": map[string]any{
				"containers": []map[string]any{{"name": "main", "image": "nginx:1.25"}},
			},
			"status": map[string]any{"phase": "Running"},
		})
	}
	return Fixture{Pods: pods, Services: services, Deployments: deployments, Namespaces: ns}
}
