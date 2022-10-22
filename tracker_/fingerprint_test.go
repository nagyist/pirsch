package tracker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFingerprint(t *testing.T) {
	SetFingerprintKeys(42, 99)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("User-Agent", "test")
	req.RemoteAddr = "127.0.0.1:80"
	assert.NotZero(t, Fingerprint(req, "salt", time.Now().UTC(), nil, nil))
}