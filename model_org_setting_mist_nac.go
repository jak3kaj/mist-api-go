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

type OrgSettingMistNac struct {
	// list of PEM-encoded ca certs
	Cacerts []string `json:"cacerts,omitempty"`
	// use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm.
	DefaultIdpId string `json:"default_idp_id,omitempty"`
	// to disable RSAE_PSS_SHA256, RSAE_PSS_SHA384, RSAE_PSS_SHA512 from server side. see https://www.openssl.org/docs/man3.0/man1/openssl-ciphers.html
	DisableRsaeAlgorithms bool `json:"disable_rsae_algorithms,omitempty"`
	// eap ssl security level see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR
	EapSslSecurityLevel int32 `json:"eap_ssl_security_level,omitempty"`
	// By default NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliancy NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled
	EuOnly bool `json:"eu_only,omitempty"`
	IdpMachineCertLookupField *AllOforgSettingMistNacIdpMachineCertLookupField `json:"idp_machine_cert_lookup_field,omitempty"`
	IdpUserCertLookupField *AllOforgSettingMistNacIdpUserCertLookupField `json:"idp_user_cert_lookup_field,omitempty"`
	Idps []OrgSettingMistNacIdp `json:"idps,omitempty"`
	ServerCert *AllOforgSettingMistNacServerCert `json:"server_cert,omitempty"`
	UseIpVersion *AllOforgSettingMistNacUseIpVersion `json:"use_ip_version,omitempty"`
	// By default NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(radsec) to reach mist-nac.  Set `use_ssl_port`==`true` to override that port with TCP43 (ssl),  This is a org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled
	UseSslPort bool `json:"use_ssl_port,omitempty"`
}
