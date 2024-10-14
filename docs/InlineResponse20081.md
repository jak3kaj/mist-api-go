# InlineResponse20081

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclPolicies** | [**[]AclPolicy**](acl_policy.md) |  | [optional] [default to null]
**AclTags** | [**map[string]AclTag**](acl_tag.md) | ACL Tags to identify traffic source or destination. Key name is the tag name | [optional] [default to null]
**AdditionalConfigCmds** | **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**DhcpSnooping** | [***DhcpSnooping**](dhcp_snooping.md) |  | [optional] [default to null]
**DnsServers** | **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] [default to null]
**DnsSuffix** | **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] [default to null]
**ExtraRoutes** | [**map[string]ExtraRoute**](extra_route.md) |  | [optional] [default to null]
**ExtraRoutes6** | [**map[string]ExtraRoute6**](extra_route6.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**ImportOrgNetworks** | **[]string** | Org Networks that we&#x27;d like to import | [optional] [default to null]
**MistNac** | [***Object**](.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Networks** | [**map[string]SwitchNetwork**](switch_network.md) | Property key is network name | [optional] [default to null]
**NtpServers** | **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**OspfAreas** | [**map[string]OspfArea**](ospf_area.md) | Junos OSPF areas | [optional] [default to null]
**PortMirroring** | [**map[string]SwitchPortMirroringProperty**](switch_port_mirroring_property.md) | Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 port mirrorings is allowed | [optional] [default to null]
**PortUsages** | [**map[string]SwitchPortUsage**](switch_port_usage.md) |  | [optional] [default to null]
**RadiusConfig** | [***Object**](.md) |  | [optional] [default to null]
**RemoteSyslog** | [***RemoteSyslog**](remote_syslog.md) |  | [optional] [default to null]
**RemoveExistingConfigs** | **bool** | by default, when we configure a device, we only clean up config we generates. Remove existing configs if enabled | [optional] [default to false]
**SnmpConfig** | [***SnmpConfig**](snmp_config.md) |  | [optional] [default to null]
**SwitchMatching** | [***Object**](.md) |  | [optional] [default to null]
**SwitchMgmt** | [***Object**](.md) |  | [optional] [default to null]
**VrfConfig** | [***VrfConfig**](vrf_config.md) |  | [optional] [default to null]
**VrfInstances** | [**map[string]SwitchVrfInstance**](switch_vrf_instance.md) | Property key is the network name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

