// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package modeljson

type Message struct {
	Body       string              `json:"body,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	Age        MessageAge          `json:"age,omitempty"`
	Queue      MessageQueue        `json:"queue,omitempty"`
	RoutingKey string              `json:"routing_key,omitempty"`
}

type MessageAge struct {
	Millis *int64 `json:"ms,omitempty"`
}

func (a MessageAge) isZero() bool {
	return a.Millis == nil
}

type MessageQueue struct {
	Name string `json:"name,omitempty"`
}

func (q MessageQueue) isZero() bool {
	return q.Name == ""
}
