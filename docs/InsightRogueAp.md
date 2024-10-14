# InsightRogueAp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** | mac of the device that had strongest signal strength for ssid/bssid pair | [default to null]
**AvgRssi** | **float64** | average signal strength of ap_mac for ssid/bssid pair | [default to null]
**Bssid** | **string** | bssid of the network detected as threat | [default to null]
**Channel** | **string** | channel over which ap_mac heard ssid/bssid pair | [default to null]
**DeltaX** | **float64** | X position relative to the reporting AP (&#x60;ap_mac&#x60;) | [optional] [default to null]
**DeltaY** | **float64** | Y position relative to the reporting AP (&#x60;ap_mac&#x60;) | [optional] [default to null]
**NumAps** | **int32** | num of aps that heard the ssid/bssid pair | [default to null]
**SeenOnLan** | **bool** | whether the reporting AP see a wireless client (on LAN) connecting to it | [optional] [default to null]
**Ssid** | **string** | ssid of the network detected as threat | [optional] [default to null]
**TimesHeard** | **int32** | represents number of times the pair was heard in the interval. Each count roughly corresponds to a minute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

