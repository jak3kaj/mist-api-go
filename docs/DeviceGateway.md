# DeviceGateway

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] [default to null]
**BgpConfig** | [**map[string]BgpConfig**](bgp_config.md) |  | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**DeviceprofileId** | **string** |  | [optional] [default to null]
**DhcpdConfig** | [***map[string]DhcpdConfigProperty**](map.md) |  | [optional] [default to null]
**DnsServers** | **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] [default to null]
**DnsSuffix** | **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] [default to null]
**ExtraRoutes** | [**map[string]GatewayExtraRoute**](gateway_extra_route.md) | Property key is the destination CIDR (e.g. \&quot;10.0.0.0/8\&quot;) | [optional] [default to null]
**ExtraRoutes6** | [**map[string]GatewayExtraRoute**](gateway_extra_route.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**IdpProfiles** | [**map[string]IdpProfile**](idp_profile.md) | Property key is the profile name | [optional] [default to null]
**Image1Url** | **string** |  | [optional] [default to null]
**Image2Url** | **string** |  | [optional] [default to null]
**Image3Url** | **string** |  | [optional] [default to null]
**IpConfigs** | [**map[string]GatewayIpConfigProperty**](gateway_ip_config_property.md) | Property key is the network name | [optional] [default to null]
**Mac** | **string** | device MAC address | [optional] [default to null]
**Managed** | **bool** |  | [optional] [default to null]
**MapId** | **string** | map where the device belongs to | [optional] [default to null]
**Model** | **string** | device Model | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**MspId** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Networks** | [**[]Network**](network.md) |  | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]
**NtpServers** | **[]string** |  | [optional] [default to null]
**OobIpConfig** | [***AllOfdeviceGatewayOobIpConfig**](AllOfdeviceGatewayOobIpConfig.md) |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**PathPreferences** | [**map[string]GatewayPathPreferences**](gateway_path_preferences.md) | Property key is the path name | [optional] [default to null]
**PortConfig** | [**map[string]GatewayPortConfig**](gateway_port_config.md) | Property key is the port name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] [default to null]
**PortMirroring** | [***GatewayPortMirroring**](gateway_port_mirroring.md) |  | [optional] [default to null]
**RouterId** | **string** | auto assigned if not set | [optional] [default to null]
**RoutingPolicies** | [**map[string]RoutingPolicy**](routing_policy.md) | Property key is the routing policy name | [optional] [default to null]
**Serial** | **string** | device Serial | [optional] [default to null]
**ServicePolicies** | [**[]ServicePolicy**](service_policy.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**TunnelConfigs** | [**map[string]TunnelConfigs**](tunnel_configs.md) | Property key is the tunnel name | [optional] [default to null]
**TunnelProviderOptions** | [***TunnelProviderOptions**](tunnel_provider_options.md) |  | [optional] [default to null]
**Type_** | **string** | Device Type. enum: &#x60;gateway&#x60; | [default to null]
**Vars** | **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] [default to null]
**VrfConfig** | [***VrfConfig**](vrf_config.md) |  | [optional] [default to null]
**VrfInstances** | [**map[string]GatewayVrfInstance**](gateway_vrf_instance.md) | Property key is the network name | [optional] [default to null]
**X** | **float64** | x in pixel | [optional] [default to null]
**Y** | **float64** | y in pixel | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

