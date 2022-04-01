package bamboohr

// import (
// 	"github.com/turbot/steampipe-plugin-sdk/plugin"
// 	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
// )

// type bambooHrConfig struct {
// 	Token  *string `cty:"apiKey"`
// 	Tenant *string `cty:"tenant"`
// }

// var ConfigSchema = map[string]*schema.Attribute{
// 	"token": {
// 		Type: schema.TypeString,
// 	},
// }

// func ConfigInstance() interface{} {
// 	return &bambooHrConfig{}
// }

// // GetConfig :: retrieve and cast connection config from query data
// func GetConfig(connection *plugin.Connection) bambooHrConfig {
// 	if connection == nil || connection.Config == nil {
// 		return bambooHrConfig{}
// 	}
// 	config, _ := connection.Config.(bambooHrConfig)
// 	return config
// }
