# AllOfdeviceprofileApBleConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconEnabled** | **bool** | whether Mist beacons is enabled | [optional] [default to false]
**BeaconRate** | **int32** | required if &#x60;beacon_rate_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, 1-10, in number-beacons-per-second | [optional] [default to 0]
**BeaconRateMode** | [***Object**](.md) |  | [optional] [default to null]
**BeamDisabled** | **[]int32** | list of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam) | [optional] [default to null]
**CustomBlePacketEnabled** | **bool** | can be enabled if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to send custom packet | [optional] [default to false]
**CustomBlePacketFrame** | **string** | The custom frame to be sent out in this beacon. The frame must be a hexstring | [optional] [default to null]
**CustomBlePacketFreqMsec** | **int32** | Frequency (msec) of data emitted by custom ble beacon | [optional] [default to 0]
**EddystoneUidAdvPower** | **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] [default to 0]
**EddystoneUidBeams** | **string** |  | [optional] [default to null]
**EddystoneUidEnabled** | **bool** | only if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;false&#x60;, Whether Eddystone-UID beacon is enabled | [optional] [default to false]
**EddystoneUidFreqMsec** | **int32** | Frequency (msec) of data emmit by Eddystone-UID beacon | [optional] [default to 0]
**EddystoneUidInstance** | **string** | Eddystone-UID instance for the device | [optional] [default to null]
**EddystoneUidNamespace** | **string** | Eddystone-UID namespace | [optional] [default to null]
**EddystoneUrlAdvPower** | **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] [default to 0]
**EddystoneUrlBeams** | **string** |  | [optional] [default to null]
**EddystoneUrlEnabled** | **bool** | only if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;false&#x60;, Whether Eddystone-URL beacon is enabled | [optional] [default to false]
**EddystoneUrlFreqMsec** | **int32** | Frequency (msec) of data emit by Eddystone-UID beacon | [optional] [default to 0]
**EddystoneUrlUrl** | **string** | URL pointed by Eddystone-URL beacon | [optional] [default to null]
**IbeaconAdvPower** | **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] [default to 0]
**IbeaconBeams** | **string** |  | [optional] [default to null]
**IbeaconEnabled** | **bool** | can be enabled if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to send iBeacon | [optional] [default to false]
**IbeaconFreqMsec** | **int32** | Frequency (msec) of data emmit for iBeacon | [optional] [default to 0]
**IbeaconMajor** | **int32** | Major number for iBeacon | [optional] [default to null]
**IbeaconMinor** | **int32** | Minor number for iBeacon | [optional] [default to null]
**IbeaconUuid** | **string** | optional, if not specified, the same UUID as the beacon will be used | [optional] [default to null]
**Power** | **int32** | required if &#x60;power_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] [default to 9]
**PowerMode** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

