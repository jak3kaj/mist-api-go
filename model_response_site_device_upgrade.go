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

type ResponseSiteDeviceUpgrade struct {
	Counts *AllOfresponseSiteDeviceUpgradeCounts `json:"counts,omitempty"`
	// current canary or rrm phase in progress
	CurrentPhase int32 `json:"current_phase,omitempty"`
	// whether to allow local AP-to-AP FW upgrade
	EnableP2p bool `json:"enable_p2p,omitempty"`
	// whether to force upgrade when requested version is same as running version
	Force bool `json:"force,omitempty"`
	// unique id for the upgrade
	Id string `json:"id"`
	// percentage of failures allowed
	MaxFailurePercentage int32 `json:"max_failure_percentage,omitempty"`
	// number of failures allowed within a canary phase or serial rollout
	MaxFailures []int32 `json:"max_failures,omitempty"`
	// reboot start time in epoch
	RebootAt int32 `json:"reboot_at,omitempty"`
	// firmware download start time in epoch
	StartTime float64 `json:"start_time,omitempty"`
	Status *AllOfresponseSiteDeviceUpgradeStatus `json:"status,omitempty"`
	Strategy *AllOfresponseSiteDeviceUpgradeStrategy `json:"strategy,omitempty"`
	// version to upgrade to
	TargetVersion string `json:"target_version,omitempty"`
	// a dictionary of rrm phase number to devices part of that phase
	UpgradePlan *interface{} `json:"upgrade_plan,omitempty"`
}
