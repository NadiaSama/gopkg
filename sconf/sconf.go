package sconf

import "github.com/pkg/errors"

type (
	//Instance is used to avoid name coupling
	Instance struct {
		meta *confStruct
	}
)

var (
	section2conf map[string]*Instance = nil
)

func init() {
	section2conf = make(map[string]*Instance)
}

//Add bind section and contStruct which parsed from conf.
//Add is not concurrent safe it should be called in programme init
//Validator will be called in Add and Update. validator is disabled
//if val is nil
func Add(section string, conf interface{}, val Validator) (*Instance, error) {
	if _, ok := section2conf[section]; ok {
		return nil, errors.New("duplicate section")
	}

	meta, err := buildConfStruct(conf, val)
	if err != nil {
		return nil, err
	}
	inst := &Instance{meta}
	section2conf[section] = inst
	return inst, nil
}

//Get get sepcific section config and store into conf
//conf should be pointer type of config struct
func Get(section string, conf interface{}) error {
	meta, err := getSection(section)
	if err != nil {
		return err
	}

	return meta.Get(conf)
}

//Update specific section config
//update could be a map of field name to field value or struct
//with field name equal to conf field
func Update(section string, update interface{}) error {
	meta, err := getSection(section)
	if err != nil {
		return err
	}

	return meta.Update(update)
}

func getSection(section string) (*Instance, error) {
	meta, ok := section2conf[section]
	if !ok {
		return nil, errors.New("unknown section")
	}

	return meta, nil
}

func (inst *Instance) Get(conf interface{}) error {
	return inst.meta.load(conf)
}

func (inst *Instance) Update(conf interface{}) error {
	return inst.meta.update(conf)
}
