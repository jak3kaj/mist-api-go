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

// additional information per alarm type
type Alarm struct {
	// UUID of the admin who acked the alarm
	AckAdminId string `json:"ack_admin_id,omitempty"`
	// Name & Email ID of the admin who acked the alarm
	AckAdminName string `json:"ack_admin_name,omitempty"`
	// Whether the alarm is acked or not
	Acked bool `json:"acked,omitempty"`
	// Epoch (seconds) when the alarm was acked
	AckedTime int32 `json:"acked_time,omitempty"`
	// additional information: List of MACs of the APs
	Aps []string `json:"aps,omitempty"`
	// List of BSSIDs
	Bssids []string `json:"bssids,omitempty"`
	// Number of incident within an alarm window
	Count int32 `json:"count"`
	// additional information: List of MACs of the gateways
	Gateways []string `json:"gateways,omitempty"`
	// Group of the alarm
	Group string `json:"group"`
	// additional information: List of Hostnames of the devices (AP/Switch/Gateway)
	Hostnames []string `json:"hostnames,omitempty"`
	// UUID of the alarm
	Id string `json:"id"`
	// Epoch (seconds) of the last incident/alarm within an alarm window
	LastSeen float64 `json:"last_seen"`
	// Text describing the alarm
	Note string `json:"note,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// Severity of the alarm
	Severity string `json:"severity"`
	SiteId string `json:"site_id,omitempty"`
	// List of SSIDs
	Ssids []string `json:"ssids,omitempty"`
	// additional information: List of MACs of the switches
	Switches []string `json:"switches,omitempty"`
	// Epoch (seconds) of the first incident/alarm
	Timestamp int32 `json:"timestamp"`
	// Key-name of the alarm type
	Type_ string `json:"type"`
}
