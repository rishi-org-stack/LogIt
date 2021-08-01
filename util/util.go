package util

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
)

func DropUnwantedField() {

}

func TransferData(o1, o2 interface{}) error {
	if reflect.TypeOf(o1) == reflect.TypeOf(o2) {
		copier.CopyWithOption(o2, o1, copier.Option{DeepCopy: true, IgnoreEmpty: true})
	}
	return fmt.Errorf("we have an error util/util.go %v", "types doesnt matches")
}
