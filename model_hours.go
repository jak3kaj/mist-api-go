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

// hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).   **Note**: If the dow is not defined then it\\u2019\\ s treated as 00:00-23:59.
type Hours struct {
	Fri string `json:"fri,omitempty"`
	Mon string `json:"mon,omitempty"`
	Sat string `json:"sat,omitempty"`
	Sun string `json:"sun,omitempty"`
	Thu string `json:"thu,omitempty"`
	Tue string `json:"tue,omitempty"`
	Wed string `json:"wed,omitempty"`
}
