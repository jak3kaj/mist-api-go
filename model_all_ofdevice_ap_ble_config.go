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

type AllOfdeviceApBleConfig struct {
	// whether Mist beacons is enabled
	BeaconEnabled bool `json:"beacon_enabled,omitempty"`
	// required if `beacon_rate_mode`==`custom`, 1-10, in number-beacons-per-second
	BeaconRate int32 `json:"beacon_rate,omitempty"`
	BeaconRateMode *any `json:"beacon_rate_mode,omitempty"`
	// list of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam)
	BeamDisabled []int32 `json:"beam_disabled,omitempty"`
	// can be enabled if `beacon_enabled`==`true`, whether to send custom packet
	CustomBlePacketEnabled bool `json:"custom_ble_packet_enabled,omitempty"`
	// The custom frame to be sent out in this beacon. The frame must be a hexstring
	CustomBlePacketFrame string `json:"custom_ble_packet_frame,omitempty"`
	// Frequency (msec) of data emitted by custom ble beacon
	CustomBlePacketFreqMsec int32 `json:"custom_ble_packet_freq_msec,omitempty"`
	// advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
	EddystoneUidAdvPower int32 `json:"eddystone_uid_adv_power,omitempty"`
	EddystoneUidBeams string `json:"eddystone_uid_beams,omitempty"`
	// only if `beacon_enabled`==`false`, Whether Eddystone-UID beacon is enabled
	EddystoneUidEnabled bool `json:"eddystone_uid_enabled,omitempty"`
	// Frequency (msec) of data emmit by Eddystone-UID beacon
	EddystoneUidFreqMsec int32 `json:"eddystone_uid_freq_msec,omitempty"`
	// Eddystone-UID instance for the device
	EddystoneUidInstance string `json:"eddystone_uid_instance,omitempty"`
	// Eddystone-UID namespace
	EddystoneUidNamespace string `json:"eddystone_uid_namespace,omitempty"`
	// advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
	EddystoneUrlAdvPower int32 `json:"eddystone_url_adv_power,omitempty"`
	EddystoneUrlBeams string `json:"eddystone_url_beams,omitempty"`
	// only if `beacon_enabled`==`false`, Whether Eddystone-URL beacon is enabled
	EddystoneUrlEnabled bool `json:"eddystone_url_enabled,omitempty"`
	// Frequency (msec) of data emit by Eddystone-UID beacon
	EddystoneUrlFreqMsec int32 `json:"eddystone_url_freq_msec,omitempty"`
	// URL pointed by Eddystone-URL beacon
	EddystoneUrlUrl string `json:"eddystone_url_url,omitempty"`
	// advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
	IbeaconAdvPower int32 `json:"ibeacon_adv_power,omitempty"`
	IbeaconBeams string `json:"ibeacon_beams,omitempty"`
	// can be enabled if `beacon_enabled`==`true`, whether to send iBeacon
	IbeaconEnabled bool `json:"ibeacon_enabled,omitempty"`
	// Frequency (msec) of data emmit for iBeacon
	IbeaconFreqMsec int32 `json:"ibeacon_freq_msec,omitempty"`
	// Major number for iBeacon
	IbeaconMajor int32 `json:"ibeacon_major,omitempty"`
	// Minor number for iBeacon
	IbeaconMinor int32 `json:"ibeacon_minor,omitempty"`
	// optional, if not specified, the same UUID as the beacon will be used
	IbeaconUuid string `json:"ibeacon_uuid,omitempty"`
	// required if `power_mode`==`custom`
	Power int32 `json:"power,omitempty"`
	PowerMode *any `json:"power_mode,omitempty"`
}
