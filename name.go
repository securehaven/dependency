package dependency

import "reflect"

func Name(dep any) string {
	return reflect.TypeOf(dep).String()
}
