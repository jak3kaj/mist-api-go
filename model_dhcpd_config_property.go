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

type DhcpdConfigProperty struct {
	// if `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used
	DnsServers []string `json:"dns_servers,omitempty"`
	// if `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// if `type`==`local` or `type6`==`local`. Property key is the MAC Address. Format is `[0-9a-f]{12}` (e.g \"5684dae9ac8b\")
	FixedBindings map[string]DhcpdConfigFixedBinding `json:"fixed_bindings,omitempty"`
	// if `type`==`local` - optional, `ip` will be used if not provided
	Gateway string `json:"gateway,omitempty"`
	// if `type`==`local`
	IpEnd string `json:"ip_end,omitempty"`
	// if `type6`==`local`
	IpEnd6 string `json:"ip_end6,omitempty"`
	// if `type`==`local`
	IpStart string `json:"ip_start,omitempty"`
	// if `type6`==`local`
	IpStart6 string `json:"ip_start6,omitempty"`
	// in seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]
	LeaseTime int32 `json:"lease_time,omitempty"`
	// if `type`==`local` or `type6`==`local`. Property key is the DHCP option number
	Options map[string]DhcpdConfigOption `json:"options,omitempty"`
	// `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,  should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.
	ServerIdOverride bool `json:"server_id_override,omitempty"`
	// if `type`==`relay`
	Servers []string `json:"servers,omitempty"`
	// if `type6`==`relay`
	Servers6 []string `json:"servers6,omitempty"`
	Type_ *AllOfdhcpdConfigPropertyType_ `json:"type,omitempty"`
	Type6 *AllOfdhcpdConfigPropertyType6 `json:"type6,omitempty"`
	// if `type`==`local` or `type6`==`local`. Property key is <enterprise number>:<sub option code>, with   * enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers)   * sub option code: 1-255, sub-option code'
	VendorEncapulated map[string]DhcpdConfigVendorOption `json:"vendor_encapulated,omitempty"`
}
