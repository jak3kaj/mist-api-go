# CaptureScan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** | filter by ap_mac | [optional] [default to null]
**Aps** | [**map[string]CaptureScanAps**](capture_scan_aps.md) | dictionary key is AP mac and value is a dictionary which contains key “band”, “bandwidth”, “channel” and “tcpdump_expression”. In case keys are missed we will take parent value if parent values are not set we will use default value | [optional] [default to null]
**Band** | [***AllOfcaptureScanBand**](AllOfcaptureScanBand.md) |  | [optional] [default to null]
**Bandwidth** | [***AllOfcaptureScanBandwidth**](AllOfcaptureScanBandwidth.md) |  | [optional] [default to null]
**Channel** | **int32** | specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values | [optional] [default to 1]
**ClientMac** | **string** | filter by client mac | [optional] [default to null]
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | [***AllOfcaptureScanFormat**](AllOfcaptureScanFormat.md) |  | [optional] [default to null]
**MaxPktLen** | **int32** | max_len of each packet to capture | [optional] [default to 512]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**TcpdumpExpression** | **string** | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. | [optional] [default to null]
**Type_** | **string** | enum: &#x60;scan&#x60; | [default to null]
**Width** | **string** | specify the bandwidth value with respect to the channel. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

