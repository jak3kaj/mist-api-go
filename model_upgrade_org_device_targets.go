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

type UpgradeOrgDeviceTargets struct {
	// list of devices MAC Addresses which cloud has requested to download firmware
	DownloadRequested []string `json:"download_requested,omitempty"`
	// list of devices MAC Addresses which have the firmware downloaded
	Downloaded []string `json:"downloaded,omitempty"`
	// list of devices MAC Addresses which have failed to upgrade
	Failed []string `json:"failed,omitempty"`
	// list of devices MAC Addresses which are rebooting
	RebootInProgress []string `json:"reboot_in_progress,omitempty"`
	// list of devices MAC Addresses which have rebooted successfully
	Rebooted []string `json:"rebooted,omitempty"`
	// list of devices MAC Addresses which cloud has scheduled an upgrade for
	Scheduled []string `json:"scheduled,omitempty"`
	// list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade
	Skipped []string `json:"skipped,omitempty"`
	// count of devices part of this upgrade
	Total int32 `json:"total,omitempty"`
	// count of devices which have upgraded successfully
	Upgraded []string `json:"upgraded,omitempty"`
}
