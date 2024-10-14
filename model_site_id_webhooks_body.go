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

type SiteIdWebhooksBody struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	// whether webhook is enabled
	Enabled bool `json:"enabled,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// if `type`=`http-post`, additional custom HTTP headers to add the headers name and value must be string, total bytes of headers name and value must be less than 1000
	Headers map[string]string `json:"headers,omitempty"`
	Id string `json:"id,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// name of the webhook
	Name string `json:"name,omitempty"`
	// required when `oauth2_grant_type`==`client_credentials`
	Oauth2ClientId string `json:"oauth2_client_id,omitempty"`
	// required when `oauth2_grant_type`==`client_credentials`
	Oauth2ClientSecret string `json:"oauth2_client_secret,omitempty"`
	Oauth2GrantType *any `json:"oauth2_grant_type,omitempty"`
	// required when `oauth2_grant_type`==`password`
	Oauth2Password string `json:"oauth2_password,omitempty"`
	// required when `type`==`oauth2`, if provided, will be used in the token request
	Oauth2Scopes []string `json:"oauth2_scopes,omitempty"`
	// required when `type`==`oauth2`
	Oauth2TokenUrl string `json:"oauth2_token_url,omitempty"`
	// required when `oauth2_grant_type`==`password`
	Oauth2Username string `json:"oauth2_username,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// only if `type`=`http-post`   when `secret` is provided, two  HTTP headers will be added:    * X-Mist-Signature-v2: HMAC_SHA256(secret, body)   * X-Mist-Signature: HMAC_SHA1(secret, body)
	Secret string `json:"secret,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// required if `type`=`splunk` If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.'
	SplunkToken string `json:"splunk_token,omitempty"`
	// N.B. For org webhooks, only alarms/audits/client-info/client-join/client-sessions/device-events/device-updowns/mxedge-events/nac-sessions/nac_events topics are supported.
	Topics []WebhookTopic `json:"topics,omitempty"`
	Type_ *any `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	// when url uses HTTPS, whether to verify the certificate
	VerifyCert bool `json:"verify_cert,omitempty"`
}
