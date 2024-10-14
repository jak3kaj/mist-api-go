# RftemplateRadioBand24

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRrmDisable** | **bool** |  | [optional] [default to false]
**AntGain** | **int32** |  | [optional] [default to 0]
**AntennaMode** | [***AllOfrftemplateRadioBand24AntennaMode**](AllOfrftemplateRadioBand24AntennaMode.md) |  | [optional] [default to null]
**Bandwidth** | [***AllOfrftemplateRadioBand24Bandwidth**](AllOfrftemplateRadioBand24Bandwidth.md) |  | [optional] [default to null]
**Channels** | **[]int32** | For RFTemplates. List of channels, null or empty array means auto | [optional] [default to []]
**Disabled** | **bool** | whether to disable the radio | [optional] [default to false]
**Power** | **int32** | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / … | [optional] [default to null]
**PowerMax** | **int32** | when power&#x3D;0, max tx power to use, HW-specific values will be used if not set | [optional] [default to 17]
**PowerMin** | **int32** | when power&#x3D;0, min tx power to use, HW-specific values will be used if not set | [optional] [default to 8]
**Preamble** | [***AllOfrftemplateRadioBand24Preamble**](AllOfrftemplateRadioBand24Preamble.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

