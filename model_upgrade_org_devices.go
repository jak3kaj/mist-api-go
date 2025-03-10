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

type UpgradeOrgDevices struct {
	// phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase.
	CanaryPhases []int32 `json:"canary_phases,omitempty"`
	// whether to allow local AP-to-AP FW upgrade
	EnableP2p bool `json:"enable_p2p,omitempty"`
	// true will force upgrade when requested version is same as running version
	Force bool `json:"force,omitempty"`
	// percentage of failures allowed across the entire upgrade(not applicable for `big_bang`)
	MaxFailurePercentage float64 `json:"max_failure_percentage,omitempty"`
	// number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used
	MaxFailures []int32 `json:"max_failures,omitempty"`
	Models []string `json:"models,omitempty"`
	P2pClusterSize int32 `json:"p2p_cluster_size,omitempty"`
	// number of parallel p2p download batches to creat
	P2pParallelism int32 `json:"p2p_parallelism,omitempty"`
	// Reboot device immediately after upgrade is completed (Available on Junos OS devices)
	Reboot bool `json:"reboot,omitempty"`
	// reboot start time in epoch seconds, default is `start_time`
	RebootAt float64 `json:"reboot_at,omitempty"`
	// percentage of AP’s that need to be present in the first rrm batch
	RrmFirstBatchPercentage int32 `json:"rrm_first_batch_percentage,omitempty"`
	// max percentage of AP’s that need to be present in each rrm batch
	RrmMaxBatchPercentage int32 `json:"rrm_max_batch_percentage,omitempty"`
	RrmMeshUpgrade *AllOfupgradeOrgDevicesRrmMeshUpgrade `json:"rrm_mesh_upgrade,omitempty"`
	RrmNodeOrder *AllOfupgradeOrgDevicesRrmNodeOrder `json:"rrm_node_order,omitempty"`
	// true will make rrm batch sizes slowly ramp up
	RrmSlowRamp bool `json:"rrm_slow_ramp,omitempty"`
	SiteIds []string `json:"site_ids,omitempty"`
	// Perform recovery snapshot after device is rebooted (Available on Junos OS devices)
	Snapshot bool `json:"snapshot,omitempty"`
	// upgrade start time in epoch seconds, default is now
	StartTime float64 `json:"start_time,omitempty"`
	Strategy *AllOfupgradeOrgDevicesStrategy `json:"strategy,omitempty"`
	// specific version / stable
	Version string `json:"version,omitempty"`
}
