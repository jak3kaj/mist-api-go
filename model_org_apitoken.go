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

// Org API Token  **Note:** `privilege` field is required to create the object, but may not be  returned in the POST API Response (only in the afterward GET)
type OrgApitoken struct {
	// email of the token creator / null if creator is deleted
	CreatedBy string `json:"created_by,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	Id string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
	LastUsed float64 `json:"last_used,omitempty"`
	// name of the token
	Name string `json:"name"`
	OrgId string `json:"org_id,omitempty"`
	// list of privileges the token has on the orgs/sites
	Privileges []PrivilegeOrg `json:"privileges,omitempty"`
	// list of allowed IP addresses from where the token can be used from. At most 10 IP addresses can be specified, cannot be changed once the API Token is created.
	SrcIps []string `json:"src_ips,omitempty"`
}
