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

type InlineResponse20025 struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	// if `idp_type`==`saml`, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process)
	CustomLogoutUrl string `json:"custom_logout_url,omitempty"`
	// if `idp_type`==`saml`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched
	DefaultRole string `json:"default_role,omitempty"`
	// random string generated during the SSO creation and used to generate the SAML URLs:   * ACS URL = `/api/v1/saml/{domain}/login` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/login`)   * Single Logout URL = `/api/v1/saml/{domain}/logout` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/logout`)
	Domain string `json:"domain,omitempty"`
	// Required if `ldap_type`==`custom`, LDAP filter that will identify the type of group
	GroupFilter string `json:"group_filter,omitempty"`
	Id string `json:"id,omitempty"`
	// if `idp_type`==`saml`. IDP Cert (used to verify the signed response)
	IdpCert string `json:"idp_cert,omitempty"`
	IdpSignAlgo *any `json:"idp_sign_algo,omitempty"`
	// Required if `idp_type`==`saml`, IDP Single-Sign-On URL
	IdpSsoUrl string `json:"idp_sso_url,omitempty"`
	IdpType *any `json:"idp_type,omitempty"`
	// if `idp_type`==`saml`, ignore any unmatched roles provided in assertion. By default, an assertion is treated as invalid for any unmatched role
	IgnoreUnmatchedRoles bool `json:"ignore_unmatched_roles,omitempty"`
	// if `idp_type`==`saml`. IDP issuer URL
	Issuer string `json:"issuer,omitempty"`
	// Required if `idp_type`==`ldap`, whole domain or a specific organization unit (container) in Search base to specify where users and groups are found in the LDAP tree
	LdapBaseDn string `json:"ldap_base_dn,omitempty"`
	// Required if `idp_type`==`ldap`, the account used to authenticate against the LDAP
	LdapBindDn string `json:"ldap_bind_dn,omitempty"`
	// Required if `idp_type`==`ldap`, the password used to authenticate against the LDAP
	LdapBindPassword string `json:"ldap_bind_password,omitempty"`
	// Required if `idp_type`==`ldap`, list of CA certificates to validate the LDAP certificate
	LdapCacerts []string `json:"ldap_cacerts,omitempty"`
	// if `idp_type`==`ldap`, LDAPS Client certificate
	LdapClientCert string `json:"ldap_client_cert,omitempty"`
	// if `idp_type`==`ldap`, Key for the `ldap_client_cert`
	LdapClientKey string `json:"ldap_client_key,omitempty"`
	// if `ldap_type`==`custom`
	LdapGroupAttr string `json:"ldap_group_attr,omitempty"`
	// if `ldap_type`==`custom`
	LdapGroupDn string `json:"ldap_group_dn,omitempty"`
	// if `idp_type`==`ldap`, whether to recursively resolve LDAP groups
	LdapResolveGroups bool `json:"ldap_resolve_groups,omitempty"`
	// if `idp_type`==`ldap`, list of LDAP/LDAPS server IP Addresses or Hostnames
	LdapServerHosts []string `json:"ldap_server_hosts,omitempty"`
	LdapType *any `json:"ldap_type,omitempty"`
	// Required if `ldap_type`==`custom`, LDAP filter that will identify the type of user
	LdapUserFilter string `json:"ldap_user_filter,omitempty"`
	// Required if `ldap_type`==`custom`,LDAP filter that will identify the type of member
	MemberFilter string `json:"member_filter,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	MspId string `json:"msp_id,omitempty"`
	MxedgeProxy *any `json:"mxedge_proxy,omitempty"`
	// name
	Name string `json:"name"`
	NameidFormat *any `json:"nameid_format,omitempty"`
	// Required if `idp_type`==`oauth`, Client Credentials
	OauthCcClientId string `json:"oauth_cc_client_id,omitempty"`
	// Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form \"-----BEGIN RSA PRIVATE KEY--....\"
	OauthCcClientSecret string `json:"oauth_cc_client_secret,omitempty"`
	// if `idp_type`==`oauth`
	OauthDiscoveryUrl string `json:"oauth_discovery_url,omitempty"`
	// if `idp_type`==`oauth`, ropc = Resource Owner Password Credentials
	OauthRopcClientId string `json:"oauth_ropc_client_id,omitempty"`
	// if `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty
	OauthRopcClientSecret string `json:"oauth_ropc_client_secret,omitempty"`
	// Required if `idp_type`==`oauth`, oauth_tenant_id
	OauthTenantId string `json:"oauth_tenant_id,omitempty"`
	OauthType *any `json:"oauth_type,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// if `idp_type`==`saml`, custom role attribute parsing scheme  Supported Role Parsing Schemes <table><tr><th>Name</th><th>Scheme</th></tr><tr><td>cn</td><td><ul><li>The expected role attribute format in SAML Assertion is “CN=cn,OU=ou1,OU=ou2,…”</li><li>CN (the key) is case insensitive and exactly 1 CN is expected (or the entire entry will be ignored)</li><li>E.g. if role attribute is “CN=cn,OU=ou1,OU=ou2” then parsed role value is “cn”</li></ul></td></tr></table>
	RoleAttrExtraction string `json:"role_attr_extraction,omitempty"`
	// if `idp_type`==`saml`, name of the attribute in SAML Assertion to extract role from
	RoleAttrFrom string `json:"role_attr_from,omitempty"`
	// if `idp_type`==`oauth`, indicates if SCIM provisioning is enabled for the OAuth IDP
	ScimEnabled bool `json:"scim_enabled,omitempty"`
	// if `idp_type`==`oauth`, scim_secret_token (generated by caller, crypto-random) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP
	ScimSecretToken string `json:"scim_secret_token,omitempty"`
	SiteId string `json:"site_id,omitempty"`
}
