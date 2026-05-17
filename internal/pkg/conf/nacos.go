package conf

import (
	"cmp"
	"os"

	"github.com/chenjl-ops/go-lib/nacos"
	log "github.com/sirupsen/logrus"
)

var NacosConfig *Specification

func NacosReadRemoteConfig() error {
	nc, err := nacos.NewNacosConfig(
		nacos.WithUrl(cmp.Or(os.Getenv("RUNTIME_CONFIG_URL"), "http://10.1.16.12")),
		nacos.WithDataId(cmp.Or(os.Getenv("RUNTIME_APP_NAME"), "op-wx-api")),
		nacos.WithPath("/nacos"),
		nacos.WithPort(8848),
		nacos.WithGroup(cmp.Or(os.Getenv("RUNTIME_GROUP"), "ops")),
		nacos.WithTenant(cmp.Or(os.Getenv("RUNTIME_ENV"), "public")),
	)

	if err != nil {
		log.Error("NacosReadRemoteConfig error: ", err)
		return err
	}

	return nc.ReadRemoteConfig(&NacosConfig)
}
