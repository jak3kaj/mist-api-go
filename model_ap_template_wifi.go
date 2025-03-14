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

type ApTemplateWifi struct {
	CiscoEnabled bool `json:"cisco_enabled,omitempty"`
	Disable11k bool `json:"disable_11k,omitempty"`
	DisableRadiosWhenPowerConstrained bool `json:"disable_radios_when_power_constrained,omitempty"`
	EnableArpSpoof bool `json:"enable_arp_spoof,omitempty"`
	EnableSharedRadioScanning bool `json:"enable_shared_radio_scanning,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	LocateConnected bool `json:"locate_connected,omitempty"`
	LocateUnconnected bool `json:"locate_unconnected,omitempty"`
	MeshAllowDfs bool `json:"mesh_allow_dfs,omitempty"`
	MeshEnableCrm bool `json:"mesh_enable_crm,omitempty"`
	MeshEnabled bool `json:"mesh_enabled,omitempty"`
	ProxyArp bool `json:"proxy_arp,omitempty"`
}
