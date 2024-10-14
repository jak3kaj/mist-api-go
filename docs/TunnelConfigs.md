# TunnelConfigs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoProvision** | [***TunnelConfigsAutoProvision**](tunnel_configs_auto_provision.md) |  | [optional] [default to null]
**IkeLifetime** | **int32** | Only if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to null]
**IkeMode** | [***AllOftunnelConfigsIkeMode**](AllOftunnelConfigsIkeMode.md) |  | [optional] [default to null]
**IkeProposals** | [**[]GatewayTemplateTunnelIkeProposal**](gateway_template_tunnel_ike_proposal.md) | if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to null]
**IpsecLifetime** | **int32** | if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to null]
**IpsecProposals** | [**[]GatewayTemplateTunnelIpsecProposal**](gateway_template_tunnel_ipsec_proposal.md) | Only if  &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to null]
**LocalId** | **string** | Only if:   * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-ipsec&#x60;   * &#x60;provider&#x60;&#x3D;&#x3D;&#x60;jse-ipsec&#x60;   * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to null]
**Mode** | [***AllOftunnelConfigsMode**](AllOftunnelConfigsMode.md) |  | [optional] [default to null]
**Networks** | **[]string** | networks reachable via this tunnel | [optional] [default to null]
**Primary** | [***GatewayTemplateTunnelNode**](gateway_template_tunnel_node.md) |  | [optional] [default to null]
**Probe** | [***AllOftunnelConfigsProbe**](AllOftunnelConfigsProbe.md) |  | [optional] [default to null]
**Protocol** | [***AllOftunnelConfigsProtocol**](AllOftunnelConfigsProtocol.md) |  | [optional] [default to null]
**Provider** | [***AllOftunnelConfigsProvider**](AllOftunnelConfigsProvider.md) |  | [optional] [default to null]
**Psk** | **string** | Only if:   * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-ipsec&#x60;   * &#x60;provider&#x60;&#x3D;&#x3D;&#x60;jse-ipsec&#x60;   * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] [default to null]
**Secondary** | [***GatewayTemplateTunnelNode**](gateway_template_tunnel_node.md) |  | [optional] [default to null]
**Version** | [***AllOftunnelConfigsVersion**](AllOftunnelConfigsVersion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

