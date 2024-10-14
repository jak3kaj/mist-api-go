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

type DeviceIdUpgradeBody struct {
	// Reboot device immediately after upgrade is completed (Available on Junos OS devices)
	Reboot bool `json:"reboot,omitempty"`
	// reboot start time in epoch
	RebootAt int32 `json:"reboot_at,omitempty"`
	// Perform recovery snapshot after device is rebooted (Available on Junos OS devices)
	Snapshot bool `json:"snapshot,omitempty"`
	// firmware download start time in epoch
	StartTime float64 `json:"start_time,omitempty"`
	// specific version / `stable`, default is to use the latest
	Version string `json:"version"`
}
