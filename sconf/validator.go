package sconf

type (
	Validator interface {
		//Validate return nil if the config validate success
		//or error message if validate fail. v is pointer
		//of config struct
		Validate(v interface{}) error
	}

	//nilValidator Validate alway return nil
	NilValidator struct {
	}
)

var (
	NoValidate = &NilValidator{}
)

func (nv *NilValidator) Validate(v interface{}) error {
	return nil
}
