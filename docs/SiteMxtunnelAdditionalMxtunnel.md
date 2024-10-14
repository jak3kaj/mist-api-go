# SiteMxtunnelAdditionalMxtunnel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | [**[]SiteMxtunnelCluster**](site_mxtunnel_cluster.md) | for AP, how to connect to tunterm or radsecproxy | [optional] [default to null]
**HelloInterval** | **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries | [optional] [default to 60]
**HelloRetries** | **int32** |  | [optional] [default to 7]
**Protocol** | [***AllOfsiteMxtunnelAdditionalMxtunnelProtocol**](AllOfsiteMxtunnelAdditionalMxtunnelProtocol.md) |  | [optional] [default to null]
**VlanIds** | **[]int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

