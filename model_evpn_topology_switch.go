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

type EvpnTopologySwitch struct {
	Config *AllOfevpnTopologySwitchConfig `json:"config,omitempty"`
	DeviceprofileId string `json:"deviceprofile_id,omitempty"`
	Downlinks []string `json:"downlinks,omitempty"`
	Esilaglinks []string `json:"esilaglinks,omitempty"`
	EvpnId int32 `json:"evpn_id,omitempty"`
	Mac string `json:"mac,omitempty"`
	Model string `json:"model,omitempty"`
	// optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.    * for CLOS, to group dist / access switches into pods   * for ERB/CRB, to group dist / esilag-access into pods
	Pod int32 `json:"pod,omitempty"`
	// by default, core switches are assumed to be connecting all pods.  if you want to limit the pods, you can specify pods.
	Pods []int32 `json:"pods,omitempty"`
	Role *AllOfevpnTopologySwitchRole `json:"role,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	SuggestedDownlinks []string `json:"suggested_downlinks,omitempty"`
	SuggestedEsilaglinks []string `json:"suggested_esilaglinks,omitempty"`
	SuggestedUplinks []string `json:"suggested_uplinks,omitempty"`
	// if not specified in the request, suggested ones will be used
	Uplinks []string `json:"uplinks,omitempty"`
}
