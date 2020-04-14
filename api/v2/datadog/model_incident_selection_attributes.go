/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// IncidentSelectionAttributes The Incident Selection attributes.
type IncidentSelectionAttributes struct {
	// Timestamp of when an selection was created.
	Created *time.Time `json:"created,omitempty"`
	// Timestamp of when an selection was modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Represents the incident ID.
	ObjectId *string `json:"object_id,omitempty"`
}

// NewIncidentSelectionAttributes instantiates a new IncidentSelectionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentSelectionAttributes() *IncidentSelectionAttributes {
	this := IncidentSelectionAttributes{}
	return &this
}

// NewIncidentSelectionAttributesWithDefaults instantiates a new IncidentSelectionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentSelectionAttributesWithDefaults() *IncidentSelectionAttributes {
	this := IncidentSelectionAttributes{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IncidentSelectionAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSelectionAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IncidentSelectionAttributes) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IncidentSelectionAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentSelectionAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSelectionAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentSelectionAttributes) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentSelectionAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *IncidentSelectionAttributes) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSelectionAttributes) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *IncidentSelectionAttributes) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *IncidentSelectionAttributes) SetObjectId(v string) {
	o.ObjectId = &v
}

func (o IncidentSelectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.ObjectId != nil {
		toSerialize["object_id"] = o.ObjectId
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentSelectionAttributes struct {
	value *IncidentSelectionAttributes
	isSet bool
}

func (v NullableIncidentSelectionAttributes) Get() *IncidentSelectionAttributes {
	return v.value
}

func (v *NullableIncidentSelectionAttributes) Set(val *IncidentSelectionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentSelectionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentSelectionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentSelectionAttributes(val *IncidentSelectionAttributes) *NullableIncidentSelectionAttributes {
	return &NullableIncidentSelectionAttributes{value: val, isSet: true}
}

func (v NullableIncidentSelectionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentSelectionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}