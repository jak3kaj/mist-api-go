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

type OrgIdTicketsBody struct {
	CaseNumber string `json:"case_number,omitempty"`
	Comments []TicketComment `json:"comments,omitempty"`
	CreatedAt int32 `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	Requester string `json:"requester,omitempty"`
	// email of the requester
	RequesterEmail string `json:"requester_email,omitempty"`
	Status *any `json:"status,omitempty"`
	Subject string `json:"subject"`
	// question (default) / bug / critical
	Type_ string `json:"type"`
	UpdatedAt int32 `json:"updated_at,omitempty"`
}
