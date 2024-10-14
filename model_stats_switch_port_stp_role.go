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
// StatsSwitchPortStpRole : if `up`==`true`. enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
type StatsSwitchPortStpRole string

// List of stats_switch_port_stp_role
const (
	ALTERNATE_StatsSwitchPortStpRole StatsSwitchPortStpRole = "alternate"
	BACKUP_StatsSwitchPortStpRole StatsSwitchPortStpRole = "backup"
	DESIGNATED_StatsSwitchPortStpRole StatsSwitchPortStpRole = "designated"
	ROOT_StatsSwitchPortStpRole StatsSwitchPortStpRole = "root"
	ROOT_PREVENTED_StatsSwitchPortStpRole StatsSwitchPortStpRole = "root-prevented"
)
