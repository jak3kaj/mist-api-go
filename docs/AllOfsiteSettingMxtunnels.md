# AllOfsiteSettingMxtunnels

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMxtunnels** | [**map[string]SiteMxtunnelAdditionalMxtunnel**](site_mxtunnel_additional_mxtunnel.md) |  | [optional] [default to null]
**ApSubnets** | **[]string** | list of subnets where we allow AP to establish Mist Tunnels from | [optional] [default to null]
**AutoPreemption** | [***Object**](.md) |  | [optional] [default to null]
**Clusters** | [**[]SiteMxtunnelCluster**](site_mxtunnel_cluster.md) | for AP, how to connect to tunterm or radsecproxy | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**HelloInterval** | **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries | [optional] [default to 60]
**HelloRetries** | **int32** |  | [optional] [default to 7]
**Hosts** | **[]string** | hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Mtu** | **int32** | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU | [optional] [default to 0]
**OrgId** | **string** |  | [optional] [default to null]
**Protocol** | [***Object**](.md) |  | [optional] [default to null]
**Radsec** | [***SiteMxtunnelRadsec**](site_mxtunnel_radsec.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**VlanIds** | **[]int32** | list of vlan_ids that will be used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

