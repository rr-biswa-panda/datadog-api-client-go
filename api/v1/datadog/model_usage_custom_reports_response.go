/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// UsageCustomReportsResponse Returns available daily custom reports.
type UsageCustomReportsResponse struct {
	// An array of available daily custom reports.
	Data *[]UsageCustomReportsData `json:"data,omitempty"`
	Meta *UsageCustomReportsMeta   `json:"meta,omitempty"`
}

// NewUsageCustomReportsResponse instantiates a new UsageCustomReportsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageCustomReportsResponse() *UsageCustomReportsResponse {
	this := UsageCustomReportsResponse{}
	return &this
}

// NewUsageCustomReportsResponseWithDefaults instantiates a new UsageCustomReportsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageCustomReportsResponseWithDefaults() *UsageCustomReportsResponse {
	this := UsageCustomReportsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsageCustomReportsResponse) GetData() []UsageCustomReportsData {
	if o == nil || o.Data == nil {
		var ret []UsageCustomReportsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCustomReportsResponse) GetDataOk() (*[]UsageCustomReportsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsageCustomReportsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []UsageCustomReportsData and assigns it to the Data field.
func (o *UsageCustomReportsResponse) SetData(v []UsageCustomReportsData) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsageCustomReportsResponse) GetMeta() UsageCustomReportsMeta {
	if o == nil || o.Meta == nil {
		var ret UsageCustomReportsMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCustomReportsResponse) GetMetaOk() (*UsageCustomReportsMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsageCustomReportsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given UsageCustomReportsMeta and assigns it to the Meta field.
func (o *UsageCustomReportsResponse) SetMeta(v UsageCustomReportsMeta) {
	o.Meta = &v
}

func (o UsageCustomReportsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableUsageCustomReportsResponse struct {
	value *UsageCustomReportsResponse
	isSet bool
}

func (v NullableUsageCustomReportsResponse) Get() *UsageCustomReportsResponse {
	return v.value
}

func (v *NullableUsageCustomReportsResponse) Set(val *UsageCustomReportsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageCustomReportsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageCustomReportsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageCustomReportsResponse(val *UsageCustomReportsResponse) *NullableUsageCustomReportsResponse {
	return &NullableUsageCustomReportsResponse{value: val, isSet: true}
}

func (v NullableUsageCustomReportsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageCustomReportsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}