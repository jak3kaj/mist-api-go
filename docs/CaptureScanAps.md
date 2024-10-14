# CaptureScanAps

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | [***AllOfcaptureScanApsBand**](AllOfcaptureScanApsBand.md) |  | [optional] [default to null]
**Channel** | **string** | specify the channel value where scan PCAP has to be started | [optional] [default to null]
**TcpdumpExpression** | **string** | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. | [optional] [default to null]
**Width** | **string** | specify the bandwidth value with respect to the channel. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

