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

// TimeseriesWidgetDefinition The timeseries visualization allows you to display the evolution of one or more metrics, log events, or Analyzed Spans over time.
type TimeseriesWidgetDefinition struct {
	// List of widget events.
	Events *[]WidgetEvent `json:"events,omitempty"`
	// Available legend sizes for a widget. Should be one of \"0\", \"2\", \"4\", \"8\", \"16\", or \"auto\".
	LegendSize *string `json:"legend_size,omitempty"`
	// List of markers.
	Markers *[]WidgetMarker `json:"markers,omitempty"`
	// List of timeseries widget requests.
	Requests []TimeseriesWidgetRequest `json:"requests"`
	// (screenboard only) Show the legend for this widget.
	ShowLegend *bool       `json:"show_legend,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title      *string          `json:"title,omitempty"`
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string                        `json:"title_size,omitempty"`
	Type      TimeseriesWidgetDefinitionType `json:"type"`
	Yaxis     *WidgetAxis                    `json:"yaxis,omitempty"`
}

// NewTimeseriesWidgetDefinition instantiates a new TimeseriesWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesWidgetDefinition(requests []TimeseriesWidgetRequest, type_ TimeseriesWidgetDefinitionType) *TimeseriesWidgetDefinition {
	this := TimeseriesWidgetDefinition{}
	this.Requests = requests
	this.Type = type_
	return &this
}

// NewTimeseriesWidgetDefinitionWithDefaults instantiates a new TimeseriesWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesWidgetDefinitionWithDefaults() *TimeseriesWidgetDefinition {
	this := TimeseriesWidgetDefinition{}
	var type_ TimeseriesWidgetDefinitionType = "timeseries"
	this.Type = type_
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetEvents() []WidgetEvent {
	if o == nil || o.Events == nil {
		var ret []WidgetEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetEventsOk() (*[]WidgetEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []WidgetEvent and assigns it to the Events field.
func (o *TimeseriesWidgetDefinition) SetEvents(v []WidgetEvent) {
	o.Events = &v
}

// GetLegendSize returns the LegendSize field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetLegendSize() string {
	if o == nil || o.LegendSize == nil {
		var ret string
		return ret
	}
	return *o.LegendSize
}

// GetLegendSizeOk returns a tuple with the LegendSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetLegendSizeOk() (*string, bool) {
	if o == nil || o.LegendSize == nil {
		return nil, false
	}
	return o.LegendSize, true
}

// HasLegendSize returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasLegendSize() bool {
	if o != nil && o.LegendSize != nil {
		return true
	}

	return false
}

// SetLegendSize gets a reference to the given string and assigns it to the LegendSize field.
func (o *TimeseriesWidgetDefinition) SetLegendSize(v string) {
	o.LegendSize = &v
}

// GetMarkers returns the Markers field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetMarkers() []WidgetMarker {
	if o == nil || o.Markers == nil {
		var ret []WidgetMarker
		return ret
	}
	return *o.Markers
}

// GetMarkersOk returns a tuple with the Markers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetMarkersOk() (*[]WidgetMarker, bool) {
	if o == nil || o.Markers == nil {
		return nil, false
	}
	return o.Markers, true
}

// HasMarkers returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasMarkers() bool {
	if o != nil && o.Markers != nil {
		return true
	}

	return false
}

// SetMarkers gets a reference to the given []WidgetMarker and assigns it to the Markers field.
func (o *TimeseriesWidgetDefinition) SetMarkers(v []WidgetMarker) {
	o.Markers = &v
}

// GetRequests returns the Requests field value
func (o *TimeseriesWidgetDefinition) GetRequests() []TimeseriesWidgetRequest {
	if o == nil {
		var ret []TimeseriesWidgetRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetRequestsOk() (*[]TimeseriesWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *TimeseriesWidgetDefinition) SetRequests(v []TimeseriesWidgetRequest) {
	o.Requests = v
}

// GetShowLegend returns the ShowLegend field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetShowLegend() bool {
	if o == nil || o.ShowLegend == nil {
		var ret bool
		return ret
	}
	return *o.ShowLegend
}

// GetShowLegendOk returns a tuple with the ShowLegend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetShowLegendOk() (*bool, bool) {
	if o == nil || o.ShowLegend == nil {
		return nil, false
	}
	return o.ShowLegend, true
}

// HasShowLegend returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasShowLegend() bool {
	if o != nil && o.ShowLegend != nil {
		return true
	}

	return false
}

// SetShowLegend gets a reference to the given bool and assigns it to the ShowLegend field.
func (o *TimeseriesWidgetDefinition) SetShowLegend(v bool) {
	o.ShowLegend = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *TimeseriesWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TimeseriesWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasTitleAlign() bool {
	if o != nil && o.TitleAlign != nil {
		return true
	}

	return false
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *TimeseriesWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasTitleSize() bool {
	if o != nil && o.TitleSize != nil {
		return true
	}

	return false
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *TimeseriesWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value
func (o *TimeseriesWidgetDefinition) GetType() TimeseriesWidgetDefinitionType {
	if o == nil {
		var ret TimeseriesWidgetDefinitionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetTypeOk() (*TimeseriesWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TimeseriesWidgetDefinition) SetType(v TimeseriesWidgetDefinitionType) {
	o.Type = v
}

// GetYaxis returns the Yaxis field value if set, zero value otherwise.
func (o *TimeseriesWidgetDefinition) GetYaxis() WidgetAxis {
	if o == nil || o.Yaxis == nil {
		var ret WidgetAxis
		return ret
	}
	return *o.Yaxis
}

// GetYaxisOk returns a tuple with the Yaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetDefinition) GetYaxisOk() (*WidgetAxis, bool) {
	if o == nil || o.Yaxis == nil {
		return nil, false
	}
	return o.Yaxis, true
}

// HasYaxis returns a boolean if a field has been set.
func (o *TimeseriesWidgetDefinition) HasYaxis() bool {
	if o != nil && o.Yaxis != nil {
		return true
	}

	return false
}

// SetYaxis gets a reference to the given WidgetAxis and assigns it to the Yaxis field.
func (o *TimeseriesWidgetDefinition) SetYaxis(v WidgetAxis) {
	o.Yaxis = &v
}

func (o TimeseriesWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.LegendSize != nil {
		toSerialize["legend_size"] = o.LegendSize
	}
	if o.Markers != nil {
		toSerialize["markers"] = o.Markers
	}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if o.ShowLegend != nil {
		toSerialize["show_legend"] = o.ShowLegend
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Yaxis != nil {
		toSerialize["yaxis"] = o.Yaxis
	}
	return json.Marshal(toSerialize)
}

type NullableTimeseriesWidgetDefinition struct {
	value *TimeseriesWidgetDefinition
	isSet bool
}

func (v NullableTimeseriesWidgetDefinition) Get() *TimeseriesWidgetDefinition {
	return v.value
}

func (v *NullableTimeseriesWidgetDefinition) Set(val *TimeseriesWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesWidgetDefinition(val *TimeseriesWidgetDefinition) *NullableTimeseriesWidgetDefinition {
	return &NullableTimeseriesWidgetDefinition{value: val, isSet: true}
}

func (v NullableTimeseriesWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
