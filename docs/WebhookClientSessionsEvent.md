# WebhookClientSessionsEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | mac address of the AP the client roamed or disconnected from | [default to null]
**ApName** | **string** | user-friendly name of the AP the client roamed or disconnected from. | [default to null]
**Band** | **string** | 5GHz or 2.4GHz band | [default to null]
**Bssid** | **string** |  | [default to null]
**ClientFamily** | **string** | device family E.g. “Mac”, “iPhone”, “Apple watch” | [default to null]
**ClientManufacture** | **string** | device manufacturer E.g. “Apple” | [default to null]
**ClientModel** | **string** | device model E.g. “8+”, “XS” | [default to null]
**ClientOs** | **string** | device operating system E.g. “Mojave”, “Windows 10”, “Linux” | [default to null]
**Connect** | **int32** | time when the user connects | [default to null]
**ConnectFloat** | **float64** | floating point connect timestamp with millisecond precision | [default to null]
**Disconnect** | **int32** | time when the user disconnects | [default to null]
**DisconnectFloat** | **float64** | floating point disconnect timestamp with millisecond precision | [default to null]
**Duration** | **int32** | the duration of the roamed or complete session indicated by termination_reason field. | [default to null]
**Mac** | **string** | the client’s mac | [default to null]
**NextAp** | **string** | the AP the client has roamed to. | [default to null]
**OrgId** | **string** |  | [default to null]
**Rssi** | **float64** | latest average RSSI before the user disconnects | [default to null]
**SiteId** | **string** |  | [default to null]
**SiteName** | **string** |  | [default to null]
**Ssid** | **string** |  | [default to null]
**TerminationReason** | **int32** | 1 disassociate - when the client disassociates. 2 inactive - when the client is timeout. 3 roamed - when the client is roamed between APs | [default to null]
**Timestamp** | **float64** |  | [default to null]
**Version** | **float64** | schema version of this message | [default to null]
**WlanId** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

