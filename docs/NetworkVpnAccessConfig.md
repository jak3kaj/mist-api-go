# NetworkVpnAccessConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisedSubnet** | **string** | if &#x60;routed&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to advertise an aggregated subnet toward HUB this is useful when there are multiple networks on SPOKE&#x27;s side | [optional] [default to null]
**AllowPing** | **bool** | whether to allow ping from vpn into this routed network | [optional] [default to null]
**DestinationNat** | [**map[string]NetworkDestinationNatProperty**](network_destination_nat_property.md) | Property key may be an IP/Port (i.e. \&quot;63.16.0.3:443\&quot;), or a port (i.e. \&quot;:2222\&quot;) | [optional] [default to null]
**NatPool** | **string** | if &#x60;routed&#x60;&#x3D;&#x3D;&#x60;false&#x60; (usually at Spoke), but some hosts needs to be reachable from Hub, a subnet is required to create and advertise the route to Hub | [optional] [default to null]
**NoReadvertiseToLanBgp** | **bool** | toward LAN-side BGP peers | [optional] [default to false]
**NoReadvertiseToLanOspf** | **bool** | toward LAN-side OSPF peers | [optional] [default to false]
**NoReadvertiseToOverlay** | **bool** | toward overlay how HUB should deal with routes it received from Spokes | [optional] [default to null]
**OtherVrfs** | **[]string** | by default, the routes are only readvertised toward the same vrf on spoke to allow it to be leaked to other vrfs | [optional] [default to null]
**Routed** | **bool** | whether this network is routable | [optional] [default to null]
**SourceNat** | [***AllOfnetworkVpnAccessConfigSourceNat**](AllOfnetworkVpnAccessConfigSourceNat.md) |  | [optional] [default to null]
**StaticNat** | [**map[string]NetworkStaticNatProperty**](network_static_nat_property.md) | Property key may be an IP Address (i.e. \&quot;172.16.0.1\&quot;), and IP Address and Port (i.e. \&quot;172.16.0.1:8443\&quot;) or a CIDR (i.e. \&quot;172.16.0.12/20\&quot;) | [optional] [default to null]
**SummarizedSubnet** | **string** | toward overlay how HUB should deal with routes it received from Spokes | [optional] [default to null]
**SummarizedSubnetToLanBgp** | **string** | toward LAN-side BGP peers | [optional] [default to null]
**SummarizedSubnetToLanOspf** | **string** | toward LAN-side OSPF peers | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

