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

type ApUplinkPortConfig struct {
	// Whether to do 802.1x against uplink switch. When enaled, AP cert will be used to do EAP-TLS and the Org's CA Cert has to be provisioned at the switch
	Dot1x bool `json:"dot1x,omitempty"`
	// by default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons.
	KeepWlansUpIfDown bool `json:"keep_wlans_up_if_down,omitempty"`
}
