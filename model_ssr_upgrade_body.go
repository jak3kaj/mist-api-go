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

type SsrUpgradeBody struct {
	Channel *Object `json:"channel,omitempty"`
	// list of 128T device IDs to upgrade
	DeviceIds []string `json:"device_ids"`
	// reboot start time in epoch seconds, default is start_time, -1 disables reboot
	RebootAt int32 `json:"reboot_at,omitempty"`
	// 128T firmware download start time in epoch seconds, default is now, -1 disables download
	StartTime int32 `json:"start_time,omitempty"`
	Strategy *Object `json:"strategy,omitempty"`
	// 128T firmware version to upgrade (e.g. 5.3.0-93)
	Version string `json:"version,omitempty"`
}
