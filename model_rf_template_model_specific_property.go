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

type RfTemplateModelSpecificProperty struct {
	AntGain24 int32 `json:"ant_gain_24,omitempty"`
	AntGain5 int32 `json:"ant_gain_5,omitempty"`
	AntGain6 int32 `json:"ant_gain_6,omitempty"`
	Band24 *AllOfrfTemplateModelSpecificPropertyBand24 `json:"band_24,omitempty"`
	Band24Usage *AllOfrfTemplateModelSpecificPropertyBand24Usage `json:"band_24_usage,omitempty"`
	Band5 *AllOfrfTemplateModelSpecificPropertyBand5 `json:"band_5,omitempty"`
	Band5On24Radio *AllOfrfTemplateModelSpecificPropertyBand5On24Radio `json:"band_5_on_24_radio,omitempty"`
	Band6 *AllOfrfTemplateModelSpecificPropertyBand6 `json:"band_6,omitempty"`
}
