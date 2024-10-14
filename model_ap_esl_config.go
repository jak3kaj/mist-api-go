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

type ApEslConfig struct {
	// Only if `type`==`imagotag` or `type`==`native`
	Cacert string `json:"cacert,omitempty"`
	// Only if `type`==`imagotag` or `type`==`native`
	Channel int32 `json:"channel,omitempty"`
	// usb_config is ignored if esl_config enabled
	Enabled bool `json:"enabled,omitempty"`
	// Only if `type`==`imagotag` or `type`==`native`
	Host string `json:"host,omitempty"`
	// Only if `type`==`imagotag` or `type`==`native`
	Port int32 `json:"port,omitempty"`
	Type_ *AllOfapEslConfigType_ `json:"type,omitempty"`
	// Only if `type`==`imagotag` or `type`==`native`
	VerifyCert bool `json:"verify_cert,omitempty"`
	// Only if `type`==`solum` or `type`==`hanshow`
	VlanId int32 `json:"vlan_id,omitempty"`
}
