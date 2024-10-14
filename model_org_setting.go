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

// Org Settings
type OrgSetting struct {
	// enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.
	ApUpdownThreshold int32 `json:"ap_updown_threshold,omitempty"`
	ApiPolicy *OrgSettingApiPolicy `json:"api_policy,omitempty"`
	AutoDeviceNaming *OrgSettingAutoDeviceNaming `json:"auto_device_naming,omitempty"`
	AutoDeviceprofileAssignment *OrgSettingAutoDeviceprofileAssignment `json:"auto_deviceprofile_assignment,omitempty"`
	AutoSiteAssignment *OrgSettingAutoSiteAssignment `json:"auto_site_assignment,omitempty"`
	BlacklistUrl string `json:"blacklist_url,omitempty"`
	// list of PEM-encoded ca certs
	Cacerts []string `json:"cacerts,omitempty"`
	Celona *OrgSettingCelona `json:"celona,omitempty"`
	Cloudshark *OrgSettingCloudshark `json:"cloudshark,omitempty"`
	Cradlepoint *OrgSettingCradlepoint `json:"cradlepoint,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	DeviceCert *AllOforgSettingDeviceCert `json:"device_cert,omitempty"`
	// enable threshold-based device down delivery via   * device-updowns webhooks topic,    * Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate)
	DeviceUpdownThreshold int32 `json:"device_updown_threshold,omitempty"`
	// whether to disallow Mist to analyze pcap files (this is required for marvis pcap)
	DisablePcap bool `json:"disable_pcap,omitempty"`
	// whether to disable remote shell access for an entire org
	DisableRemoteShell bool `json:"disable_remote_shell,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	GatewayMgmt *OrgSettingGatewayMgmt `json:"gateway_mgmt,omitempty"`
	// enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.
	GatewayUpdownThreshold int32 `json:"gateway_updown_threshold,omitempty"`
	Id string `json:"id,omitempty"`
	Installer *OrgSettingInstaller `json:"installer,omitempty"`
	Jcloud *OrgSettingJcloud `json:"jcloud,omitempty"`
	Juniper *AccountJuniperInfo `json:"juniper,omitempty"`
	Mgmt *AllOforgSettingMgmt `json:"mgmt,omitempty"`
	MistNac *OrgSettingMistNac `json:"mist_nac,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	MspId string `json:"msp_id,omitempty"`
	MxedgeFipsEnabled bool `json:"mxedge_fips_enabled,omitempty"`
	MxedgeMgmt *MxedgeMgmt `json:"mxedge_mgmt,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	PasswordPolicy *AllOforgSettingPasswordPolicy `json:"password_policy,omitempty"`
	Pcap *OrgSettingPcap `json:"pcap,omitempty"`
	PcapBucketVerified bool `json:"pcap_bucket_verified,omitempty"`
	PortChannelization *map[string]string `json:"port_channelization,omitempty"`
	Security *OrgSettingSecurity `json:"security,omitempty"`
	SimpleAlert *AllOforgSettingSimpleAlert `json:"simple_alert,omitempty"`
	SwitchMgmt *OrgSettingSwitchMgmt `json:"switch_mgmt,omitempty"`
	// enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.
	SwitchUpdownThreshold int32 `json:"switch_updown_threshold,omitempty"`
	SyntheticTest *SynthetictestConfig `json:"synthetic_test,omitempty"`
	// list of tags
	Tags []string `json:"tags,omitempty"`
	// automatically logout the user when UI session is inactive. `0` means disabled
	UiIdleTimeout int32 `json:"ui_idle_timeout,omitempty"`
	VpnOptions *OrgSettingVpnOptions `json:"vpn_options,omitempty"`
	WanPma *OrgSettingWanPma `json:"wan_pma,omitempty"`
	WiredPma *OrgSettingWiredPma `json:"wired_pma,omitempty"`
	WirelessPma *OrgSettingWirelessPma `json:"wireless_pma,omitempty"`
}
