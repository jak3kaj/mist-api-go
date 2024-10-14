# CaptureSwitch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | [***AllOfcaptureSwitchFormat**](AllOfcaptureSwitchFormat.md) |  | [optional] [default to null]
**MaxPktLen** | **int32** | max_len of each packet to capture | [optional] [default to 512]
**NumPackets** | **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ports** | [**map[string]CaptureSwitchPortsTcpdumpExpression**](capture_switch_ports_tcpdump_expression.md) | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request | [optional] [default to null]
**Switches** | [**map[string]CaptureSwitchSwitches**](capture_switch_switches.md) | Property key is the switch mac | [optional] [default to null]
**TcpdumpExpression** | **string** | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. | [optional] [default to null]
**Type_** | **string** | enum: &#x60;switch&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

