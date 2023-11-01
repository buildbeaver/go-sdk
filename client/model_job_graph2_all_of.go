/*
BuildBeaver Dynamic Build API - OpenAPI 3.0

This is the BuildBeaver Dynamic Build API.

API version: 0.3.00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the JobGraph2AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobGraph2AllOf{}

// JobGraph2AllOf struct for JobGraph2AllOf
type JobGraph2AllOf struct {
	// The set of steps within the job
	Steps []Step `json:"steps"`
	AdditionalProperties map[string]interface{}
}

type _JobGraph2AllOf JobGraph2AllOf

// NewJobGraph2AllOf instantiates a new JobGraph2AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobGraph2AllOf(steps []Step) *JobGraph2AllOf {
	this := JobGraph2AllOf{}
	this.Steps = steps
	return &this
}

// NewJobGraph2AllOfWithDefaults instantiates a new JobGraph2AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobGraph2AllOfWithDefaults() *JobGraph2AllOf {
	this := JobGraph2AllOf{}
	return &this
}

// GetSteps returns the Steps field value
func (o *JobGraph2AllOf) GetSteps() []Step {
	if o == nil {
		var ret []Step
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *JobGraph2AllOf) GetStepsOk() ([]Step, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *JobGraph2AllOf) SetSteps(v []Step) {
	o.Steps = v
}

func (o JobGraph2AllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobGraph2AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["steps"] = o.Steps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobGraph2AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varJobGraph2AllOf := _JobGraph2AllOf{}

	if err = json.Unmarshal(bytes, &varJobGraph2AllOf); err == nil {
		*o = JobGraph2AllOf(varJobGraph2AllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "steps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobGraph2AllOf struct {
	value *JobGraph2AllOf
	isSet bool
}

func (v NullableJobGraph2AllOf) Get() *JobGraph2AllOf {
	return v.value
}

func (v *NullableJobGraph2AllOf) Set(val *JobGraph2AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJobGraph2AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJobGraph2AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobGraph2AllOf(val *JobGraph2AllOf) *NullableJobGraph2AllOf {
	return &NullableJobGraph2AllOf{value: val, isSet: true}
}

func (v NullableJobGraph2AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobGraph2AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

