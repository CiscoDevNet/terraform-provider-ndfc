package nd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSetRaw tests the Body::SetRaw method.
func TestSetRaw(t *testing.T) {
	name := Body{}.SetRaw("a", `{"name":"a"}`).Res().Get("a.name").Str
	assert.Equal(t, "a", name)
}

// TestDelete tests the Body::Delete method.
func TestDelete(t *testing.T) {
	body := Body{}
	body = body.SetRaw("a", `{"name":"a"}`)
	assert.Equal(t, "a", body.Res().Get("a.name").Str)
	body = body.Delete("a.name")
	assert.Equal(t, "", body.Res().Get("a.name").Str)
}
