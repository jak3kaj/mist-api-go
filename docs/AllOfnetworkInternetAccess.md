# AllOfnetworkInternetAccess

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateSimpleServicePolicy** | **bool** |  | [optional] [default to false]
**DestinationNat** | [**map[string]NetworkDestinationNatProperty**](network_destination_nat_property.md) | Property key may be an IP/Port (i.e. \&quot;63.16.0.3:443\&quot;), or a port (i.e. \&quot;:2222\&quot;) | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Restricted** | **bool** | by default, all access is allowed, to only allow certain traffic, make &#x60;restricted&#x60;&#x3D;&#x60;true&#x60; and define service_policies | [optional] [default to false]
**StaticNat** | [**map[string]NetworkStaticNatProperty**](network_static_nat_property.md) | Property key may be an IP Address (i.e. \&quot;172.16.0.1\&quot;), and IP Address and Port (i.e. \&quot;172.16.0.1:8443\&quot;) or a CIDR (i.e. \&quot;172.16.0.12/20\&quot;) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

