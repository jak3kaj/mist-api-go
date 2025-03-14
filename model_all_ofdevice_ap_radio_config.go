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

type AllOfdeviceApRadioConfig struct {
	AllowRrmDisable bool `json:"allow_rrm_disable,omitempty"`
	// antenna gain for 2.4G - for models with external antenna only
	AntGain24 int32 `json:"ant_gain_24,omitempty"`
	// antenna gain for 5G - for models with external antenna only
	AntGain5 int32 `json:"ant_gain_5,omitempty"`
	// antenna gain for 6G - for models with external antenna only
	AntGain6 int32 `json:"ant_gain_6,omitempty"`
	AntennaMode *any `json:"antenna_mode,omitempty"`
	Band24 *any `json:"band_24,omitempty"`
	Band24Usage *any `json:"band_24_usage,omitempty"`
	Band5 *any `json:"band_5,omitempty"`
	Band5On24Radio *any `json:"band_5_on_24_radio,omitempty"`
	Band6 *any `json:"band_6,omitempty"`
	// to make an outdoor operate indoor. for an outdoor-ap, some channels are disallowed by default, this allows the user to use it as an indoor-ap
	IndoorUse bool `json:"indoor_use,omitempty"`
	// whether scanning radio is enabled
	ScanningEnabled bool `json:"scanning_enabled,omitempty"`
}
