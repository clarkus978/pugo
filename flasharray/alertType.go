// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Alert struct {
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}
