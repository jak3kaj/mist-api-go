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

type AllOfrfTemplateModelSpecificPropertyBand6 struct {
	AllowRrmDisable bool `json:"allow_rrm_disable,omitempty"`
	AntGain int32 `json:"ant_gain,omitempty"`
	AntennaMode *any `json:"antenna_mode,omitempty"`
	Bandwidth *any `json:"bandwidth,omitempty"`
	// For RFTemplates. List of channels, null or empty array means auto
	Channels []int32 `json:"channels,omitempty"`
	// whether to disable the radio
	Disabled bool `json:"disabled,omitempty"`
	// TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …
	Power int32 `json:"power,omitempty"`
	// when power=0, max tx power to use, HW-specific values will be used if not set
	PowerMax int32 `json:"power_max,omitempty"`
	// when power=0, min tx power to use, HW-specific values will be used if not set
	PowerMin int32 `json:"power_min,omitempty"`
	Preamble *any `json:"preamble,omitempty"`
	// for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we'll fallback to Low Power Indoor if AFC failed
	StandardPower bool `json:"standard_power,omitempty"`
}
