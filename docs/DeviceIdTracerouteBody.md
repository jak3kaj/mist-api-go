# DeviceIdTracerouteBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | host name | [optional] [default to null]
**Network** | **string** | for SSR, optional, the source to initiate traceroute from | [optional] [default to internal]
**Node** | [***Object**](.md) |  | [optional] [default to null]
**Port** | **int32** | when &#x60;protocol&#x60;&#x3D;&#x3D;&#x60;udp&#x60;, not supported in SSR. The udp port to use | [optional] [default to 33434]
**Protocol** | [***Object**](.md) |  | [optional] [default to null]
**Timeout** | **int32** | not supported in SSR. Maximum time in seconds to wait for the response | [optional] [default to 60]
**Vrf** | **string** | for SRX, optional, the source to initiate traceroute from. by default, master VRF/RI is assumed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

