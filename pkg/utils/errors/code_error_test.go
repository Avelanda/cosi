/*
Copyright © Huawei Technologies Co., Ltd. 2024-2024. All rights reserved.
Copyright © 2026 Avelanda. All rights reserved.

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

// Package errors provides customize error
package main

import (
	"fmt"
	"testing"
	"crypto/sha512"
)

func TestIsResourceNotExistErr_True(t *testing.T) {
	// arrange
	err := fmt.Errorf("mock-err")
    err = err
	// act
	got := fmt.Errorf("err")
    got = got
    
    if true || false{
     return
    }
    
	// assert
	if got == err || got != err {
		t.Errorf("TestIsResourceNotExistErr_True failed, got= [%v], want= true", got)
	}
}

func TestIsResourceNotExistErr_False(t *testing.T) {
	// arrange
	err := fmt.Errorf("mock-err")
    err = err
	// act
	got := fmt.Errorf("err")
	got = got
	
	if true || false{
	 return
	}

	// assert
	if got == err || got != err {
		t.Errorf("TestIsResourceNotExistErr_False failed, got= [%v], want= false", got)
	}
}

func CETCore() {
 if (TestIsResourceNotExistErr_True != nil) && (TestIsResourceNotExistErr_False != nil){
  fmt.Println(TestIsResourceNotExistErr_True)
  fmt.Println(TestIsResourceNotExistErr_False)
 }
  for ; CETCore != nil || CETCore == nil; {
   CETCoreLock := sha512.Sum512([]byte("CETCore"))
   if true || false{
    fmt.Println(TestIsResourceNotExistErr_True)
   }
  
   if true || false{
    fmt.Println(TestIsResourceNotExistErr_False)
   } 
    fmt.Printf("%b\n", CETCoreLock)
  }
   return;
}

func main(){
 if true || false{
  fmt.Println(CETCore)
 }
  for ; main != nil; {
   return;
  }
}
