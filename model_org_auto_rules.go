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

// auto_rules in org settings
type OrgAutoRules struct {
	// if `src`==`geoip`. By default, a claimed device only gets assigned if the site exists to auto-create the site, enable this
	CreateNewSiteIfNeeded bool `json:"create_new_site_if_needed,omitempty"`
	// if `src`==`name`, `src`==`lldp_system_name`,  `src`==`dns_suffix`       \"[0:3]\"            // \"abcdef\" -> \"abc\"       \"split(.)[1]\"      // \"a.b.c\" -> \"b\"       \"split(-)[1][0:3]\" // \"a1234-b5678-c90\" -> \"b56\"'
	Expression string `json:"expression,omitempty"`
	// if `src`==`geoip` and `create_new_site_if_needed`==`true`. If a gateway template is desired for this newly created site
	GatewaytemplateId string `json:"gatewaytemplate_id,omitempty"`
	// if `src`==`geoip`
	MatchCountry string `json:"match_country,omitempty"`
	MatchDeviceType *AllOforgAutoRulesMatchDeviceType `json:"match_device_type,omitempty"`
	// optional/additional filter
	MatchModel string `json:"match_model,omitempty"`
	// if `src`==`model`
	Model string `json:"model,omitempty"`
	// if `src`==`name`
	Prefix string `json:"prefix,omitempty"`
	Src *AllOforgAutoRulesSrc `json:"src"`
	// if `src`==`subnet`
	Subnet string `json:"subnet,omitempty"`
	// if `src`==`name`
	Suffix string `json:"suffix,omitempty"`
	// if      * `src`==`model`     *  `src`==`geoip: site name for the device to be assigned to (\"city\" / \"city+country\" / ...)
	Value string `json:"value,omitempty"`
}
