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

type InlineResponse20072 struct {
	// can be set to true to allow the override by usermac result
	AllowUsermacOverride bool `json:"allow_usermac_override,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	// if `type`==`egress_vlan_names`, list of egress vlans to return
	EgressVlanNames []string `json:"egress_vlan_names,omitempty"`
	// if `type`==`gbp_tag`
	GbpTag int32 `json:"gbp_tag,omitempty"`
	Id string `json:"id,omitempty"`
	Match *any `json:"match,omitempty"`
	// This field is applicable only when `type`==`match`   * `false`: means it is sufficient to match any of the values (i.e., match-any behavior)   * `true`: means all values should be matched (i.e., match-all behavior)   Currently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`
	MatchAll bool `json:"match_all,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	OrgId string `json:"org_id,omitempty"`
	// if `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field \"radius_attrs\".  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_attrs in the result of a given rule.
	RadiusAttrs []string `json:"radius_attrs,omitempty"`
	// if `type`==`radius_group`
	RadiusGroup string `json:"radius_group,omitempty"`
	// if `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field \"radius_vendor_attrs\".  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule.
	RadiusVendorAttrs []string `json:"radius_vendor_attrs,omitempty"`
	// if `type`==`session_timeout, in seconds
	SessionTimeout int32 `json:"session_timeout,omitempty"`
	Type_ *any `json:"type"`
	UsernameAttr *any `json:"username_attr,omitempty"`
	// if `type`==`match`
	Values []string `json:"values,omitempty"`
	// if `type`==`vlan`
	Vlan string `json:"vlan,omitempty"`
}
