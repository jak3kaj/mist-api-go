/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// GatewayPortUsage : port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan`
type GatewayPortUsage string

// List of gateway_port_usage
const (
	HA_CONTROL_GatewayPortUsage GatewayPortUsage = "ha_control"
	HA_DATA_GatewayPortUsage GatewayPortUsage = "ha_data"
	LAN_GatewayPortUsage GatewayPortUsage = "lan"
	WAN_GatewayPortUsage GatewayPortUsage = "wan"
)
