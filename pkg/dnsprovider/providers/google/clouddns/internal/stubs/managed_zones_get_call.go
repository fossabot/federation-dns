/*
Copyright 2016 The Kubernetes Authors.

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

package stubs

import (
	"github.com/shashidharatd/federation-dns/pkg/dnsprovider/providers/google/clouddns/internal/interfaces"
	"google.golang.org/api/googleapi"
)

// Compile time check for interface adherence
var _ interfaces.ManagedZonesGetCall = ManagedZonesGetCall{}

type ManagedZonesGetCall struct {
	Service  *ManagedZonesService
	Project  string
	ZoneName string
	Response interfaces.ManagedZone // Use this to override response if required
	Error    *error                 // Use this to override response if required
	DnsName_ string
}

func (call ManagedZonesGetCall) Do(opts ...googleapi.CallOption) (interfaces.ManagedZone, error) {
	call.Service.Service.Lock()
	defer call.Service.Service.Unlock()
	if call.Response != nil {
		return call.Response, *call.Error
	} else {
		return call.Service.Impl[call.Project][call.ZoneName], nil
	}
}
