// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treebidimap

import (
	"encoding/json"

	"github.com/rahul1534/go18ds/containers"
	"github.com/rahul1534/go18ds/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map[string, string])(nil)
	var _ containers.JSONDeserializer = (*Map[string, string])(nil)
}

// ToJSON outputs the JSON representation of the map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[string]V)
	it := m.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[string, V]) FromJSON(data []byte) error {
	elements := make(map[string]V)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.Put(key, value)
		}
	}
	return err
}
