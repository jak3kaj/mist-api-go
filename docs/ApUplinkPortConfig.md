# ApUplinkPortConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dot1x** | **bool** | Whether to do 802.1x against uplink switch. When enaled, AP cert will be used to do EAP-TLS and the Org&#x27;s CA Cert has to be provisioned at the switch | [optional] [default to false]
**KeepWlansUpIfDown** | **bool** | by default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

