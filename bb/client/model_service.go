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

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	// Unique name of the service, within the parent job.
	Name string `json:"name"`
	// Name of the Docker image for the service to run.
	Image string `json:"image"`
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
	AwsAuth *DockerAWSAuth `json:"aws_auth,omitempty"`
	// A list of environment variables to export prior to starting the service.
	Environment []EnvVar `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _Service Service

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(name string, image string, environment []EnvVar) *Service {
	this := Service{}
	this.Name = name
	this.Image = image
	this.Environment = environment
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetImage returns the Image field value
func (o *Service) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *Service) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *Service) SetImage(v string) {
	o.Image = v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *Service) GetBasicAuth() DockerBasicAuth {
	if o == nil || IsNil(o.BasicAuth) {
		var ret DockerBasicAuth
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBasicAuthOk() (*DockerBasicAuth, bool) {
	if o == nil || IsNil(o.BasicAuth) {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *Service) HasBasicAuth() bool {
	if o != nil && !IsNil(o.BasicAuth) {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given DockerBasicAuth and assigns it to the BasicAuth field.
func (o *Service) SetBasicAuth(v DockerBasicAuth) {
	o.BasicAuth = &v
}

// GetAwsAuth returns the AwsAuth field value if set, zero value otherwise.
func (o *Service) GetAwsAuth() DockerAWSAuth {
	if o == nil || IsNil(o.AwsAuth) {
		var ret DockerAWSAuth
		return ret
	}
	return *o.AwsAuth
}

// GetAwsAuthOk returns a tuple with the AwsAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAwsAuthOk() (*DockerAWSAuth, bool) {
	if o == nil || IsNil(o.AwsAuth) {
		return nil, false
	}
	return o.AwsAuth, true
}

// HasAwsAuth returns a boolean if a field has been set.
func (o *Service) HasAwsAuth() bool {
	if o != nil && !IsNil(o.AwsAuth) {
		return true
	}

	return false
}

// SetAwsAuth gets a reference to the given DockerAWSAuth and assigns it to the AwsAuth field.
func (o *Service) SetAwsAuth(v DockerAWSAuth) {
	o.AwsAuth = &v
}

// GetEnvironment returns the Environment field value
func (o *Service) GetEnvironment() []EnvVar {
	if o == nil {
		var ret []EnvVar
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Service) GetEnvironmentOk() ([]EnvVar, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment, true
}

// SetEnvironment sets field value
func (o *Service) SetEnvironment(v []EnvVar) {
	o.Environment = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["image"] = o.Image
	if !IsNil(o.BasicAuth) {
		toSerialize["basic_auth"] = o.BasicAuth
	}
	if !IsNil(o.AwsAuth) {
		toSerialize["aws_auth"] = o.AwsAuth
	}
	toSerialize["environment"] = o.Environment

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Service) UnmarshalJSON(bytes []byte) (err error) {
	varService := _Service{}

	if err = json.Unmarshal(bytes, &varService); err == nil {
		*o = Service(varService)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "image")
		delete(additionalProperties, "basic_auth")
		delete(additionalProperties, "aws_auth")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


