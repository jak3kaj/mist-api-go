# MxclustersMxclusterIdBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**MistDas** | [***Object**](.md) |  | [optional] [default to null]
**MistNac** | [***MxclusterNac**](mxcluster_nac.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**MxedgeMgmt** | [***MxedgeMgmt**](mxedge_mgmt.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Proxy** | [***Object**](.md) |  | [optional] [default to null]
**Radsec** | [***Object**](.md) |  | [optional] [default to null]
**RadsecTls** | [***MxclusterRadsecTls**](mxcluster_radsec_tls.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**TuntermApSubnets** | **[]string** | list of subnets where we allow AP to establish Mist Tunnels from | [optional] [default to null]
**TuntermDhcpdConfig** | [***Object**](.md) |  | [optional] [default to null]
**TuntermExtraRoutes** | [**map[string]MxclusterTuntermExtraRoute**](mxcluster_tunterm_extra_route.md) | extra routes for Mist Tunneled VLANs. Property key is a CIDR | [optional] [default to null]
**TuntermHosts** | **[]string** | hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) | [optional] [default to null]
**TuntermHostsOrder** | **[]int32** | list of index of tunterm_hosts | [optional] [default to null]
**TuntermHostsSelection** | [***Object**](.md) |  | [optional] [default to null]
**TuntermMonitoring** | [**[][]TuntermMonitoringItem**](array.md) |  | [optional] [default to null]
**TuntermMonitoringDisabled** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

