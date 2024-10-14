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

type ResponseLogSearchItem struct {
	// admin id
	AdminId string `json:"admin_id"`
	// name of the admin that performs the action
	AdminName string `json:"admin_name"`
	// field values after the change
	After *interface{} `json:"after,omitempty"`
	// field values prior to the change
	Before *interface{} `json:"before,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Id string `json:"id,omitempty"`
	// log message
	Message string `json:"message"`
	// org id
	OrgId string `json:"org_id"`
	SiteId string `json:"site_id"`
	// start time, in epoch
	Timestamp float64 `json:"timestamp"`
}
