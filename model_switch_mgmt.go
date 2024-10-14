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

// Switch settings
type SwitchMgmt struct {
	// ap_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used.
	ApAffinityThreshold int32 `json:"ap_affinity_threshold,omitempty"`
	// Set Banners for switches. Allows markup formatting
	CliBanner string `json:"cli_banner,omitempty"`
	// Sets timeout for switches
	CliIdleTimeout int32 `json:"cli_idle_timeout,omitempty"`
	// the rollback timer for commit confirmed
	ConfigRevertTimer int32 `json:"config_revert_timer,omitempty"`
	// Enable to provide the FQDN with DHCP option 81
	DhcpOptionFqdn bool `json:"dhcp_option_fqdn,omitempty"`
	DisableOobDownAlarm bool `json:"disable_oob_down_alarm,omitempty"`
	// Property key is the user name. For Local user authentication
	LocalAccounts map[string]ConfigSwitchLocalAccountsUser `json:"local_accounts,omitempty"`
	MxedgeProxyHost string `json:"mxedge_proxy_host,omitempty"`
	MxedgeProxyPort int32 `json:"mxedge_proxy_port,omitempty"`
	ProtectRe *AllOfswitchMgmtProtectRe `json:"protect_re,omitempty"`
	Radius *AllOfswitchMgmtRadius `json:"radius,omitempty"`
	RootPassword string `json:"root_password,omitempty"`
	Tacacs *Tacacs `json:"tacacs,omitempty"`
	// to use mxedge as proxy
	UseMxedgeProxy bool `json:"use_mxedge_proxy,omitempty"`
}
