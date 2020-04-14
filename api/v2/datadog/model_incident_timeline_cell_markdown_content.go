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

// IncidentTimelineCellMarkdownContent A Markdown cell content schema.
type IncidentTimelineCellMarkdownContent struct {
	// The markdown content of the cell.
	Content string `json:"content"`
}

// NewIncidentTimelineCellMarkdownContent instantiates a new IncidentTimelineCellMarkdownContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTimelineCellMarkdownContent(content string) *IncidentTimelineCellMarkdownContent {
	this := IncidentTimelineCellMarkdownContent{}
	this.Content = content
	return &this
}

// NewIncidentTimelineCellMarkdownContentWithDefaults instantiates a new IncidentTimelineCellMarkdownContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTimelineCellMarkdownContentWithDefaults() *IncidentTimelineCellMarkdownContent {
	this := IncidentTimelineCellMarkdownContent{}
	return &this
}

// GetContent returns the Content field value
func (o *IncidentTimelineCellMarkdownContent) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownContent) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *IncidentTimelineCellMarkdownContent) SetContent(v string) {
	o.Content = v
}

func (o IncidentTimelineCellMarkdownContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

// AsIncidentTimelineCellCommonContentOneOf wraps this instance of IncidentTimelineCellMarkdownContent in IncidentTimelineCellCommonContentOneOf
func (s *IncidentTimelineCellMarkdownContent) AsIncidentTimelineCellCommonContentOneOf() IncidentTimelineCellCommonContentOneOf {
	return IncidentTimelineCellCommonContentOneOf{IncidentTimelineCellCommonContentOneOfInterface: s}
}

type NullableIncidentTimelineCellMarkdownContent struct {
	value *IncidentTimelineCellMarkdownContent
	isSet bool
}

func (v NullableIncidentTimelineCellMarkdownContent) Get() *IncidentTimelineCellMarkdownContent {
	return v.value
}

func (v *NullableIncidentTimelineCellMarkdownContent) Set(val *IncidentTimelineCellMarkdownContent) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTimelineCellMarkdownContent) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTimelineCellMarkdownContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTimelineCellMarkdownContent(val *IncidentTimelineCellMarkdownContent) *NullableIncidentTimelineCellMarkdownContent {
	return &NullableIncidentTimelineCellMarkdownContent{value: val, isSet: true}
}

func (v NullableIncidentTimelineCellMarkdownContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTimelineCellMarkdownContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}