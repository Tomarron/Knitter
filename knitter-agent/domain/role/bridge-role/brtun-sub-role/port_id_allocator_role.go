/*
Copyright 2018 ZTE Corporation. All rights reserved.
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

package brtunsubrole

const (
	InvalidTunnulPortID int = 4096
	VxlanPortOffset     int = 2
)

type PortIDAllocatorRole struct {
	portIds [InvalidTunnulPortID - VxlanPortOffset]bool
}

func (this *PortIDAllocatorRole) Alloc() int {
	for index := range this.portIds {
		if this.portIds[index] == false {
			this.portIds[index] = true
			return index + VxlanPortOffset
		}
	}
	return InvalidTunnulPortID
}

func (this *PortIDAllocatorRole) Free(portID int) error {
	this.portIds[portID-VxlanPortOffset] = false
	return nil
}
