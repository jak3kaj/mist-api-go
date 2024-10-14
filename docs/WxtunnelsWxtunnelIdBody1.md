# WxtunnelsWxtunnelIdBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**Dmvpn** | [***Object**](.md) |  | [optional] [default to null]
**ForMgmt** | **bool** | determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan | [optional] [default to false]
**ForSite** | **bool** |  | [optional] [default to null]
**HelloInterval** | **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries. | [optional] [default to 60]
**HelloRetries** | **int32** |  | [optional] [default to 7]
**Hostname** | **string** | optional, overwrite the hostname in SCCRQ control message, default is \\u201C\\u201D or null, %H and %M can be used, which will be replace with corresponding values:   * %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC   * %M: MAC (e.g. 5c5b350e0060) | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Ipsec** | [***Object**](.md) |  | [optional] [default to null]
**IsStatic** | **bool** | whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled. | [optional] [default to false]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Mtu** | **int32** | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU | [optional] [default to 0]
**Name** | **string** | The name of the tunnel | [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Peers** | **[]string** | list of remote peers’ IP or hostname | [optional] [default to null]
**RouterId** | **string** | optional, overwrite the router-id in SCCRQ control message, default is “0” or null, can also be an IPv4 address | [optional] [default to null]
**Secret** | **string** | secret, ‘’ if no auth is used | [optional] [default to null]
**Sessions** | [**[]WxlanTunnelSession**](wxlan_tunnel_session.md) | sessions to be established with the tunnel. Has to be &gt;&#x3D; 1 in order for this tunnel to be useful. For management tunnel, it can only have 1 | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**UdpPort** | **int32** | udp port if &#x60;use_udp&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] [default to null]
**UseUdp** | **bool** | whether to use UDP instead of IP (proto&#x3D;115, which is default of L2TPv3) | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

