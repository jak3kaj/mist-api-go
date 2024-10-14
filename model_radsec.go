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

// Radsec settings
type Radsec struct {
	CoaEnabled bool `json:"coa_enabled,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	IdleTimeout int32 `json:"idle_timeout,omitempty"`
	// To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org mxedge(s) identified by mxcluster_ids
	MxclusterIds []string `json:"mxcluster_ids,omitempty"`
	// default is site.mxedge.radsec.proxy_hosts which must be a superset of all wlans[*].radsec.proxy_hosts when radsec.proxy_hosts are not used, tunnel peers (org or site mxedges) are used irrespective of use_site_mxedge
	ProxyHosts []string `json:"proxy_hosts,omitempty"`
	// name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge.
	ServerName string `json:"server_name,omitempty"`
	// List of Radsec Servers. Only if not Mist Edge.
	Servers []RadsecServer `json:"servers,omitempty"`
	// use mxedge(s) as radsecproxy
	UseMxedge bool `json:"use_mxedge,omitempty"`
	// To use Site mxedges when this WLAN does not use mxtunnel
	UseSiteMxedge bool `json:"use_site_mxedge,omitempty"`
}
