# CaptureGateway

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | [***AllOfcaptureGatewayFormat**](AllOfcaptureGatewayFormat.md) |  | [optional] [default to null]
**Gateways** | [**map[string]CaptureGatewayGateways**](capture_gateway_gateways.md) | List of SSRs. Property key is the SSR MAC | [optional] [default to null]
**MaxPktLen** | **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ports** | **[]string** | dict of port which uses port id as the key | [optional] [default to null]
**Type_** | **string** | enum: &#x60;gateway&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

