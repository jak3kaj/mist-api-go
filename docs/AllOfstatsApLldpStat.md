# AllOfstatsApLldpStat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | **string** |  | [optional] [default to null]
**LldpMedSupported** | **bool** | whether it support LLDP-MED | [optional] [default to null]
**MgmtAddr** | **string** | switch’s management address (if advertised), can be IPv4, IPv6, or MAC | [optional] [default to null]
**MgmtAddrs** | **[]string** |  | [optional] [default to null]
**PortDesc** | **string** | ge-0/0/4 | [optional] [default to null]
**PortId** | **string** |  | [optional] [default to null]
**PowerAllocated** | **float64** | in mW, provided/allocated by PSE | [optional] [default to null]
**PowerDraw** | **float64** | in mW, total power needed by PD | [optional] [default to null]
**PowerRequestCount** | **int32** | number of negotiations, if it keeps increasing, we don’ t have a stable power | [optional] [default to null]
**PowerRequested** | **float64** | in mW, the current power requested by PD | [optional] [default to null]
**SystemDesc** | **string** | description provided by switch | [optional] [default to null]
**SystemName** | **string** | name of the switch | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

