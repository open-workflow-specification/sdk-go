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

import "encoding/json"

type CallHTTP struct {
	TaskBase `json:",inline"` // Inline TaskBase fields
	Call     string           `json:"call" validate:"required,eq=http"`
	With     HTTPArguments    `json:"with" validate:"required"`
}

func (c *CallHTTP) GetBase() *TaskBase {
	return &c.TaskBase
}

type HTTPArguments struct {
	Method   string               `json:"method" validate:"required,oneofci=GET POST PUT DELETE PATCH"`
	Endpoint *Endpoint            `json:"endpoint" validate:"required"`
	Headers  *ObjectOrRuntimeExpr `json:"headers,omitempty" validate:"omitempty,object_or_runtime_expr"`
	Body     json.RawMessage      `json:"body,omitempty"`
	Query    *ObjectOrRuntimeExpr `json:"query,omitempty" validate:"omitempty,object_or_runtime_expr"`
	Output   string               `json:"output,omitempty" validate:"omitempty,oneof=raw content response"`
	Redirect bool                 `json:"redirect,omitempty"`
}

type CallOpenAPI struct {
	TaskBase `json:",inline"` // Inline TaskBase fields
	Call     string           `json:"call" validate:"required,eq=openapi"`
	With     OpenAPIArguments `json:"with" validate:"required"`
}

func (c *CallOpenAPI) GetBase() *TaskBase {
	return &c.TaskBase
}

type OpenAPIArguments struct {
	Document       *ExternalResource                  `json:"document" validate:"required"`
	OperationID    string                             `json:"operationId" validate:"required"`
	Parameters     map[string]interface{}             `json:"parameters,omitempty"`
	Authentication *ReferenceableAuthenticationPolicy `json:"authentication,omitempty"`
	Output         string                             `json:"output,omitempty" validate:"omitempty,oneof=raw content response"`
	Redirect       bool                               `json:"redirect,omitempty"`
}

type CallGRPC struct {
	TaskBase `json:",inline"`
	Call     string        `json:"call" validate:"required,eq=grpc"`
	With     GRPCArguments `json:"with" validate:"required"`
}

func (c *CallGRPC) GetBase() *TaskBase {
	return &c.TaskBase
}

type GRPCArguments struct {
	Proto          *ExternalResource                  `json:"proto" validate:"required"`
	Service        GRPCService                        `json:"service" validate:"required"`
	Method         string                             `json:"method" validate:"required"`
	Arguments      map[string]interface{}             `json:"arguments,omitempty"`
	Authentication *ReferenceableAuthenticationPolicy `json:"authentication,omitempty" validate:"omitempty"`
}

type GRPCService struct {
	Name           string                             `json:"name" validate:"required"`
	Host           string                             `json:"host" validate:"required,hostname_rfc1123"`
	Port           int                                `json:"port" validate:"required,min=0,max=65535"`
	Authentication *ReferenceableAuthenticationPolicy `json:"authentication,omitempty"`
}

type CallAsyncAPI struct {
	TaskBase `json:",inline"`
	Call     string            `json:"call" validate:"required,eq=asyncapi"`
	With     AsyncAPIArguments `json:"with" validate:"required"`
}

func (c *CallAsyncAPI) GetBase() *TaskBase {
	return &c.TaskBase
}

type AsyncAPIArguments struct {
	Document       *ExternalResource                  `json:"document" validate:"required"`
	Channel        string                             `json:"channel,omitempty"`
	Operation      string                             `json:"operation,omitempty"`
	Server         *AsyncAPIServer                    `json:"server,omitempty"`
	Protocol       string                             `json:"protocol,omitempty" validate:"omitempty,oneof=amqp amqp1 anypointmq googlepubsub http ibmmq jms kafka mercure mqtt mqtt5 nats pulsar redis sns solace sqs stomp ws"`
	Message        *AsyncAPIOutboundMessage           `json:"message,omitempty"`
	Subscription   *AsyncAPISubscription              `json:"subscription,omitempty"`
	Authentication *ReferenceableAuthenticationPolicy `json:"authentication,omitempty" validate:"omitempty"`
}

