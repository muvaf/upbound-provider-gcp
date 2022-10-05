/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EntryFulfillmentMessagesObservation struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []EntryFulfillmentMessagesTextObservation `json:"text,omitempty" tf:"text,omitempty"`
}

type EntryFulfillmentMessagesParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []EntryFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type EntryFulfillmentMessagesTextObservation struct {

	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type EntryFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type EntryFulfillmentObservation struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []EntryFulfillmentMessagesObservation `json:"messages,omitempty" tf:"messages,omitempty"`
}

type EntryFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []EntryFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type EventHandlersTriggerFulfillmentMessagesObservation struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []TriggerFulfillmentMessagesTextObservation `json:"text,omitempty" tf:"text,omitempty"`
}

type EventHandlersTriggerFulfillmentMessagesParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []TriggerFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type EventHandlersTriggerFulfillmentObservation struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []EventHandlersTriggerFulfillmentMessagesObservation `json:"messages,omitempty" tf:"messages,omitempty"`
}

type EventHandlersTriggerFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []EventHandlersTriggerFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type FillBehaviorObservation struct {

	// The fulfillment to provide the initial prompt that the agent can present to the user in order to fill the parameter.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InitialPromptFulfillment []InitialPromptFulfillmentObservation `json:"initialPromptFulfillment,omitempty" tf:"initial_prompt_fulfillment,omitempty"`
}

type FillBehaviorParameters struct {

	// The fulfillment to provide the initial prompt that the agent can present to the user in order to fill the parameter.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InitialPromptFulfillment []InitialPromptFulfillmentParameters `json:"initialPromptFulfillment,omitempty" tf:"initial_prompt_fulfillment,omitempty"`
}

type FormObservation struct {

	// Parameters to collect from the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Parameters []FormParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FormParameters struct {

	// Parameters to collect from the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Parameters []FormParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FormParametersObservation struct {

	// Defines fill behavior for the parameter.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	FillBehavior []FillBehaviorObservation `json:"fillBehavior,omitempty" tf:"fill_behavior,omitempty"`
}

type FormParametersParameters struct {

	// The human-readable name of the parameter, unique within the form.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The entity type of the parameter.
	// Format: projects/-/locations/-/agents/-/entityTypes/ for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects//locations//agents//entityTypes/ for developer entity types.
	// +kubebuilder:validation:Optional
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// Defines fill behavior for the parameter.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	FillBehavior []FillBehaviorParameters `json:"fillBehavior,omitempty" tf:"fill_behavior,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kubebuilder:validation:Optional
	IsList *bool `json:"isList,omitempty" tf:"is_list,omitempty"`

	// Indicates whether the parameter content should be redacted in log.
	// If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	// +kubebuilder:validation:Optional
	Redact *bool `json:"redact,omitempty" tf:"redact,omitempty"`

	// Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
	// Required parameters must be filled before form filling concludes.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type InitialPromptFulfillmentMessagesObservation struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []InitialPromptFulfillmentMessagesTextObservation `json:"text,omitempty" tf:"text,omitempty"`
}

type InitialPromptFulfillmentMessagesParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []InitialPromptFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type InitialPromptFulfillmentMessagesTextObservation struct {

	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type InitialPromptFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type InitialPromptFulfillmentObservation struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []InitialPromptFulfillmentMessagesObservation `json:"messages,omitempty" tf:"messages,omitempty"`
}

type InitialPromptFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []InitialPromptFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type PageEventHandlersObservation struct {

	// The unique identifier of this event handler.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []EventHandlersTriggerFulfillmentObservation `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type PageEventHandlersParameters struct {

	// The name of the event to handle.
	// +kubebuilder:validation:Optional
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// The target flow to transition to.
	// Format: projects//locations//agents//flows/.
	// +kubebuilder:validation:Optional
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects//locations//agents//flows//pages/.
	// +kubebuilder:validation:Optional
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []EventHandlersTriggerFulfillmentParameters `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type PageObservation struct {

	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EntryFulfillment []EntryFulfillmentObservation `json:"entryFulfillment,omitempty" tf:"entry_fulfillment,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EventHandlers []PageEventHandlersObservation `json:"eventHandlers,omitempty" tf:"event_handlers,omitempty"`

	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Form []FormObservation `json:"form,omitempty" tf:"form,omitempty"`

	// an identifier for the resource with format {{parent}}/pages/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the page.
	// Format: projects//locations//agents//flows//pages/.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TransitionRoutes []PageTransitionRoutesObservation `json:"transitionRoutes,omitempty" tf:"transition_routes,omitempty"`
}

type PageParameters struct {

	// The human-readable name of the page, unique within the agent.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EntryFulfillment []EntryFulfillmentParameters `json:"entryFulfillment,omitempty" tf:"entry_fulfillment,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EventHandlers []PageEventHandlersParameters `json:"eventHandlers,omitempty" tf:"event_handlers,omitempty"`

	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Form []FormParameters `json:"form,omitempty" tf:"form,omitempty"`

	// The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The flow to create a page for.
	// Format: projects//locations//agents//flows/.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("start_flow",true)
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects//locations//agents//flows//transitionRouteGroups/.
	// +kubebuilder:validation:Optional
	TransitionRouteGroups []*string `json:"transitionRouteGroups,omitempty" tf:"transition_route_groups,omitempty"`

	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TransitionRoutes []PageTransitionRoutesParameters `json:"transitionRoutes,omitempty" tf:"transition_routes,omitempty"`
}

type PageTransitionRoutesObservation struct {

	// The unique identifier of this transition route.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []PageTransitionRoutesTriggerFulfillmentObservation `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type PageTransitionRoutesParameters struct {

	// The condition to evaluate against form parameters or session parameters.
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// The unique identifier of an Intent.
	// Format: projects//locations//agents//intents/. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +kubebuilder:validation:Optional
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// The target flow to transition to.
	// Format: projects//locations//agents//flows/.
	// +kubebuilder:validation:Optional
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects//locations//agents//flows//pages/.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/dialogflowcx/v1beta1.Page
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// Reference to a Page in dialogflowcx to populate targetPage.
	// +kubebuilder:validation:Optional
	TargetPageRef *v1.Reference `json:"targetPageRef,omitempty" tf:"-"`

	// Selector for a Page in dialogflowcx to populate targetPage.
	// +kubebuilder:validation:Optional
	TargetPageSelector *v1.Selector `json:"targetPageSelector,omitempty" tf:"-"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []PageTransitionRoutesTriggerFulfillmentParameters `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type PageTransitionRoutesTriggerFulfillmentObservation struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []TransitionRoutesTriggerFulfillmentMessagesObservation `json:"messages,omitempty" tf:"messages,omitempty"`
}

type PageTransitionRoutesTriggerFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Messages []TransitionRoutesTriggerFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects//locations//agents//webhooks/.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesObservation struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []TransitionRoutesTriggerFulfillmentMessagesTextObservation `json:"text,omitempty" tf:"text,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []TransitionRoutesTriggerFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesTextObservation struct {

	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type TriggerFulfillmentMessagesTextObservation struct {

	// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type TriggerFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

// PageSpec defines the desired state of Page
type PageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PageParameters `json:"forProvider"`
}

// PageStatus defines the observed state of Page.
type PageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Page is the Schema for the Pages API. A Dialogflow CX conversation (session) can be described and visualized as a state machine.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Page struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PageSpec   `json:"spec"`
	Status            PageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PageList contains a list of Pages
type PageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Page `json:"items"`
}

// Repository type metadata.
var (
	Page_Kind             = "Page"
	Page_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Page_Kind}.String()
	Page_KindAPIVersion   = Page_Kind + "." + CRDGroupVersion.String()
	Page_GroupVersionKind = CRDGroupVersion.WithKind(Page_Kind)
)

func init() {
	SchemeBuilder.Register(&Page{}, &PageList{})
}