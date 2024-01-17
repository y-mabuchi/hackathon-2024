package model

type ChildGender string

const (
	ChildGenderBoy  ChildGender = "CHILD_GENDER_BOY"
	ChildGenderGirl             = "CHILD_GENDER_GIRL"
)

func (c ChildGender) String() string {
	return string(c)
}

func (c ChildGender) Valid() error {
	switch c {
	case ChildGenderBoy, ChildGenderGirl:
		return nil
	default:
		return NewErrInvalidArgument(
			"child gender type",
			string(c),
		)
	}
}

func ParseChildGender(s string) (ChildGender, error) {
	cg := ChildGender(s)
	err := cg.Valid()
	if err != nil {
		return cg, err
	}

	return cg, nil
}
