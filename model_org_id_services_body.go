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

type OrgIdServicesBody struct {
	// if `type`==`custom`, ip subnets (e.g. 10.0.0.0/8)
	Addresses []string `json:"addresses,omitempty"`
	// when `type`==`app_categories`, list of application categories are available through /api/v1/const/app_categories
	AppCategories []string `json:"app_categories,omitempty"`
	// when `type`==`app_categories`, list of application categories are available through /api/v1/const/app_subcategories
	AppSubcategories []string `json:"app_subcategories,omitempty"`
	// when `type`==`apps`, list of applications are available through:   * /api/v1/const/applications   * /api/v1/const/gateway_applications   * /insight/top_app_by-bytes?wired=true
	Apps []string `json:"apps,omitempty"`
	// 0 means unlimited
	ClientLimitDown int32 `json:"client_limit_down,omitempty"`
	// 0 means unlimited
	ClientLimitUp int32 `json:"client_limit_up,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	Description string `json:"description,omitempty"`
	Dscp *Object `json:"dscp,omitempty"`
	FailoverPolicy *Object `json:"failover_policy,omitempty"`
	// if `type`==`custom`, web filtering
	Hostnames []string `json:"hostnames,omitempty"`
	Id string `json:"id,omitempty"`
	MaxJitter *Object `json:"max_jitter,omitempty"`
	MaxLatency *Object `json:"max_latency,omitempty"`
	MaxLoss *Object `json:"max_loss,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// 0 means unlimited
	ServiceLimitDown int32 `json:"service_limit_down,omitempty"`
	// 0 means unlimited
	ServiceLimitUp int32 `json:"service_limit_up,omitempty"`
	// whether to enable measure SLE
	SleEnabled bool `json:"sle_enabled,omitempty"`
	// when `type`==`custom`, optional, if it doesn't exist, http and https is assumed
	Specs []ServiceSpec `json:"specs,omitempty"`
	SsrRelaxedTcpStateEnforcement bool `json:"ssr_relaxed_tcp_state_enforcement,omitempty"`
	TrafficClass *Object `json:"traffic_class,omitempty"`
	// values from `/api/v1/consts/traffic_types`
	TrafficType string `json:"traffic_type,omitempty"`
	Type_ *Object `json:"type,omitempty"`
	// when `type`==`urls`, no need for spec as URL can encode the ports being used
	Urls []string `json:"urls,omitempty"`
}
