# AllOfstatsMxedgeCpuStat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | [**map[string]CpuStat**](cpu_stat.md) |  | [optional] [default to null]
**Idle** | **int32** | percentage of Idle, Idle/(Idle + Busy) since last sampling | [optional] [default to null]
**Interrupt** | **int32** | percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling | [optional] [default to null]
**System** | **int32** | percentage of System, System/(Idle + Busy) since last sampling | [optional] [default to null]
**Usage** | **int32** | percentage of load, Busy/(Idle + Busy) since last sampling | [optional] [default to null]
**User** | **int32** | percentage of User, User/(Idle + Busy) since last sampling | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

