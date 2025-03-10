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

type LicenseSub struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	// end date of the license term
	EndTime int32 `json:"end_time,omitempty"`
	Id string `json:"id,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	OrderId string `json:"order_id,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// number of devices entitled for this license
	Quantity int32 `json:"quantity,omitempty"`
	// Number of licenses left in this subscription
	RemainingQuantity int32 `json:"remaining_quantity,omitempty"`
	// start date of the license term
	StartTime int32 `json:"start_time,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	Type_ *AllOflicenseSubType_ `json:"type,omitempty"`
}
