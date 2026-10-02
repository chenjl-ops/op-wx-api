package conf

import (
	"cmp"
	"fmt"
	"os"
	"strconv"

	"github.com/chenjl-ops/go-lib/nacos"
	log "github.com/sirupsen/logrus"
)

var NacosConfig *Specification

func NacosReadRemoteConfig() error {
	fmt.Println(
		os.Getenv("RUNTIME_NACOS_USERNAME"),
		os.Getenv("RUNTIME_CONFIG_URL"),
		os.Getenv("RUNTIME_NACOS_SCHEME"),
		getUint64Env("RUNTIME_NACOS_PORT", 8848),
	)

	nc, err := nacos.NewNacosConfig(
		nacos.WithUrl(cmp.Or(os.Getenv("RUNTIME_CONFIG_URL"), "10.1.16.12")),
		nacos.WithDataId(cmp.Or(os.Getenv("RUNTIME_APP_NAME"), "op-wx-api")),
		nacos.WithPath("/nacos"),
		nacos.WithPort(getUint64Env("RUNTIME_NACOS_PORT", 8848)),
		nacos.WithGroup(cmp.Or(os.Getenv("RUNTIME_GROUP"), "ops")),
		nacos.WithTenant(cmp.Or(os.Getenv("RUNTIME_ENV"), "public")),
		nacos.WithUserName(cmp.Or(os.Getenv("RUNTIME_NACOS_USERNAME"), "nacos")),
		nacos.WithPassword(cmp.Or(os.Getenv("RUNTIME_NACOS_PASSWORD"), "nacos")),
		nacos.WithScheme(cmp.Or(os.Getenv("RUNTIME_NACOS_SCHEME"), "http")),
	)

	if err != nil {
		log.Error("NacosReadRemoteConfig error: ", err)
		return err
	}

	return nc.ReadRemoteConfig(&NacosConfig)
}

// 处理环境变量获取的string => uint64 同时支持default
func getUint64Env(key string, defaultValue uint64) uint64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return n
}
