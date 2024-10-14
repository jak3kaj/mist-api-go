# SwitchPortUsage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllNetworks** | **bool** | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60; whether to trunk all network/vlans | [optional] [default to false]
**AllowDhcpd** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; if DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state.  When it is not defined, it means using the system’s default setting which depends on whether the port is a access or trunk port. | [optional] [default to null]
**AllowMultipleSupplicants** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; | [optional] [default to false]
**BypassAuthWhenServerDown** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; bypass auth for known clients if set to true when RADIUS server is down | [optional] [default to false]
**BypassAuthWhenServerDownForUnkonwnClient** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60; bypass auth for all (including unknown clients) if set to true when RADIUS server is down | [optional] [default to false]
**Description** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; | [optional] [default to null]
**DisableAutoneg** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; if speed and duplex are specified, whether to disable autonegotiation | [optional] [default to false]
**Disabled** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; whether the port is disabled | [optional] [default to false]
**Duplex** | [***AllOfswitchPortUsageDuplex**](AllOfswitchPortUsageDuplex.md) |  | [optional] [default to null]
**DynamicVlanNetworks** | **[]string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60;, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return | [optional] [default to null]
**EnableMacAuth** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; whether to enable MAC Auth | [optional] [default to false]
**EnableQos** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; | [optional] [default to false]
**GuestNetwork** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) | [optional] [default to null]
**InterSwitchLink** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; inter_switch_link is used together with \&quot;isolation\&quot; under networks NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together | [optional] [default to false]
**MacAuthOnly** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;enable_mac_auth&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] [default to null]
**MacAuthPreferred** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; + &#x60;enable_mac_auth&#x60;&#x3D;&#x3D;&#x60;true&#x60; + &#x60;mac_auth_only&#x60;&#x3D;&#x3D;&#x60;false&#x60;, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x. | [optional] [default to null]
**MacAuthProtocol** | [***AllOfswitchPortUsageMacAuthProtocol**](AllOfswitchPortUsageMacAuthProtocol.md) |  | [optional] [default to null]
**MacLimit** | **int32** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform | [optional] [default to 0]
**Mode** | [***AllOfswitchPortUsageMode**](AllOfswitchPortUsageMode.md) |  | [optional] [default to null]
**Mtu** | **int32** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. | [optional] [default to null]
**Networks** | **[]string** | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60;, the list of network/vlans | [optional] [default to null]
**PersistMac** | **bool** | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;access&#x60; and &#x60;port_auth&#x60;!&#x3D;&#x60;dot1x&#x60; whether the port should retain dynamically learned MAC addresses | [optional] [default to false]
**PoeDisabled** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; whether PoE capabilities are disabled for a port | [optional] [default to false]
**PortAuth** | [***AllOfswitchPortUsagePortAuth**](AllOfswitchPortUsagePortAuth.md) |  | [optional] [default to null]
**PortNetwork** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; native network/vlan for untagged traffic | [optional] [default to null]
**ReauthInterval** | **int32** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60; reauthentication interval range | [optional] [default to 3600]
**ResetDefaultWhen** | [***AllOfswitchPortUsageResetDefaultWhen**](AllOfswitchPortUsageResetDefaultWhen.md) |  | [optional] [default to null]
**Rules** | [**[]SwitchPortUsageDynamicRule**](switch_port_usage_dynamic_rule.md) | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;dynamic&#x60; | [optional] [default to null]
**ServerFailNetwork** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; sets server fail fallback vlan | [optional] [default to null]
**ServerRejectNetwork** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; when radius server reject / fails | [optional] [default to null]
**Speed** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; speed, default is auto to automatically negotiate speed | [optional] [default to null]
**StormControl** | [***AllOfswitchPortUsageStormControl**](AllOfswitchPortUsageStormControl.md) |  | [optional] [default to null]
**StpEdge** | **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; when enabled, the port is not expected to receive BPDU frames | [optional] [default to false]
**StpNoRootPort** | **bool** |  | [optional] [default to false]
**StpP2p** | **bool** |  | [optional] [default to false]
**VoipNetwork** | **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

