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

// GroupWidgetDefinition The groups widget allows you to keep similar graphs together on your timeboard. Each group has a custom header, can hold one to many graphs, and is collapsible
type GroupWidgetDefinition struct {
	LayoutType WidgetLayoutType `json:"layout_type"`
	// Title of the widget
	Title *string `json:"title,omitempty"`
	// Type of the widget
	Type    string   `json:"type"`
	Widgets []Widget `json:"widgets"`
}

// NewGroupWidgetDefinition instantiates a new GroupWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupWidgetDefinition(layoutType WidgetLayoutType, type_ string, widgets []Widget) *GroupWidgetDefinition {
	this := GroupWidgetDefinition{}
	this.LayoutType = layoutType
	this.Type = type_
	this.Widgets = widgets
	return &this
}

// NewGroupWidgetDefinitionWithDefaults instantiates a new GroupWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWidgetDefinitionWithDefaults() *GroupWidgetDefinition {
	this := GroupWidgetDefinition{}
	var type_ string = "group"
	this.Type = type_
	return &this
}

// GetLayoutType returns the LayoutType field value
func (o *GroupWidgetDefinition) GetLayoutType() WidgetLayoutType {
	if o == nil {
		var ret WidgetLayoutType
		return ret
	}

	return o.LayoutType
}

// SetLayoutType sets field value
func (o *GroupWidgetDefinition) SetLayoutType(v WidgetLayoutType) {
	o.LayoutType = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GroupWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *GroupWidgetDefinition) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GroupWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value
func (o *GroupWidgetDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *GroupWidgetDefinition) SetType(v string) {
	o.Type = v
}

// GetWidgets returns the Widgets field value
func (o *GroupWidgetDefinition) GetWidgets() []Widget {
	if o == nil {
		var ret []Widget
		return ret
	}

	return o.Widgets
}

// SetWidgets sets field value
func (o *GroupWidgetDefinition) SetWidgets(v []Widget) {
	o.Widgets = v
}

func (o GroupWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["layout_type"] = o.LayoutType
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["widgets"] = o.Widgets
	}
	return json.Marshal(toSerialize)
}

// AsWidgetDefinition wraps this instance of GroupWidgetDefinition in WidgetDefinition
func (s *GroupWidgetDefinition) AsWidgetDefinition() WidgetDefinition {
	return WidgetDefinition{WidgetDefinitionInterface: s}
}

type NullableGroupWidgetDefinition struct {
	value *GroupWidgetDefinition
	isSet bool
}

func (v NullableGroupWidgetDefinition) Get() *GroupWidgetDefinition {
	return v.value
}

func (v NullableGroupWidgetDefinition) Set(val *GroupWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v NullableGroupWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupWidgetDefinition(val *GroupWidgetDefinition) *NullableGroupWidgetDefinition {
	return &NullableGroupWidgetDefinition{value: val, isSet: true}
}

func (v NullableGroupWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}