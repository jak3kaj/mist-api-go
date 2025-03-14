/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mist

type VbeaconsVbeaconIdBody struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Id string `json:"id,omitempty"`
	// bluetooth tag major
	Major int32 `json:"major,omitempty"`
	// map where the device belongs to
	MapId string `json:"map_id,omitempty"`
	// a message that can be displayed when the sdkclient gets near the vbeacon
	Message string `json:"message,omitempty"`
	// bluetooth tag minor
	Minor int32 `json:"minor,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// name / label of the device
	Name string `json:"name,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// required if `power_mode`==`custom`, -30 - 100, in dBm. For default power_mode, power = 4 dBm.
	Power int32 `json:"power,omitempty"`
	PowerMode *any `json:"power_mode,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// URL to show, optional
	Url string `json:"url,omitempty"`
	// bluetooth tag UUID
	Uuid string `json:"uuid,omitempty"`
	// the name to be used in wayfinding_path or wayfinding_grid blob
	WayfindingNodename string `json:"wayfinding_nodename,omitempty"`
	// x in pixel
	X float64 `json:"x,omitempty"`
	// y in pixel
	Y float64 `json:"y,omitempty"`
}
