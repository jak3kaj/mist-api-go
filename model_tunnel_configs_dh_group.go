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
// TunnelConfigsDhGroup : Only if `provider`== `custom-ipsec`. enum:   * 1   * 2 (1024-bit)   * 5   * 14 (default, 2048-bit)   * 15 (3072-bit)   * 16 (4096-bit)   * 19 (256-bit ECP)   * 20 (384-bit ECP)   * 21 (521-bit ECP)   * 24 (2048-bit ECP)
type TunnelConfigsDhGroup string

// List of tunnel_configs_dh_group
const (
	1__TunnelConfigsDhGroup TunnelConfigsDhGroup = "1"
	14__TunnelConfigsDhGroup TunnelConfigsDhGroup = "14"
	15__TunnelConfigsDhGroup TunnelConfigsDhGroup = "15"
	16__TunnelConfigsDhGroup TunnelConfigsDhGroup = "16"
	19__TunnelConfigsDhGroup TunnelConfigsDhGroup = "19"
	2__TunnelConfigsDhGroup TunnelConfigsDhGroup = "2"
	20__TunnelConfigsDhGroup TunnelConfigsDhGroup = "20"
	21__TunnelConfigsDhGroup TunnelConfigsDhGroup = "21"
	24__TunnelConfigsDhGroup TunnelConfigsDhGroup = "24"
	5__TunnelConfigsDhGroup TunnelConfigsDhGroup = "5"
)
