package util

import (
	"fmt"
	"github.com/tonyjt/tgo_v2/log"
	"strconv"
)

func Float64ToInt(value float64, multiplied float64) (intValue int, err error) {

	aString := fmt.Sprintf("%.0f", value*multiplied)

	intValue, err = strconv.Atoi(aString)

	if err != nil {
		log.Errorf("%f to int failed,error:%s", value, err.Error())
	}
	return
}

func Float64ToInt64(value float64, multiplied float64) (intValue int64, err error) {

	aString := fmt.Sprintf("%.0f", value*multiplied)

	intValue, err = strconv.ParseInt(aString, 10, 64)

	if err != nil {
		log.Errorf("%f to int failed,error:%s", value, err.Error())
	}
	return
}
