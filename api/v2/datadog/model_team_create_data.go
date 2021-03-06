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

// TeamCreateData Team data for a create request.
type TeamCreateData struct {
	Attributes    *TeamCreateAttributes `json:"attributes,omitempty"`
	Relationships *TeamRelationships    `json:"relationships,omitempty"`
	Type          TeamType              `json:"type"`
}

// NewTeamCreateData instantiates a new TeamCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamCreateData(type_ TeamType) *TeamCreateData {
	this := TeamCreateData{}
	this.Type = type_
	return &this
}

// NewTeamCreateDataWithDefaults instantiates a new TeamCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamCreateDataWithDefaults() *TeamCreateData {
	this := TeamCreateData{}
	var type_ TeamType = "teams"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TeamCreateData) GetAttributes() TeamCreateAttributes {
	if o == nil || o.Attributes == nil {
		var ret TeamCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamCreateData) GetAttributesOk() (*TeamCreateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TeamCreateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TeamCreateAttributes and assigns it to the Attributes field.
func (o *TeamCreateData) SetAttributes(v TeamCreateAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TeamCreateData) GetRelationships() TeamRelationships {
	if o == nil || o.Relationships == nil {
		var ret TeamRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamCreateData) GetRelationshipsOk() (*TeamRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TeamCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TeamRelationships and assigns it to the Relationships field.
func (o *TeamCreateData) SetRelationships(v TeamRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value
func (o *TeamCreateData) GetType() TeamType {
	if o == nil {
		var ret TeamType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamCreateData) GetTypeOk() (*TeamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TeamCreateData) SetType(v TeamType) {
	o.Type = v
}

func (o TeamCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTeamCreateData struct {
	value *TeamCreateData
	isSet bool
}

func (v NullableTeamCreateData) Get() *TeamCreateData {
	return v.value
}

func (v *NullableTeamCreateData) Set(val *TeamCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamCreateData(val *TeamCreateData) *NullableTeamCreateData {
	return &NullableTeamCreateData{value: val, isSet: true}
}

func (v NullableTeamCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
