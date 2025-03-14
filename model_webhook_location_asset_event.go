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

type WebhookLocationAssetEvent struct {
	BatteryVoltage int32 `json:"battery_voltage,omitempty"`
	EddystoneUidInstance string `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace string `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl string `json:"eddystone_url_url,omitempty"`
	IbeaconMajor int32 `json:"ibeacon_major,omitempty"`
	IbeaconMinor int32 `json:"ibeacon_minor,omitempty"`
	IbeaconUuid string `json:"ibeacon_uuid,omitempty"`
	Mac string `json:"mac,omitempty"`
	MapId string `json:"map_id,omitempty"`
	// optional, BLE manufacturing company ID
	MfgCompanyId int32 `json:"mfg_company_id,omitempty"`
	// optional, BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”)
	MfgData string `json:"mfg_data,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	Timestamp int32 `json:"timestamp,omitempty"`
	Type_ string `json:"type,omitempty"`
	// x, in meter
	X float64 `json:"x,omitempty"`
	// y, in meter
	Y float64 `json:"y,omitempty"`
}
