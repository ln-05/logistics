package inits

import (
	"encoding/json"
	"log"
	"logistics_srv/basic/appconfig"
	"logistics_srv/basic/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitNacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         appconfig.AppConf.NamespaceID, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      appconfig.AppConf.IdAddr,
			ContextPath: "/nacos",
			Port:        uint64(appconfig.AppConf.Port),
			Scheme:      "http",
		},
	}

	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: appconfig.AppConf.DataId,
		Group:  appconfig.AppConf.Group})
	log.Println(content)
	var my appconfig.Config
	json.Unmarshal([]byte(content), &my)
	global.Nacos = my
	log.Println(global.Nacos)
}
