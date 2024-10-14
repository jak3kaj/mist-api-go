# CaptureMxedge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | [***AllOfcaptureMxedgeFormat**](AllOfcaptureMxedgeFormat.md) |  | [optional] [default to null]
**MaxPktLen** | **int32** | max_len of each packet to capture | [optional] [default to 128]
**Mxedges** | [**map[string]CaptureMxedgeMxedges**](capture_mxedge_mxedges.md) |  | [optional] [default to null]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Type_** | **string** | enum: &#x60;mxedge&#x60; | [default to null]
**TzspHost** | **string** | Required if &#x60;format&#x60;&#x3D;&#x3D;&#x60;tzsp&#x60;. Remote host accessible to mxedges over the network for receiving the captured packets. | [optional] [default to null]
**TzspPort** | **int32** | if &#x60;format&#x60;&#x3D;&#x3D;&#x60;tzsp&#x60;. Port on remote host for receiving the captured packets | [optional] [default to 37008]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

