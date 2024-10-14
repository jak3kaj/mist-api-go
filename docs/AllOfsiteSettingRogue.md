# AllOfsiteSettingRogue

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | whether or not rogue detection is enabled | [optional] [default to false]
**HoneypotEnabled** | **bool** | whether or not honeypot detection is enabled | [optional] [default to false]
**MinDuration** | **int32** | minimum duration for a bssid to be considered rogue | [optional] [default to 10]
**MinRssi** | **int32** | minimum RSSI for an AP to be considered rogue (ignoring APs that’s far away) | [optional] [default to -80]
**WhitelistedBssids** | **[]string** | list of BSSIDs to whitelist. Ex: \&quot;cc-:8e-:6f-:d4-:bf-:16\&quot;, \&quot;cc-8e-6f-d4-bf-16\&quot;, \&quot;cc-73-*\&quot;, \&quot;cc:82:*\&quot; | [optional] [default to null]
**WhitelistedSsids** | **[]string** | list of SSIDs to whitelist | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

