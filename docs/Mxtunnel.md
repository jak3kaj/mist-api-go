# Mxtunnel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnchorMxtunnelIds** | **[]string** | list of anchor mxtunnels used for forming edge to edge tunnels | [optional] [default to null]
**AutoPreemption** | [***AllOfmxtunnelAutoPreemption**](AllOfmxtunnelAutoPreemption.md) |  | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**HelloInterval** | **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by &#x60;hello_retries&#x60;. | [optional] [default to 60]
**HelloRetries** | **int32** |  | [optional] [default to 7]
**Id** | **string** |  | [optional] [default to null]
**Ipsec** | [***MxtunnelIpsec**](mxtunnel_ipsec.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Mtu** | **int32** | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU | [optional] [default to 0]
**MxclusterIds** | **[]string** | list of mxclusters to deploy this tunnel to | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Protocol** | [***AllOfmxtunnelProtocol**](AllOfmxtunnelProtocol.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**VlanIds** | **[]int32** | list of vlan_ids that will be used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

