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

// Switch port config
type JunosPortConfig struct {
	// To disable LACP support for the AE interface
	AeDisableLacp bool `json:"ae_disable_lacp,omitempty"`
	// Users could force to use the designated AE name
	AeIdx int32 `json:"ae_idx,omitempty"`
	// to use fast timeout
	AeLacpSlow bool `json:"ae_lacp_slow,omitempty"`
	Aggregated bool `json:"aggregated,omitempty"`
	// if want to generate port up/down alarm
	Critical bool `json:"critical,omitempty"`
	Description string `json:"description,omitempty"`
	// if `speed` and `duplex` are specified, whether to disable autonegotiation
	DisableAutoneg bool `json:"disable_autoneg,omitempty"`
	Duplex *AllOfjunosPortConfigDuplex `json:"duplex,omitempty"`
	// Enable dynamic usage for this port. Set to `dynamic` to enable.
	DynamicUsage string `json:"dynamic_usage,omitempty"`
	Esilag bool `json:"esilag,omitempty"`
	// media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation
	Mtu int32 `json:"mtu,omitempty"`
	// prevent helpdesk to override the port config
	NoLocalOverwrite bool `json:"no_local_overwrite,omitempty"`
	PoeDisabled bool `json:"poe_disabled,omitempty"`
	Speed *AllOfjunosPortConfigSpeed `json:"speed,omitempty"`
	// port usage name.   If EVPN is used, use `evpn_uplink`or `evpn_downlink`
	Usage string `json:"usage"`
}
