package base

import "github.com/diegojromerolopez/congolway/pkg/neighborhood"

const DefaultRules = "23/3"
const DefaultRowLimitation = "limited"
const DefaultColLimitation = "limited"
const DefaultGridType = "dok"
const DefaultGeneration = 0
const DefaultNeighborhoodType = neighborhood.MOORE

// GolConf : configuration for Game of Life instances
type GolConf struct {
	rules            string
	gridType         string
	rowLimitation    string
	colLimitation    string
	generation       int
	neighborhoodType int
}

// NewDefaultGolConf : returns a default configuration
// for Game of Life instances
func NewDefaultGolConf() *GolConf {
	return &GolConf{
		DefaultRules,
		DefaultGridType,
		DefaultRowLimitation,
		DefaultColLimitation,
		DefaultGeneration,
		DefaultNeighborhoodType}
}

// NewGolConf : returns a default configuration
// but overwritting the fields that are passed
// as a map
func NewGolConf(overwrittenAttrs map[string]interface{}) *GolConf {
	gconf := &GolConf{
		DefaultRules,
		DefaultGridType,
		DefaultRowLimitation,
		DefaultColLimitation,
		DefaultGeneration,
		DefaultNeighborhoodType}

	if overwrittenAttrs["rules"] != nil {
		gconf.rules = overwrittenAttrs["rules"].(string)
	}
	if overwrittenAttrs["gridType"] != nil {
		gconf.gridType = overwrittenAttrs["gridType"].(string)
	}
	if overwrittenAttrs["rowLimitation"] != nil {
		gconf.rowLimitation = overwrittenAttrs["rowLimitation"].(string)
	}
	if overwrittenAttrs["colLimitation"] != nil {
		gconf.colLimitation = overwrittenAttrs["colLimitation"].(string)
	}
	if overwrittenAttrs["generation"] != nil {
		gconf.generation = overwrittenAttrs["generation"].(int)
	}
	if overwrittenAttrs["neighborhoodType"] != nil {
		gconf.neighborhoodType = overwrittenAttrs["neighborhoodType"].(int)
	}
	return gconf
}

func (gc *GolConf) Rules() string {
	return gc.rules
}

func (gc *GolConf) GridType() string {
	return gc.gridType
}

func (gc *GolConf) RowLimitation() string {
	return gc.rowLimitation
}

func (gc *GolConf) ColLimitation() string {
	return gc.colLimitation
}

func (gc *GolConf) Generation() int {
	return gc.generation
}

func (gc *GolConf) NeighborhoodType() int {
	return gc.neighborhoodType
}
