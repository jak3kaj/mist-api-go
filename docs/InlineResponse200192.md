# InlineResponse200192

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** |  | [optional] [default to null]
**Aps** | **[]string** | List of target APs to capture packets | [optional] [default to null]
**ClientMac** | **string** |  | [optional] [default to null]
**Duration** | **int32** |  | [optional] [default to null]
**Failed** | **[]string** | List of APs where configuration attempt failed | [optional] [default to null]
**Format** | [***Object**](.md) |  | [optional] [default to null]
**Gateways** | **[]string** | Information on gateways to capture packets on if a gateway capture type is specified | [optional] [default to null]
**Id** | **string** | unique id for the capture | [default to null]
**IncludesMcast** | **bool** |  | [optional] [default to null]
**MaxNumPackets** | **int32** | max number of packets configured by user | [optional] [default to null]
**MaxPktLen** | **int32** |  | [optional] [default to null]
**NumPackets** | **int32** | total number of packets captured by all AP, not applicable for type [client, new_assoc] | [optional] [default to null]
**Ok** | **[]string** | List of target APs successfully configured to capture packets | [optional] [default to null]
**PcapAps** | [**map[string]ResponsePcapAp**](response_pcap_ap.md) |  | [optional] [default to null]
**RadiotapTcpdumpExpression** | **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;radiotap&#x60;, radiotap_tcpdump_expression expression provided by the user | [optional] [default to null]
**ScanTcpdumpExpression** | **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;scan&#x60;, scan_tcpdump_expression provided by the user | [optional] [default to null]
**Ssid** | **string** |  | [optional] [default to null]
**StartedTime** | **int32** |  | [optional] [default to null]
**Switches** | **[]string** | Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured. | [optional] [default to null]
**TcpdumpExpression** | **string** | tcpdump expression provided by the user (common) | [optional] [default to null]
**Type_** | [***Object**](.md) |  | [default to null]
**TzspHost** | **string** | Required if &#x60;format&#x60;&#x3D;&#x3D;&#x60;tzsp&#x60;. Remote host accessible to mxedges over the network for receiving the captured packets. | [optional] [default to null]
**TzspPort** | **int32** | if &#x60;format&#x60;&#x3D;&#x3D;&#x60;tzsp&#x60;. Port on remote host for receiving the captured packets | [optional] [default to null]
**WiredTcpdumpExpression** | **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;wired&#x60;, wired_tcpdump_expression provided by the user | [optional] [default to null]
**WirelessTcpdumpExpression** | **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;‘wireless’&#x60;, wireless_tcpdump_expression provided by the user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

