package weichat

import (
	"op-wx-api/internal/pkg/conf"
	"testing"

	log "github.com/sirupsen/logrus"
)

// TestGetAccessToken
func TestGetAccessToken(t *testing.T) {
	var err error
	err = conf.NacosReadRemoteConfig()
	if nil != err {
		log.Fatal(err)
	}

	_, err = getAccessToken()
	if err != nil {
		t.Fatal(err)
	}
}
