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

type InlineResponse20012 struct {
	AdminId string `json:"admin_id,omitempty"`
	ComplianceStatus *any `json:"compliance_status,omitempty"`
	// if admin account is not an Org API Token
	Email string `json:"email,omitempty"`
	// if admin account is not an Org API Token
	EnableTwoFactor bool `json:"enable_two_factor,omitempty"`
	ExpireTime int32 `json:"expire_time,omitempty"`
	// if admin account is not an Org API Token for an invite, this is the original first name used
	FirstName string `json:"first_name,omitempty"`
	// if admin account is not an Org API Token, how long the invite should be valid
	Hours int32 `json:"hours,omitempty"`
	// if admin account is not an Org API Token for an invite, this is the original last name used
	LastName string `json:"last_name,omitempty"`
	// for Org API Token Only
	Name string `json:"name,omitempty"`
	// when it doesn’t exist, it’s assumed true on EU (i.e. no tracking, the user has to opt-in); otherwise, the user would have to opt-out
	NoTracking bool `json:"no_tracking,omitempty"`
	// if admin account is not an Org API Token
	OauthGoogle bool `json:"oauth_google,omitempty"`
	// password last modified time, in epoch
	PasswordModifiedTime float64 `json:"password_modified_time,omitempty"`
	// if admin account is not an Org API Token phone number (numbers only, including country code)
	Phone string `json:"phone,omitempty"`
	// if admin account is not an Org API Token secondary phone number (numbers only, including country code)
	Phone2 string `json:"phone2,omitempty"`
	// list of privileges the admin has
	Privileges []AdminPrivilege `json:"privileges,omitempty"`
	SessionExpiry int64 `json:"session_expiry,omitempty"`
	Tags []string `json:"tags,omitempty"`
	// if admin account is not an Org API Token two factor status
	TwoFactorVerified bool `json:"two_factor_verified,omitempty"`
	// if admin account is not an Org API Token an admin login via_sso is more restircted. (password and email cannot be changed)
	ViaSso bool `json:"via_sso,omitempty"`
}
