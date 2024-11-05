package handler

import "errors"

var ERR_NO_DS = errors.New("Datastore not found")

type (
	registry struct {
		ds map[string]interface{}
	}

	dsInitFunc func() (string, interface{})
)

func (r *registry) Get(name string) (interface{}, error) {
	ds, ok := r.ds[name]
	if !ok {
		return nil, ERR_NO_DS
	}

	return ds, nil
}

func NewDsRegistry(dsInitFuncs ...dsInitFunc) *registry {
	dsMap := make(map[string]interface{})
	for _, init := range dsInitFuncs {
		name, ds := init()
		dsMap[name] = ds
	}
	return &registry{ds: dsMap}
}
