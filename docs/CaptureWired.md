# CaptureWired

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** |  | [optional] [default to null]
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | [***AllOfcaptureWiredFormat**](AllOfcaptureWiredFormat.md) |  | [optional] [default to null]
**MaxPktLen** | **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**TcpdumpExpression** | **string** | tcpdump expression | [optional] [default to null]
**Type_** | **string** | enum: &#x60;wired&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

