# CaptureNewAssoc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** |  | [optional] [default to null]
**ClientMac** | **string** | client mac, required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60;; optional otherwise | [optional] [default to null]
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**IncludesMcast** | **bool** |  | [optional] [default to false]
**MaxPktLen** | **int32** |  | [optional] [default to 128]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 100]
**Ssid** | **string** | optional filter by ssid | [optional] [default to null]
**Type_** | **string** | enum: &#x60;new_assoc&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

