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

type WebhookSdkclientScanDataEventScanDataItem struct {
	// mac address of the AP associated with the BSSID scanned
	Ap string `json:"ap"`
	Band *AllOfwebhookSdkclientScanDataEventScanDataItemBand `json:"band"`
	// BSSID found during client’s background scan for wifi
	Bssid string `json:"bssid"`
	// channel of the band found in the scan
	Channel int32 `json:"channel"`
	// client’s RSSI relative to the BSSID scanned
	Rssi float64 `json:"rssi"`
	// ESSID containing the BSSID scanned
	Ssid string `json:"ssid"`
	// time the scan of the particular BSSID occurred
	Timestamp float64 `json:"timestamp"`
}
