package gen

import (
	"math/rand"

	"github.com/qiangyt/loggen/pkg/config"
)

type PidGeneratorT struct {
	pIdArray []uint32
}

type PidGenerator = *PidGeneratorT

func NewPidGenerator(cfg config.Pid) PidGenerator {
	pIdArray := []uint32{}
	for idx := 0; idx < int(cfg.Amount); idx++ {
		pIdArange := int32(cfg.End - cfg.Begin)
		pId := cfg.Begin + uint32(rand.Int31n(pIdArange))
		pIdArray = append(pIdArray, pId)
	}

	return &PidGeneratorT{pIdArray}
}

func (me PidGenerator) Next() uint32 {
	return uint32(rand.Intn(len(me.pIdArray)))
}
