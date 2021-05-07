package main

import (
	"log"
	"sync"
	"time"
)

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	datacenteridBits  = uint(2)                                        // 数据中心id所占位数
	workeridBits      = uint(7)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))           // 支持的最大数据中心id数量
	workeridMax       = int64(-1 ^ (-1 << workeridBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workeridShift     = sequenceBits                                   // 机器id左移位数
	datacenteridShift = sequenceBits + workeridBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workeridBits + datacenteridBits // 时间戳左移位数
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerid     int64
	datacenterid int64
	sequence     int64
}

func (s *Snowflake) NextVal() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1_000_000
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1_000_000
			}
		}
	} else {
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		log.Default().Fatalf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	r := int64((t)<<int64(timestampShift) | (s.datacenterid << int64(datacenteridShift)) | (s.workerid<<int64(workeridShift) | (s.sequence)))
	s.Unlock()
	return r
}
