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

type OrgIdRftemplatesBody struct {
	AntGain24 int32 `json:"ant_gain_24,omitempty"`
	AntGain5 int32 `json:"ant_gain_5,omitempty"`
	AntGain6 int32 `json:"ant_gain_6,omitempty"`
	Band24 *any `json:"band_24,omitempty"`
	Band24Usage *any `json:"band_24_usage,omitempty"`
	Band5 *any `json:"band_5,omitempty"`
	Band5On24Radio *any `json:"band_5_on_24_radio,omitempty"`
	Band6 *any `json:"band_6,omitempty"`
	// optional, country code to use. If specified, this gets applied to all sites using the RF Template
	CountryCode string `json:"country_code,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Id string `json:"id,omitempty"`
	// overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. \"AP63\")
	ModelSpecific map[string]RfTemplateModelSpecificProperty `json:"model_specific,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// The name of the RF template
	Name string `json:"name"`
	OrgId string `json:"org_id,omitempty"`
	// whether scanning radio is enabled
	ScanningEnabled bool `json:"scanning_enabled,omitempty"`
}
