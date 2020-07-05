package sconf

import "github.com/pkg/errors"

var (
	section2conf map[string]*confStruct = make(map[string]*confStruct)
)

//Add bind section and contStruct which parsed from conf
//currently nested structure is not support
func Add(section string, conf interface{}) error {
	if _, ok := section2conf[section]; ok {
		return errors.New("duplicate section")
	}

	meta, err := buildConfStruct(conf)
	if err != nil {
		return err
	}
	section2conf[section] = meta
	return nil
}

//Load load sepcific section config and store into conf
//conf should be pointer type of config struct
func Load(section string, conf interface{}) error {
	meta, err := getSection(section)
	if err != nil {
		return err
	}

	return meta.load(conf)
}

//Update specific section config
func Update(section string, update map[string]interface{}) error {
	meta, err := getSection(section)
	if err != nil {
		return err
	}

	return meta.update(update)
}

func getSection(section string) (*confStruct, error) {
	meta, ok := section2conf[section]
	if !ok {
		return nil, errors.New("unknown section")
	}

	return meta, nil
}
