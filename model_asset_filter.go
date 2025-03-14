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

// Asset Filter
type AssetFilter struct {
	ApMac string `json:"ap_mac,omitempty"`
	Beam int32 `json:"beam,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	// whether the asset filter is disabled
	Disasbled bool `json:"disasbled,omitempty"`
	// eddystone uid namespace used to filter assets
	EddystoneUidNamespace string `json:"eddystone_uid_namespace,omitempty"`
	// eddystone url used to filter assets
	EddystoneUrl string `json:"eddystone_url,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// ibeacon major value used to filter assets
	IbeaconMajor int32 `json:"ibeacon_major,omitempty"`
	// ibeacon uuid used to filter assets
	IbeaconUuid string `json:"ibeacon_uuid,omitempty"`
	Id string `json:"id,omitempty"`
	// ble manufacturing-specific company-id used to filter assets
	MfgCompanyId int32 `json:"mfg_company_id,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	OrgId string `json:"org_id,omitempty"`
	Rssi int32 `json:"rssi,omitempty"`
	// ble service data uuid used to filter assets
	ServiceUuid string `json:"service_uuid,omitempty"`
	SiteId string `json:"site_id,omitempty"`
}
