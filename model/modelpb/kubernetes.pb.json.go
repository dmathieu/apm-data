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

package modelpb

import "github.com/elastic/apm-data/model/internal/modeljson"

func (k *Kubernetes) toModelJSON(out *modeljson.Kubernetes) {
	*out = modeljson.Kubernetes{
		Namespace: k.Namespace,
		Node: modeljson.KubernetesNode{
			Name: k.NodeName,
		},
		Pod: modeljson.KubernetesPod{
			Name: k.PodName,
			UID:  k.PodUid,
		},
	}
}
