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

type DevicesDeviceMacBody struct {
	DeviceprofileName string `json:"deviceprofile_name,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Height float64 `json:"height,omitempty"`
	MapId string `json:"map_id,omitempty"`
	Name string `json:"name"`
	Orientation int32 `json:"orientation,omitempty"`
	// Onlif this is to replace an existing device
	ReplacingMac string `json:"replacing_mac,omitempty"`
	// optional role for switch / gateway
	Role string `json:"role,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	SiteName string `json:"site_name,omitempty"`
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}
