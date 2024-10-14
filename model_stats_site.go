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

// Site statistics
type StatsSite struct {
	Address string `json:"address"`
	AlarmtemplateId string `json:"alarmtemplate_id"`
	CountryCode string `json:"country_code"`
	CreatedTime float64 `json:"created_time"`
	Id string `json:"id"`
	Lat float64 `json:"lat"`
	Latlng *LatLng `json:"latlng"`
	Lng float64 `json:"lng"`
	ModifiedTime float64 `json:"modified_time"`
	MspId string `json:"msp_id"`
	Name string `json:"name"`
	NetworktemplateId string `json:"networktemplate_id"`
	NumAp int32 `json:"num_ap"`
	NumApConnected int32 `json:"num_ap_connected"`
	NumClients int32 `json:"num_clients"`
	NumDevices int32 `json:"num_devices"`
	NumDevicesConnected int32 `json:"num_devices_connected"`
	NumGateway int32 `json:"num_gateway"`
	NumGatewayConnected int32 `json:"num_gateway_connected"`
	NumSwitch int32 `json:"num_switch"`
	NumSwitchConnected int32 `json:"num_switch_connected"`
	OrgId string `json:"org_id"`
	RftemplateId string `json:"rftemplate_id"`
	SecpolicyId *Object `json:"secpolicy_id,omitempty"`
	SitegroupIds []string `json:"sitegroup_ids"`
	Timezone string `json:"timezone"`
	Tzoffset int32 `json:"tzoffset"`
}
