package modules

import (
	"fmt"
	"github.com/davidscholberg/go-i3barjson"
	"io/ioutil"
	"strings"
)

type Raid struct {
	BlockIndex     int    `yaml:"block_index"`
	UpdateInterval int    `yaml:"update_interval"`
	Label          string `yaml:"label"`
	UpdateSignal   int    `yaml:"update_signal"`
}

func (c Raid) GetBlockIndex() int {
	return c.BlockIndex
}

func (c Raid) GetUpdateFunc() func(b *i3barjson.Block, c BlockConfig) {
	return updateRaidBlock
}

func (c Raid) GetUpdateInterval() int {
	return c.UpdateInterval
}

func (c Raid) GetUpdateSignal() int {
	return c.UpdateSignal
}

func updateRaidBlock(b *i3barjson.Block, c BlockConfig) {
	cfg := c.(Raid)
	fullTextFmt := fmt.Sprintf("%s%%s", cfg.Label)
	mdstatPath := "/proc/mdstat"
	stats, err := ioutil.ReadFile(mdstatPath)
	if err != nil {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}
	i := strings.Index(string(stats), "_")
	if i != -1 {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, "degraded")
		return
	}
	b.Urgent = false
	b.FullText = fmt.Sprintf(fullTextFmt, "ok")
}
