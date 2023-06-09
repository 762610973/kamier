package core

import (
	"sync"

	zlog "compute/log"
)

type value []byte

type prepareValue struct {
	sync.Mutex
	preparedValue map[int64]value
	senderMap     map[int64][]chan value
}

func newPreparedValue() *prepareValue {
	p := new(prepareValue)
	p.preparedValue = make(map[int64]value)
	p.senderMap = make(map[int64][]chan value)
	return p
}

func (p *prepareValue) fetchValue(step int64, pb chan value) {
	zlog.Debug("another node fetch value from self node")
	p.Lock()
	defer p.Unlock()
	v, ok := p.preparedValue[step]
	if ok {
		zlog.Debug("already prepared value")
		pb <- v
	} else {
		zlog.Debug("haven't prepare value, so put a channel to sender list")
		if _, ok = p.senderMap[step]; !ok {
			zlog.Debug("sender list not exists")
			p.senderMap[step] = make([]chan value, 0, 1)
			p.senderMap[step] = append(p.senderMap[step], pb)
		} else {
			zlog.Debug("sender map exists")
			p.senderMap[step] = append(p.senderMap[step], pb)
		}
	}
}

func (p *prepareValue) prepareValue(step int64, v value) {
	p.Lock()
	defer p.Unlock()
	zlog.Debug("prepare value......")
	_, ok := p.preparedValue[step]
	if !ok {
		zlog.Debug("[value] is not exists, so put value to prepare map")
		p.preparedValue[step] = v
	}
	sendList, ok := p.senderMap[step]
	if ok {
		zlog.Debug("sender map is existed, so range list to send value")
		for _, val := range sendList {
			val <- v
		}
		delete(p.senderMap, step)
	}
	zlog.Debug("prepare value complete")

}
