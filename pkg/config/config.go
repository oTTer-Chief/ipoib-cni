/*
2022 NVIDIA CORPORATION & AFFILIATES

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

package config

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Mellanox/ipoib-cni/pkg/types"
)

// LoadConf parses and validates stdin netconf and returns NetConf object
func LoadConf(bytes []byte) (*types.NetConf, string, error) {
	n := &types.NetConf{}
	if err := json.Unmarshal(bytes, n); err != nil {
		return nil, "", fmt.Errorf("failed to load netconf: %v", err)
	}
	if n.Master == "" {
		return nil, "", fmt.Errorf("host master interface is missing")
	}
	if n.Pkey != "" {
		tmpPkey, err := strconv.ParseUint(n.Pkey, 0, 15)
		if err != nil {
			return nil, "", fmt.Errorf("invalid Pkey: %w", err)
		}
		if tmpPkey < 1 || tmpPkey > 32767 {
			return nil, "", fmt.Errorf("invalid Pkey %s (must be between 0x0001 and 0x7fff inclusive)", n.Pkey)
		}
	}
	return n, n.CNIVersion, nil
}
