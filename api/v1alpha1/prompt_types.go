/*
Copyright 2023 KubeAGI.

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

package v1alpha1

import (
	llmzhipuai "github.com/kubeagi/arcadia/pkg/llms/zhipuai"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PromptSpec defines the desired state of Prompt
type PromptSpec struct {
	// LLM serivice name(CRD LLM)
	LLM string `json:"llm"`
	// ZhiPuAIParams defines the params of ZhiPuAI
	ZhiPuAIParams *llmzhipuai.ModelParams `json:"zhiPuAIParams,omitempty"`
}

// PromptStatus defines the observed state of Prompt
type PromptStatus struct {
	// ConditionedStatus is the current status
	ConditionedStatus `json:",inline"`

	// Data retrieved after LLM Call
	Data []byte `json:"data"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Prompt is the Schema for the prompts API
type Prompt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PromptSpec   `json:"spec,omitempty"`
	Status PromptStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PromptList contains a list of Prompt
type PromptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Prompt `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Prompt{}, &PromptList{})
}