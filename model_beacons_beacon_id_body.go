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

type BeaconsBeaconIdBody struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	// Eddystone-UID instance (6 bytes) in hexstring format
	EddystoneInstance string `json:"eddystone_instance,omitempty"`
	// Eddystone-UID namespace (10 bytes) in hexstring format
	EddystoneNamespace string `json:"eddystone_namespace,omitempty"`
	// Eddystone-URL url
	EddystoneUrl string `json:"eddystone_url,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// bluetooth tag major
	IbeaconMajor int32 `json:"ibeacon_major,omitempty"`
	// bluetooth tag minor
	IbeaconMinor int32 `json:"ibeacon_minor,omitempty"`
	// bluetooth tag UUID
	IbeaconUuid string `json:"ibeacon_uuid,omitempty"`
	Id string `json:"id,omitempty"`
	// optiona, MAC of the beacon, currently used only to identify battery voltage
	Mac string `json:"mac,omitempty"`
	// map where the device belongs to
	MapId string `json:"map_id,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// name / label of the device
	Name string `json:"name,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// in dBm
	Power int32 `json:"power,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	Type_ *Object `json:"type,omitempty"`
	// x in pixel
	X float64 `json:"x,omitempty"`
	// y in pixel
	Y float64 `json:"y,omitempty"`
}
