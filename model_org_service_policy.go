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

type OrgServicePolicy struct {
	Action *AllOforgServicePolicyAction `json:"action,omitempty"`
	Antivirus *AllOforgServicePolicyAntivirus `json:"antivirus,omitempty"`
	Appqoe *AllOforgServicePolicyAppqoe `json:"appqoe,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	Ewf []ServicePolicyEwfRule `json:"ewf,omitempty"`
	Id string `json:"id,omitempty"`
	Idp *IdpConfig `json:"idp,omitempty"`
	// access within the same VRF
	LocalRouting bool `json:"local_routing,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// by default, we derive all paths available and use them optionally, you can customize by using `path_preference`
	PathPreference string `json:"path_preference,omitempty"`
	Secintel *AllOforgServicePolicySecintel `json:"secintel,omitempty"`
	Services []string `json:"services,omitempty"`
	SslProxy *AllOforgServicePolicySslProxy `json:"ssl_proxy,omitempty"`
	Tenants []string `json:"tenants,omitempty"`
}
