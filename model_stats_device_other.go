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

type StatsDeviceOther struct {
	ConfigStatus string `json:"config_status,omitempty"`
	LastConfig int32 `json:"last_config,omitempty"`
	LastSeen int32 `json:"last_seen,omitempty"`
	Mac string `json:"mac,omitempty"`
	Status string `json:"status,omitempty"`
	Uptime int32 `json:"uptime,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	VendorSpecific *AllOfstatsDeviceOtherVendorSpecific `json:"vendor_specific,omitempty"`
	Version string `json:"version,omitempty"`
}
