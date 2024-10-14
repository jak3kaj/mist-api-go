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

type OrgIdWxrulesBody struct {
	Action *any `json:"action,omitempty"`
	ApplyTags []string `json:"apply_tags,omitempty"`
	// blocked apps (always blocking, ignoring action), the key of Get Application List
	BlockedApps []string `json:"blocked_apps,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	// List of WxTag UUID to indicate these tags are allowed access
	DstAllowWxtags []string `json:"dst_allow_wxtags"`
	// List of WxTag UUID to indicate these tags are blocked access
	DstDenyWxtags []string `json:"dst_deny_wxtags"`
	// List of WxTag UUID
	DstWxtags []string `json:"dst_wxtags,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Id string `json:"id,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// the order how rules would be looked up, > 0 and bigger order got matched first, -1 means LAST, uniqueness not checked
	Order int32 `json:"order"`
	OrgId string `json:"org_id,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// List of WxTag UUID to determine if this rule would match
	SrcWxtags []string `json:"src_wxtags"`
	// Only for Org Level WxRule
	TemplateId string `json:"template_id,omitempty"`
}
