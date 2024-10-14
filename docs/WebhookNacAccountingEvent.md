# WebhookNacAccountingEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | mac address of the AP the client roamed or disconnected from | [optional] [default to null]
**AuthType** | **string** | radius authentication type | [optional] [default to null]
**Bssid** | **string** | it’s the MAC physical address of the access point | [optional] [default to null]
**ClientIp** | **string** | IP Address of client | [optional] [default to null]
**ClientType** | **string** | client type E.g. “wired”, “wireless”, “vty” | [optional] [default to null]
**Mac** | **string** | the client’s mac | [optional] [default to null]
**NasVendor** | **string** | NAS Device vendor name E.g. “Juniper”, “Cisco” | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**RxPkts** | **int32** | number of packets received | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**Ssid** | **string** | ESSID | [optional] [default to null]
**Timestamp** | **float64** | sampling time (in epoch) | [optional] [default to null]
**TxPkts** | **int32** | number of packets sent | [optional] [default to null]
**Type_** | **string** | type of event. E.g. “ACCOUNTING_START”, “ACCOUNTING_UPDATE”, “ACCOUNTING_STOP” | [optional] [default to null]
**Username** | **string** | username authenticated with | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

