/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// ApiKeyListResponse struct for ApiKeyListResponse
type ApiKeyListResponse struct {
	ApiKeys *[]ApiKey `json:"api_keys,omitempty"`
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise.
func (o *ApiKeyListResponse) GetApiKeys() []ApiKey {
	if o == nil || o.ApiKeys == nil {
		var ret []ApiKey
		return ret
	}
	return *o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyListResponse) GetApiKeysOk() ([]ApiKey, bool) {
	if o == nil || o.ApiKeys == nil {
		var ret []ApiKey
		return ret, false
	}
	return *o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *ApiKeyListResponse) HasApiKeys() bool {
	if o != nil && o.ApiKeys != nil {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []ApiKey and assigns it to the ApiKeys field.
func (o *ApiKeyListResponse) SetApiKeys(v []ApiKey) {
	o.ApiKeys = &v
}

type NullableApiKeyListResponse struct {
	Value        ApiKeyListResponse
	ExplicitNull bool
}

func (v NullableApiKeyListResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApiKeyListResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}