// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIMetaListItem API meta list item
//
// swagger:model APIMetaListItem
type APIMetaListItem struct {

	// api name
	APIName string `json:"apiName,omitempty"`

	// api type
	APIType string `json:"apiType,omitempty"`

	// gateway envs
	GatewayEnvs []string `json:"gateway-envs"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this API meta list item
func (m *APIMetaListItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this API meta list item based on context it is used
func (m *APIMetaListItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIMetaListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIMetaListItem) UnmarshalBinary(b []byte) error {
	var res APIMetaListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
