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

type AllOfroutingPolicyTermAction struct {
	Accept bool `json:"accept,omitempty"`
	AddCommunity []string `json:"add_community,omitempty"`
	// for SSR, hub decides how VRF routes are leaked on spoke
	AddTargetVrfs []string `json:"add_target_vrfs,omitempty"`
	// when used as export policy, optional
	Community []string `json:"community,omitempty"`
	// when used as export policy, optional. To exclude certain AS
	ExcludeAsPath []string `json:"exclude_as_path,omitempty"`
	ExcludeCommunity []string `json:"exclude_community,omitempty"`
	// when used as export policy, optional
	ExportCommunitites []string `json:"export_communitites,omitempty"`
	// optional, for an import policy, local_preference can be changed
	LocalPreference string `json:"local_preference,omitempty"`
	// when used as export policy, optional. By default, the local AS will be prepended, to change it
	PrependAsPath []string `json:"prepend_as_path,omitempty"`
}
