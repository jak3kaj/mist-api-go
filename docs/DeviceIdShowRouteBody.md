# DeviceIdShowRouteBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value. | [optional] [default to 0]
**Interval** | **int32** | rate at which output will refresh | [optional] [default to 0]
**Neighbor** | **string** | IP of the neighbor | [optional] [default to null]
**Node** | [***HaClusterNode**](ha_cluster_node.md) |  | [optional] [default to null]
**Prefix** | **string** | route prefix | [optional] [default to null]
**Protocol** | [***Object**](.md) |  | [optional] [default to null]
**Route** | **string** | if specified, dump bot received and advertised, if not specified, both will be shown   * for SSR, show bgp neighbors 10.250.18.202 received-routes/advertised-routes   * for SRX and Switches, show route receive_protocol/advertise_protocol bgp 192.168.255.12&#x27; | [optional] [default to null]
**Vrf** | **string** | VRF name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

