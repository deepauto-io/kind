/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schedx

import "strings"

const (
	// Split is the split string.
	Split = "<SPLIT>"
)

type KindType string

const (
	API  KindType = "api"
	WEB  KindType = "web"
	TEAM KindType = "team"
)

func (k KindType) String() string {
	return string(k)
}

type GPTsType string

const (
	GPT3 GPTsType = "gpt-3"
	GPT4 GPTsType = "gpt-4"
)

func (t GPTsType) String() string {
	return string(t)
}

// Model return model name.
func Model(model string) string {
	if IsGPT4(model) {
		return GPT4.String()
	}
	return GPT3.String()
}

// IsGPT4 return true if model is gpt-4.
func IsGPT4(model string) bool {
	model = strings.ToLower(model)
	if strings.Contains(model, "gpt-4") {
		return true
	}
	if strings.Contains(model, "o1") {
		return true
	}
	if strings.Contains(model, "o3") {
		return true
	}
	if strings.Contains(model, "o") {
		return true
	}
	return false
}
