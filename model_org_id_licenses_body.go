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

type OrgIdLicensesBody struct {
	// if `op`==`unamend`, the ID of the operation to cancel
	AmendmentId string `json:"amendment_id,omitempty"`
	// if `op`==`amend`, the id of the org where the license is moved
	DstOrgId string `json:"dst_org_id,omitempty"`
	// if `op`==`annotate`
	Notes string `json:"notes,omitempty"`
	Op *any `json:"op"`
	// if `op`==`amend`, the number of licenses to move
	Quantity int32 `json:"quantity,omitempty"`
	// if `op`==`amend` or `op`==`delete`, the ID of the subscription to use
	SubscriptionId string `json:"subscription_id,omitempty"`
}
