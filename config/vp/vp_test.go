package vp

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mgr := New()
	assert.NotNil(t, mgr)
	assert.NotNil(t, mgr.vp)
	assert.NotNil(t, mgr.conf)
	assert.Equal(t, filepath.Join(".", "data", "config.yml"), *mgr.file)
}
