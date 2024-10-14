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

type InlineResponse200137 struct {
	// proposal on band 2.4G, key is ap_id, value is the proposal
	Band24 map[string]RrmBand `json:"band_24"`
	Band24Metric *RrmBandMetric `json:"band_24_metric"`
	// proposal on band 5G, key is ap_id, value is the proposal
	Band5 map[string]RrmBand `json:"band_5"`
	Band5Metric *RrmBandMetric `json:"band_5_metric"`
	// proposal on band 6G, key is ap_id, value is the proposal
	Band6 map[string]RrmBand `json:"band_6,omitempty"`
	Band6Metric *RrmBandMetric `json:"band_6_metric,omitempty"`
	Rftemplate *any `json:"rftemplate"`
	RftemplateId string `json:"rftemplate_id"`
	RftemplateName string `json:"rftemplate_name"`
	Status *any `json:"status"`
	// time where the status was updated
	Timestamp float64 `json:"timestamp"`
}
