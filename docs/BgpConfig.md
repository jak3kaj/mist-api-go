# BgpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | **string** |  | [optional] [default to null]
**BfdMinimumInterval** | **int32** | when bfd_multiplier is configured alone. Default:   * 1000 if &#x60;type&#x60;&#x3D;&#x3D;&#x60;external&#x60;   * 350 &#x60;type&#x60;&#x3D;&#x3D;&#x60;internal&#x60; | [optional] [default to 350]
**BfdMultiplier** | **int32** | when bfd_minimum_interval_is_configured alone | [optional] [default to 3]
**Communities** | [**[]BgpConfigCommunity**](bgp_config_community.md) |  | [optional] [default to null]
**DisableBfd** | **bool** | BFD provides faster path failure detection and is enabled by default | [optional] [default to false]
**Export** | **string** |  | [optional] [default to null]
**ExportPolicy** | **string** | default export policies if no per-neighbor policies defined | [optional] [default to null]
**ExtendedV4Nexthop** | **bool** | by default, either inet/net6 unicast depending on neighbor IP family (v4 or v6) for v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this | [optional] [default to null]
**GracefulRestartTime** | **int32** | &#x60;0&#x60; means disable | [optional] [default to 0]
**HoldTime** | **int32** |  | [optional] [default to 90]
**Import_** | **string** |  | [optional] [default to null]
**ImportPolicy** | **string** | default import policies if no per-neighbor policies defined | [optional] [default to null]
**LocalAs** | **int32** |  | [optional] [default to null]
**NeighborAs** | **int32** |  | [optional] [default to null]
**Neighbors** | [**map[string]BgpConfigNeighbors**](bgp_config_neighbors.md) | if per-neighbor as is desired. Property key is the neighbor address | [optional] [default to null]
**Networks** | **[]string** | if &#x60;type&#x60;!&#x3D;&#x60;external&#x60;or &#x60;via&#x60;&#x3D;&#x3D;&#x60;wan&#x60;networks where we expect BGP neighbor to connect to/from | [optional] [default to null]
**NoReadvertiseToOverlay** | **bool** | by default, we&#x27;ll re-advertise all learned BGP routers toward overlay | [optional] [default to false]
**TunnelName** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;tunnel&#x60; | [optional] [default to null]
**Type_** | [***AllOfbgpConfigType_**](AllOfbgpConfigType_.md) |  | [optional] [default to null]
**Via** | [***AllOfbgpConfigVia**](AllOfbgpConfigVia.md) |  | [optional] [default to null]
**VpnName** | **string** |  | [optional] [default to null]
**WanName** | **string** | if &#x60;via&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

