# RftemplatesRftemplateIdBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntGain24** | **int32** |  | [optional] [default to null]
**AntGain5** | **int32** |  | [optional] [default to null]
**AntGain6** | **int32** |  | [optional] [default to null]
**Band24** | [***Object**](.md) |  | [optional] [default to null]
**Band24Usage** | [***Object**](.md) |  | [optional] [default to null]
**Band5** | [***Object**](.md) |  | [optional] [default to null]
**Band5On24Radio** | [***Object**](.md) |  | [optional] [default to null]
**Band6** | [***Object**](.md) |  | [optional] [default to null]
**CountryCode** | **string** | optional, country code to use. If specified, this gets applied to all sites using the RF Template | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**ModelSpecific** | [**map[string]RfTemplateModelSpecificProperty**](rf_template_model_specific_property.md) | overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. \&quot;AP63\&quot;) | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** | The name of the RF template | [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**ScanningEnabled** | **bool** | whether scanning radio is enabled | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

