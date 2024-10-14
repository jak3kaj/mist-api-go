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

// Junos IP Config
type GatewayPortConfigIpConfig struct {
	// except for out-of_band interface (vme/em0/fxp0)
	Dns []string `json:"dns,omitempty"`
	// except for out-of_band interface (vme/em0/fxp0)
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// except for out-of_band interface (vme/em0/fxp0)
	Gateway string `json:"gateway,omitempty"`
	Ip string `json:"ip,omitempty"`
	// used only if `subnet` is not specified in `networks`
	Netmask string `json:"netmask,omitempty"`
	// optional, the network to be used for mgmt
	Network string `json:"network,omitempty"`
	// if `type`==`pppoe`
	PoserPassword string `json:"poser_password,omitempty"`
	PppoeAuth *AllOfgatewayPortConfigIpConfigPppoeAuth `json:"pppoe_auth,omitempty"`
	// if `type`==`pppoe`
	PppoeUsername string `json:"pppoe_username,omitempty"`
	Type_ *AllOfgatewayPortConfigIpConfigType_ `json:"type,omitempty"`
}