type AsyncAPIServer struct {
	Name      string                 `json:"name" validate:"required"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type AsyncAPIOutboundMessage struct {
	Payload map[string]interface{} `json:"payload,omitempty" validate:"omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty" validate:"omitempty"`
}

type AsyncAPISubscription struct {
	Filter  *RuntimeExpression                `json:"filter,omitempty"`
	Consume *AsyncAPIMessageConsumptionPolicy `json:"consume" validate:"required"`
}

type AsyncAPIMessageConsumptionPolicy struct {
	For    *Duration          `json:"for,omitempty"`
	Amount int                `json:"amount,omitempty" validate:"required_without_all=While Until"`
	While  *RuntimeExpression `json:"while,omitempty" validate:"required_without_all=Amount Until"`
	Until  *RuntimeExpression `json:"until,omitempty" validate:"required_without_all=Amount While"`
}

type CallA2A struct {
	TaskBase `json:",inline"`
	Call     string       `json:"call" validate:"required,eq=a2a"`
	With     A2AArguments `json:"with" validate:"required"`
}

func (c *CallA2A) GetBase() *TaskBase {
	return &c.TaskBase
}

type A2AArguments struct {
	AgentCard  *ExternalResource    `json:"agentCard,omitempty"`
	Server     *Endpoint            `json:"server,omitempty"`
	Method     string               `json:"method" validate:"required,oneof='message/send' 'message/stream' 'tasks/get' 'tasks/list' 'tasks/cancel' 'tasks/resubscribe' 'tasks/pushNotificationConfig/set' 'tasks/pushNotificationConfig/get' 'tasks/pushNotificationConfig/list' 'tasks/pushNotificationConfig/delete' 'agent/getAuthenticatedExtendedCard'"`
	Parameters *ObjectOrRuntimeExpr `json:"parameters,omitempty"`
}

type CallMCP struct {
	TaskBase `json:",inline"`
	Call     string       `json:"call" validate:"required,eq=mcp"`
	With     MCPArguments `json:"with" validate:"required"`
}

func (c *CallMCP) GetBase() *TaskBase {
	return &c.TaskBase
}

type MCPArguments struct {
	ProtocolVersion string               `json:"protocolVersion,omitempty"`
	Method          string               `json:"method" validate:"required,oneof='tools/list' 'tools/call' 'prompts/list' 'prompts/get' 'resources/list' 'resources/read' 'resources/templates/list'"`
	Parameters      *ObjectOrRuntimeExpr `json:"parameters,omitempty"`
	Timeout         *Duration            `json:"timeout,omitempty"`
	Transport       *McpCallTransport    `json:"transport" validate:"required,mcp_transport"`
	Client          *McpClient           `json:"client,omitempty"`
}

type McpCallTransport struct {
	HTTP    *McpHttpTransport  `json:"http,omitempty"`
	Stdio   *McpStdioTransport `json:"stdio,omitempty"`
	Options map[string]string  `json:"options,omitempty"`
}

type McpHttpTransport struct {
	Endpoint *Endpoint         `json:"endpoint" validate:"required"`
	Headers  map[string]string `json:"headers,omitempty"`
}

type McpStdioTransport struct {
	Command     string            `json:"command" validate:"required"`
	Arguments   []string          `json:"arguments,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
}

type McpClient struct {
	Name    string `json:"name" validate:"required"`
	Version string `json:"version" validate:"required"`
}

type CallFunction struct {
	TaskBase `json:",inline"`       // Inline TaskBase fields
	Call     string                 `json:"call" validate:"required,ne=http,ne=openapi,ne=grpc,ne=asyncapi,ne=a2a,ne=mcp"`
	With     map[string]interface{} `json:"with,omitempty"`
}

func (c *CallFunction) GetBase() *TaskBase {
	return &c.TaskBase
}
