# InlineResponse20084

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**DisallowMistServices** | **bool** | whether to disallow Mist Devices in the network | [optional] [default to false]
**Gateway** | **string** |  | [optional] [default to null]
**Gateway6** | **string** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**InternalAccess** | [***NetworkInternalAccess**](network_internal_access.md) |  | [optional] [default to null]
**InternetAccess** | [***Object**](.md) |  | [optional] [default to null]
**Isolation** | **bool** | whether to allow clients in the network to talk to each other | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**RoutedForNetworks** | **[]string** | for a Network (usually LAN), it can be routable to other networks (e.g. OSPF) | [optional] [default to null]
**Subnet** | **string** |  | [optional] [default to null]
**Subnet6** | **string** |  | [optional] [default to null]
**Tenants** | [**map[string]NetworkTenant**](network_tenant.md) |  | [optional] [default to null]
**VlanId** | [***Object**](.md) |  | [optional] [default to null]
**VpnAccess** | [**map[string]NetworkVpnAccessConfig**](network_vpn_access_config.md) | Property key is the VPN name. Whether this network can be accessed from vpn | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

