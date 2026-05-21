package fakeapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixtureSizes(t *testing.T) {
	tests := []struct {
		name string
		fn   func() Fixture
		pods int
	}{
		{"small", SmallFixture, 10},
		{"medium", MediumFixture, 200},
		{"large", LargeFixture, 2000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fx := tt.fn()
			assert.Len(t, fx.Pods, tt.pods)
			assert.NotEmpty(t, fx.Namespaces)
			assert.NotEmpty(t, fx.Services)
			assert.NotEmpty(t, fx.Deployments)
		})
	}
}

func TestFixtureDeterminism(t *testing.T) {
	a := MediumFixture()
	b := MediumFixture()
	assert.Equal(t, a.Pods[0]["metadata"], b.Pods[0]["metadata"],
		"fixtures must be deterministic for reproducible runs")
}
