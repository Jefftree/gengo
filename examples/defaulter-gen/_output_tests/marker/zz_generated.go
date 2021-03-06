// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package marker

import (
	"encoding/json"
	"reflect"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Defaulted{}, func(obj interface{}) { SetObjectDefaults_Defaulted(obj.(*Defaulted)) })
	scheme.AddTypeDefaultingFunc(&DefaultedWithFunction{}, func(obj interface{}) { SetObjectDefaults_DefaultedWithFunction(obj.(*DefaultedWithFunction)) })
	return nil
}

func SetObjectDefaults_Defaulted(in *Defaulted) {
	if reflect.ValueOf(in.Field).IsZero() {
		in.Field = "bar"
	}
	if reflect.ValueOf(in.OtherField).IsZero() {
		in.OtherField = 0
	}
	if reflect.ValueOf(in.List).IsZero() {
		if err := json.Unmarshal([]byte(`["foo", "bar"]`), &in.List); err != nil {
			panic(err)
		}
	}
	for i := range in.List {
		if reflect.ValueOf(in.List[i]).IsZero() {
			var ptrVar1 string = "apple"
			ptrVar0 := &ptrVar1
			in.List[i] = ptrVar0
		}
	}
	if reflect.ValueOf(in.Sub).IsZero() {
		if err := json.Unmarshal([]byte(`{"s": "foo", "i": 5}`), &in.Sub); err != nil {
			panic(err)
		}
	}
	if in.Sub != nil {
		if reflect.ValueOf(in.Sub.I).IsZero() {
			in.Sub.I = 1
		}
	}
	if reflect.ValueOf(in.StructList).IsZero() {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.StructList); err != nil {
			panic(err)
		}
	}
	for i := range in.StructList {
		a := &in.StructList[i]
		if reflect.ValueOf(a.I).IsZero() {
			a.I = 1
		}
	}
	if reflect.ValueOf(in.PtrStructList).IsZero() {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.PtrStructList); err != nil {
			panic(err)
		}
	}
	for i := range in.PtrStructList {
		a := in.PtrStructList[i]
		if a != nil {
			if reflect.ValueOf(a.I).IsZero() {
				a.I = 1
			}
		}
	}
	if reflect.ValueOf(in.StringList).IsZero() {
		if err := json.Unmarshal([]byte(`["foo"]`), &in.StringList); err != nil {
			panic(err)
		}
	}
	if reflect.ValueOf(in.OtherSub.I).IsZero() {
		in.OtherSub.I = 1
	}
	if reflect.ValueOf(in.Map).IsZero() {
		if err := json.Unmarshal([]byte(`{"foo": "bar"}`), &in.Map); err != nil {
			panic(err)
		}
	}
	for i_Map := range in.Map {
		if reflect.ValueOf(in.Map[i_Map]).IsZero() {
			var ptrVar1 string = "apple"
			ptrVar0 := &ptrVar1
			in.Map[i_Map] = ptrVar0
		}
	}
	if reflect.ValueOf(in.AliasPtr).IsZero() {
		var ptrVar1 string = "banana"
		ptrVar0 := &ptrVar1
		in.AliasPtr = ptrVar0
	}
}

func SetObjectDefaults_DefaultedWithFunction(in *DefaultedWithFunction) {
	SetDefaults_DefaultedWithFunction(in)
	if reflect.ValueOf(in.S1).IsZero() {
		in.S1 = "default_marker"
	}
	if reflect.ValueOf(in.S2).IsZero() {
		in.S2 = "default_marker"
	}
}
