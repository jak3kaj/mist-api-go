# RoutingPolicyTermMatching

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsPath** | **[]string** | takes regular expression | [optional] [default to null]
**Community** | **[]string** |  | [optional] [default to null]
**Network** | **[]string** |  | [optional] [default to null]
**Prefix** | **[]string** | zero or more criteria/filter can be specified to match the term, all criteria have to be met | [optional] [default to null]
**Protocol** | **[]string** | &#x60;direct&#x60;, &#x60;bgp&#x60;, &#x60;osp&#x60;, ... | [optional] [default to null]
**RouteExists** | [***RoutingPolicyTermMatchingRouteExists**](routing_policy_term_matching_route_exists.md) |  | [optional] [default to null]
**VpnNeighborMac** | **[]string** | overlay-facing criteria (used for bgp_config where via&#x3D;vpn) | [optional] [default to null]
**VpnPath** | **[]string** | overlay-facing criteria (used for bgp_config where via&#x3D;vpn) ordered- | [optional] [default to null]
**VpnPathSla** | [***RoutingPolicyTermMatchingVpnPathSla**](routing_policy_term_matching_vpn_path_sla.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

