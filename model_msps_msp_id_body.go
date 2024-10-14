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

type MspsMspIdBody struct {
	AllowMist bool `json:"allow_mist,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	Id string `json:"id,omitempty"`
	// For advanced tier (uMSPs) only
	LogoUrl string `json:"logo_url,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name,omitempty"`
	Tier *any `json:"tier,omitempty"`
	// For advanced tier (uMSPs) only
	Url string `json:"url,omitempty"`
}
