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
// SearchSiteSwOrGwPortsDeviceType : enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch`
type SearchSiteSwOrGwPortsDeviceType string

// List of search_site_sw_or_gw_ports_device_type
const (
	AP_SearchSiteSwOrGwPortsDeviceType SearchSiteSwOrGwPortsDeviceType = "ap"
	BLE_SearchSiteSwOrGwPortsDeviceType SearchSiteSwOrGwPortsDeviceType = "ble"
	GATEWAY_SearchSiteSwOrGwPortsDeviceType SearchSiteSwOrGwPortsDeviceType = "gateway"
	MXEDGE_SearchSiteSwOrGwPortsDeviceType SearchSiteSwOrGwPortsDeviceType = "mxedge"
	NAC_SearchSiteSwOrGwPortsDeviceType SearchSiteSwOrGwPortsDeviceType = "nac"
	SWITCH__SearchSiteSwOrGwPortsDeviceType SearchSiteSwOrGwPortsDeviceType = "switch"
)
