package main

import (
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/robertlagrant/steampipe-plugin-bamboohr/bamboohr"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: bamboohr.Plugin})
}
