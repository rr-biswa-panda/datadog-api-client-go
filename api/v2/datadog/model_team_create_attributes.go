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

// TeamCreateAttributes The team's attributes for a create request.
type TeamCreateAttributes struct {
	// Name of the team.
	Name string `json:"name"`
}

// NewTeamCreateAttributes instantiates a new TeamCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamCreateAttributes(name string) *TeamCreateAttributes {
	this := TeamCreateAttributes{}
	this.Name = name
	return &this
}

// NewTeamCreateAttributesWithDefaults instantiates a new TeamCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamCreateAttributesWithDefaults() *TeamCreateAttributes {
	this := TeamCreateAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *TeamCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamCreateAttributes) SetName(v string) {
	o.Name = v
}

func (o TeamCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTeamCreateAttributes struct {
	value *TeamCreateAttributes
	isSet bool
}

func (v NullableTeamCreateAttributes) Get() *TeamCreateAttributes {
	return v.value
}

func (v *NullableTeamCreateAttributes) Set(val *TeamCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamCreateAttributes(val *TeamCreateAttributes) *NullableTeamCreateAttributes {
	return &NullableTeamCreateAttributes{value: val, isSet: true}
}

func (v NullableTeamCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
