# AllOfstatsApRadioStatBand24

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | [***Object**](.md) |  | [optional] [default to null]
**Channel** | **int32** | current channel the radio is running on | [optional] [default to null]
**DynamicChainingEnalbed** | **bool** | Use dynamic chaining for downlink | [optional] [default to null]
**Mac** | **string** | radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af) | [optional] [default to null]
**NoiseFloor** | **int32** |  | [optional] [default to null]
**NumClients** | **int32** |  | [optional] [default to null]
**NumWlans** | **int32** | how many WLANs are applied to the radio | [optional] [default to null]
**Power** | **int32** | transmit power (in dBm) | [optional] [default to null]
**RxBytes** | **int32** |  | [optional] [default to null]
**RxPkts** | **int32** |  | [optional] [default to null]
**TxBytes** | **int32** |  | [optional] [default to null]
**TxPkts** | **int32** |  | [optional] [default to null]
**Usage** | **string** |  | [optional] [default to null]
**UtilAll** | **int32** | all utilization in percentage | [optional] [default to null]
**UtilNonWifi** | **int32** | reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise | [optional] [default to null]
**UtilRxInBss** | **int32** | reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS | [optional] [default to null]
**UtilRxOtherBss** | **int32** | reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS | [optional] [default to null]
**UtilTx** | **int32** | transmission utilization in percentage | [optional] [default to null]
**UtilUndecodableWifi** | **int32** | reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio | [optional] [default to null]
**UtilUnknownWifi** | **int32** | reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

