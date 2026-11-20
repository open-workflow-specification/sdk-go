// Copyright 2025 The Open Workflow Specification Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

var _ Object = &ForIn{}

// ForTask represents a task configuration to iterate over a collection.
type ForTask struct {
	TaskBase `json:",inline"`     // Inline TaskBase fields
	For      ForTaskConfiguration `json:"for" validate:"required"`
	While    string               `json:"while,omitempty"`
	Do       *TaskList            `json:"do" validate:"required,dive"`
}

func (f *ForTask) GetBase() *TaskBase {
	return &f.TaskBase
}

// ForIn represents either a runtime expression string or an inline array to enumerate.
type ForIn struct {
	Value interface{} `json:"-" validate:"for_in"`
}

func NewForIn(value interface{}) *ForIn {
	return &ForIn{Value: value}
}

func (f *ForIn) String() string {
	if f == nil || f.Value == nil {
		return ""
	}
	return fmt.Sprintf("%v", f.Value)
}

func (f *ForIn) GetValue() interface{} {
	if f == nil {
		return nil
	}
	return f.Value
}

func (f *ForIn) AsSlice() []interface{} {
	if f == nil {
		return nil
	}
	if s, ok := f.Value.([]interface{}); ok {
		return s
	}
	return nil
}

func (f *ForIn) AsString() string {
	if f == nil {
		return ""
	}
	if s, ok := f.Value.(string); ok {
		return s
	}
	if expr, ok := f.Value.(RuntimeExpression); ok {
		return expr.String()
	}
	return ""
}

func (f *ForIn) UnmarshalJSON(data []byte) error {
	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		f.Value = asString
		return nil
	}

	var asSlice []interface{}
	if err := json.Unmarshal(data, &asSlice); err == nil {
		f.Value = asSlice
		return nil
	}

	return errors.New("ForIn must be a runtime expression string or an array")
}

func (f *ForIn) MarshalJSON() ([]byte, error) {
	if f == nil || f.Value == nil {
		return nil, errors.New("ForIn cannot be nil")
	}
	switch v := f.Value.(type) {
	case RuntimeExpression:
		return json.Marshal(v.String())
	default:
		return json.Marshal(v)
	}
}

// ForTaskConfiguration defines the loop configuration for iterating over a collection.
type ForTaskConfiguration struct {
	Each string `json:"each,omitempty"`                // Variable name for the current item
	In   *ForIn `json:"in" validate:"required,for_in"` // Runtime expression or inline array for the collection
	At   string `json:"at,omitempty"`                  // Variable name for the current index
}
