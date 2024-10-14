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

type InlineResponse200168 struct {
	// battery voltage, in mV
	BatteryVoltage float64 `json:"battery_voltage,omitempty"`
	Beam int32 `json:"beam,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	EddystoneUidInstance string `json:"eddystone_uid_instance,omitempty"`
	EddystoneUidNamespace string `json:"eddystone_uid_namespace,omitempty"`
	EddystoneUrlUrl string `json:"eddystone_url_url,omitempty"`
	IbeaconMajor int32 `json:"ibeacon_major,omitempty"`
	IbeaconMinor int32 `json:"ibeacon_minor,omitempty"`
	IbeaconUuid string `json:"ibeacon_uuid,omitempty"`
	// last seen timestamp
	LastSeen float64 `json:"last_seen,omitempty"`
	// bluetooth MAC
	Mac string `json:"mac"`
	// map where the device belongs to
	MapId string `json:"map_id,omitempty"`
	// name / label of the device
	Name string `json:"name,omitempty"`
	Rssi int32 `json:"rssi,omitempty"`
	// only send this for individual asset stat
	Rssizones []AssetRssiZone `json:"rssizones,omitempty"`
	Temperatur float64 `json:"temperatur,omitempty"`
	// x in pixel
	X float64 `json:"x,omitempty"`
	// y in pixel
	Y float64 `json:"y,omitempty"`
	// only send this for individual asset stat
	Zones []AssetZone `json:"zones,omitempty"`
}
