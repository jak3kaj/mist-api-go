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
// NacTagType : enum: `egress_vlan_names`, `gbp_tag`, `match`, `radius_attrs`, `radius_group`, `radius_vendor_attrs`, `session_timeout`, `username_attr`, `vlan`
type NacTagType string

// List of nac_tag_type
const (
	EGRESS_VLAN_NAMES_NacTagType NacTagType = "egress_vlan_names"
	GBP_TAG_NacTagType NacTagType = "gbp_tag"
	MATCH_NacTagType NacTagType = "match"
	RADIUS_ATTRS_NacTagType NacTagType = "radius_attrs"
	RADIUS_GROUP_NacTagType NacTagType = "radius_group"
	RADIUS_VENDOR_ATTRS_NacTagType NacTagType = "radius_vendor_attrs"
	SESSION_TIMEOUT_NacTagType NacTagType = "session_timeout"
	USERNAME_ATTR_NacTagType NacTagType = "username_attr"
	VLAN_NacTagType NacTagType = "vlan"
)
