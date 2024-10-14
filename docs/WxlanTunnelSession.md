# WxlanTunnelSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApAsSessionId** | **string** | if &#x60;use_ap_as_session_ids&#x60;&#x3D;&#x3D;&#x60;true&#x60;, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id | [optional] [default to null]
**Comment** | **string** | optional, user-specified string for display purpose | [optional] [default to null]
**EnableCookie** | **bool** |  | [optional] [default to null]
**Ethertype** | [***AllOfwxlanTunnelSessionEthertype**](AllOfwxlanTunnelSessionEthertype.md) |  | [optional] [default to null]
**LocalSessionId** | **int32** | 1-2147483647 | [optional] [default to null]
**Pseudo8021adEnabled** | **bool** | optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers. | [optional] [default to false]
**RemoteId** | **string** | remote-id of the session, has to be unique in the same tunnel | [optional] [default to null]
**RemoteSessionId** | **int32** | 1-2147483647 | [optional] [default to null]
**UseApAsSessionIds** | **bool** | whether to use AP (last 4 bytes of MAC currently) as session ids | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

