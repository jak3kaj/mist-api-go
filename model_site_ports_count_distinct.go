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
// SitePortsCountDistinct : enum: `full_duplex`, `mac`, `neighbor_mac`, `neighbor_port_desc`, `neighbor_system_name`, `poe_disabled`, `poe_mode`, `poe_on`, `port_id`, `port_mac`, `speed`, `up`
type SitePortsCountDistinct string

// List of site_ports_count_distinct
const (
	FULL_DUPLEX_SitePortsCountDistinct SitePortsCountDistinct = "full_duplex"
	MAC_SitePortsCountDistinct SitePortsCountDistinct = "mac"
	NEIGHBOR_MAC_SitePortsCountDistinct SitePortsCountDistinct = "neighbor_mac"
	NEIGHBOR_PORT_DESC_SitePortsCountDistinct SitePortsCountDistinct = "neighbor_port_desc"
	NEIGHBOR_SYSTEM_NAME_SitePortsCountDistinct SitePortsCountDistinct = "neighbor_system_name"
	POE_DISABLED_SitePortsCountDistinct SitePortsCountDistinct = "poe_disabled"
	POE_MODE_SitePortsCountDistinct SitePortsCountDistinct = "poe_mode"
	POE_ON_SitePortsCountDistinct SitePortsCountDistinct = "poe_on"
	PORT_ID_SitePortsCountDistinct SitePortsCountDistinct = "port_id"
	PORT_MAC_SitePortsCountDistinct SitePortsCountDistinct = "port_mac"
	SPEED_SitePortsCountDistinct SitePortsCountDistinct = "speed"
	UP_SitePortsCountDistinct SitePortsCountDistinct = "up"
)
