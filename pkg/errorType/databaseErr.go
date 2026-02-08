package errortype

import (
	"errors"
	"fmt"
	"strings"
)

var ConstraintErrType = errors.New("duplicate")
var NotFoundErrType = errors.New("not found")

func ConstraintErr(err error, constraints ...string) error {
	if err == nil {
		return nil
	}
	if !strings.Contains(err.Error(), "2067") {
		return err
	}
	return fmt.Errorf("%w: %s", ConstraintErrType, strings.Join(constraints, ","))
}

func NotFoundErr(field string) error {
	return fmt.Errorf("%s %w", field, NotFoundErrType)
}
