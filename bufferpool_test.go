// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufferpool

import (
	"testing"
)

func TestPool_Get(t *testing.T) {
	p := New("bufferpool", 0x10000, 100)
	b := p.Get()
	t.Logf("%v", p)
	p.Put(b)
	t.Logf("%v", p)
	b = p.Get()
	t.Logf("%v", p)
	p.Put(b)
	t.Logf("%v", p)
}
