// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"bytes"
	"encoding/json"

	"github.com/daichi-m/go18ds/containers"
	"github.com/daichi-m/go18ds/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map[string, string])(nil)
	var _ containers.JSONDeserializer = (*Map[string, string])(nil)
}

// ToJSON outputs the JSON representation of map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	it := m.Iterator()
	lastIndex := m.Size() - 1
	index := 0

	for it.Next() {
		km, err := json.Marshal(it.Key())
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(it.Value())
		if err != nil {
			return nil, err
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}

	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// FromJSON populates map from the input JSON representation.
//func (m *Map) FromJSON(data []byte) error {
//	elements := make(map[string]interface{})
//	err := json.Unmarshal(data, &elements)
//	if err == nil {
//		m.Clear()
//		for key, value := range elements {
//			m.Put(key, value)
//		}
//	}
//	return err
//}

// FromJSON populates map from the input JSON representation.
func (m *Map[string, V]) FromJSON(data []byte) error {
	elements := make(map[string]V)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	index := make(map[string]int)
	var keys []string
	for key := range elements {
		keys = append(keys, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(data, esc)
	}

	byIndex := func(key1, key2 string) int {
		index1 := index[key1]
		index2 := index[key2]
		return index1 - index2
	}

	utils.Sort(keys, byIndex)

	m.Clear()

	for _, key := range keys {
		m.Put(key, elements[key])
	}

	return nil
}
