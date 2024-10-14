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

type OrgIdPsksBody struct {
	// sso id for psk created from psk portal
	AdminSsoId string `json:"admin_sso_id,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	// email to send psk expiring notifications to
	Email string `json:"email,omitempty"`
	// Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)
	ExpireTime int32 `json:"expire_time,omitempty"`
	// Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire
	ExpiryNotificationTime int32 `json:"expiry_notification_time,omitempty"`
	Id string `json:"id,omitempty"`
	// if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`
	Mac string `json:"mac,omitempty"`
	// if `usage`==`macs`, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000
	Macs []string `json:"macs,omitempty"`
	// For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited)
	MaxUsage int32 `json:"max_usage,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	Note string `json:"note,omitempty"`
	// If set to true, reminder notification will be sent when psk is about to expire
	NotifyExpiry bool `json:"notify_expiry,omitempty"`
	// If set to true, notification will be sent when psk is created or edited
	NotifyOnCreateOrEdit bool `json:"notify_on_create_or_edit,omitempty"`
	// previous passphrase of the PSK if it has been rotated
	OldPassphrase string `json:"old_passphrase,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// passphrase of the PSK (8-63 character or 64 in hex)
	Passphrase string `json:"passphrase"`
	Role string `json:"role,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// SSID this PSK should be applicable to
	Ssid string `json:"ssid"`
	Usage *Object `json:"usage,omitempty"`
	VlanId *Object `json:"vlan_id,omitempty"`
}
