# Mxedge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Magic** | **string** |  | [optional] [default to null]
**Model** | **string** |  | [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**MxagentRegistered** | **bool** |  | [optional] [default to null]
**MxclusterId** | **string** | MxCluster this MxEdge belongs to | [optional] [default to null]
**MxedgeMgmt** | [***MxedgeMgmt**](mxedge_mgmt.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Note** | **string** |  | [optional] [default to null]
**NtpServers** | **[]string** |  | [optional] [default to null]
**OobIpConfig** | [***AllOfmxedgeOobIpConfig**](AllOfmxedgeOobIpConfig.md) |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Proxy** | [***AllOfmxedgeProxy**](AllOfmxedgeProxy.md) |  | [optional] [default to null]
**Services** | **[]string** | list of services to run, tunterm only for now | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**TuntermDhcpdConfig** | [***AllOfmxedgeTuntermDhcpdConfig**](AllOfmxedgeTuntermDhcpdConfig.md) |  | [optional] [default to null]
**TuntermExtraRoutes** | [**map[string]MxedgeTuntermExtraRoute**](mxedge_tunterm_extra_route.md) | Property key is a CIDR | [optional] [default to null]
**TuntermIgmpSnoopingConfig** | [***MxedgeTuntermIgmpSnoopingConfig**](mxedge_tunterm_igmp_snooping_config.md) |  | [optional] [default to null]
**TuntermIpConfig** | [***AllOfmxedgeTuntermIpConfig**](AllOfmxedgeTuntermIpConfig.md) |  | [optional] [default to null]
**TuntermMonitoring** | [**[][]TuntermMonitoringItem**](array.md) |  | [optional] [default to null]
**TuntermMulticastConfig** | [***MxedgeTuntermMulticastConfig**](mxedge_tunterm_multicast_config.md) |  | [optional] [default to null]
**TuntermOtherIpConfigs** | [**map[string]MxedgeTuntermOtherIpConfig**](mxedge_tunterm_other_ip_config.md) | ip configs by VLAN ID. Property key is the VLAN ID | [optional] [default to null]
**TuntermPortConfig** | [***AllOfmxedgeTuntermPortConfig**](AllOfmxedgeTuntermPortConfig.md) |  | [optional] [default to null]
**TuntermRegistered** | **bool** |  | [optional] [default to null]
**TuntermSwitchConfig** | [***AllOfmxedgeTuntermSwitchConfig**](AllOfmxedgeTuntermSwitchConfig.md) |  | [optional] [default to null]
**Versions** | [***AllOfmxedgeVersions**](AllOfmxedgeVersions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

