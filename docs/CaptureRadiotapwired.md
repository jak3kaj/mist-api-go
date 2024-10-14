# CaptureRadiotapwired

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** |  | [optional] [default to null]
**Band** | [***AllOfcaptureRadiotapwiredBand**](AllOfcaptureRadiotapwiredBand.md) |  | [optional] [default to null]
**ClientMac** | **string** |  | [optional] [default to null]
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | [***AllOfcaptureRadiotapwiredFormat**](AllOfcaptureRadiotapwiredFormat.md) |  | [optional] [default to null]
**MaxPktLen** | **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**RadiotapTcpdumpExpression** | **string** | tcpdump expression for radiotap interface (802.11 + radio headers) | [optional] [default to null]
**Ssid** | **string** |  | [optional] [default to null]
**TcpdumpExpression** | **string** | tcpdump expression common for wired,radiotap | [optional] [default to null]
**Type_** | **string** | enum: &#x60;radiotap,wired&#x60; | [default to null]
**WiredTcpdumpExpression** | **string** | tcpdump expression for wired | [optional] [default to null]
**WirelessTcpdumpExpression** | **string** | tcpdump expression for radiotap interface (802.11) | [optional] [default to null]
**WlanId** | **string** | wlan id associated with the respective ssid. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

