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

type InventoryUpdate struct {
	// if `op`==`assign`, a **cloud-ready** switch/gateway will be managed/configured by Mist by default, this disabled the behavior
	DisableAutoConfig bool `json:"disable_auto_config,omitempty"`
	// if `op`==`assign`, `op`==`unassign`, `op`==`upgrade_to_mist`or `op`==`downgrade_to_jsi` , list of MAC, e.g. [\"5c5b350e0001\"]
	Macs []string `json:"macs,omitempty"`
	// if `op`==`assign`, an **adopted** switch/gateway will not be managed/configured by Mist by default, this enables the behavior
	Managed bool `json:"managed,omitempty"`
	// if `op`==`assign`, if true, treat site assignment against an already assigned AP as error
	NoReassign bool `json:"no_reassign,omitempty"`
	Op *AllOfinventoryUpdateOp `json:"op"`
	// if `op`==`delete`, list of serial numbers, e.g. [\"FXLH2015150025\"]
	Serials []string `json:"serials,omitempty"`
	// if `op`==`assign`, target site id
	SiteId string `json:"site_id,omitempty"`
}
