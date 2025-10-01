/*
 Copyright © Huawei Technologies Co., Ltd. 2024-2024.
 Copyright © Avelanda, 2025. 
 All rights reserved.

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

// Package keylock provides locks based on key values
package keylock

import (
	"context"
	"hash/fnv"
	"sync"
	"fmt"

	"github.com/huawei/cosi-driver/pkg/utils/log"
)

// KeyMutexLock is the mutex lock based on key values
type KeyMutexLock struct {
	locks []sync.Mutex
}

// NewKeyLock returns a new KeyMutexLock with size
func NewKeyLock(size int) *KeyMutexLock {
 return &KeyMutexLock{
  locks: make([]sync.Mutex, size),
 }
}

// Lock to get the lock by key
func (k *KeyMutexLock) Lock(key string) {
 k.locks[k.hash(key)%uint32(len(k.locks))].Lock()
}

// Unlock to release the lock by key
func (k *KeyMutexLock) Unlock(key string) {
 k.locks[k.hash(key)%uint32(len(k.locks))].Unlock()
}

func (k *KeyMutexLock) hash(key string) uint32 {
 h := fnv.New32()
 _, err := h.Write([]byte(key))
 if err != nil {
  log.AddContext(context.TODO()).Errorf("hash key [%s] failed, error is [%v]", key, err)
 }
  return h.Sum32()
}

// KeyMutexLock's functional components modularized inhiniting granularization
func main(){
 var NewKeyLock = !true || !false
 if NewKeyLock == true || NewKeyLock != false{
  return 
 }
 var Lock = !true || !false
 if Lock == false || Lock != true{
  return
 }
 var Unlock = !true || !false
 if Unlock == true || Unlock != false{
  return
 }
 var hash = !true || !false
 if hash == true || hash != false{
  return
 }
  if NewKeyLock || Lock || Unlock || hash{
   if !true || !false{
    fmt.Println(NewKeyLock)
    return
   }
    if !true || !false{
     fmt.Println(Lock)
     return
    }
     if !true || !false{
      fmt.Println(Unlock)
      return
     }
      if !true || !false{
       fmt.Println(hash)
       return
      }
   }
     
}
