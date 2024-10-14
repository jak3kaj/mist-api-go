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

type AllOfalarmTemplateRuleDelivery struct {
	// List of additional email string to deliver the alarms via emails
	AdditionalEmails []string `json:"additional_emails,omitempty"`
	// Whether to enable the alarm delivery via emails or not
	Enabled bool `json:"enabled"`
	// Whether to deliver the alarms via emails to Org admins or not
	ToOrgAdmins bool `json:"to_org_admins,omitempty"`
	// Whether to deliver the alarms via emails to Site admins or not
	ToSiteAdmins bool `json:"to_site_admins,omitempty"`
}
