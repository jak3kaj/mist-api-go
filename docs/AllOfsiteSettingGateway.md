# AllOfsiteSettingGateway

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] [default to null]
**BgpConfig** | [**map[string]BgpConfig**](bgp_config.md) |  | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**DhcpdConfig** | [***map[string]DhcpdConfigProperty**](map.md) |  | [optional] [default to null]
**DnsOverride** | **bool** |  | [optional] [default to false]
**DnsServers** | **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] [default to null]
**DnsSuffix** | **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] [default to null]
**ExtraRoutes** | [**map[string]GatewayExtraRoute**](gateway_extra_route.md) | Property key is the destination CIDR (e.g. \&quot;10.0.0.0/8\&quot;) | [optional] [default to null]
**ExtraRoutes6** | [**map[string]GatewayExtraRoute**](gateway_extra_route.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] [default to null]
**GatewayMatching** | [***Object**](.md) |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**IdpProfiles** | [**map[string]IdpProfile**](idp_profile.md) | Property key is the profile name | [optional] [default to null]
**IpConfigs** | [**map[string]GatewayIpConfigProperty**](gateway_ip_config_property.md) | Property key is the network name | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Networks** | [**[]Network**](network.md) |  | [optional] [default to null]
**NtpOverride** | **bool** |  | [optional] [default to false]
**NtpServers** | **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] [default to null]
**OobIpConfig** | [***Object**](.md) |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**PathPreferences** | [**map[string]GatewayPathPreferences**](gateway_path_preferences.md) | Property key is the path name | [optional] [default to null]
**PortConfig** | [**map[string]GatewayPortConfig**](gateway_port_config.md) | Property key is the port(s) name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] [default to null]
**RouterId** | **string** | auto assigned if not set | [optional] [default to null]
**RoutingPolicies** | [**map[string]RoutingPolicy**](routing_policy.md) | Property key is the routing policy name | [optional] [default to null]
**ServicePolicies** | [**[]ServicePolicy**](service_policy.md) |  | [optional] [default to null]
**TunnelConfigs** | [**map[string]TunnelConfigs**](tunnel_configs.md) | Property key is the tunnel name | [optional] [default to null]
**TunnelProviderOptions** | [***TunnelProviderOptions**](tunnel_provider_options.md) |  | [optional] [default to null]
**Type_** | [***Object**](.md) |  | [optional] [default to null]
**VrfConfig** | [***VrfConfig**](vrf_config.md) |  | [optional] [default to null]
**VrfInstances** | [**map[string]GatewayVrfInstance**](gateway_vrf_instance.md) | Property key is the network name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

