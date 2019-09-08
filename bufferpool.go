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
	"fmt"
	"sync"
)

type Buffer []byte

// const PacketBufferSize = 0x10000

type Pool struct {
	mutex    sync.Mutex `webformhide:"" stringformhide:""`
	name     string
	buffPool []Buffer
	length   int
	count    int
}

func New(name string, length, count int) *Pool {
	if name == "" {
		name = "Pool"
	}
	return &Pool{
		name:     name,
		buffPool: make([]Buffer, 0, count),
		length:   length,
		count:    count,
	}
}

func (p *Pool) String() string {
	return fmt.Sprintf("%s[%v %v/%v]",
		p.name, p.length, len(p.buffPool), p.count,
	)
}

func (p *Pool) Get() Buffer {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	var rtn Buffer
	if l := len(p.buffPool); l > 0 {
		rtn = p.buffPool[l-1]
		p.buffPool = p.buffPool[:l-1]
	} else {
		rtn = make(Buffer, p.length)
	}
	return rtn
}

func (p *Pool) Put(pb Buffer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.buffPool) < p.count {
		p.buffPool = append(p.buffPool, pb)
	}
}
