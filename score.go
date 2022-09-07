package elastic

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

type Score float64

var (
	posiInf = Score(math.Inf(0))
	negaInf = Score(math.Inf(-1))
)

const (
	posiInfExp = `"Infinity"`
	negaInfExp = `"-Infinity"`
)

const (
	posiInfStr = "Infinity"
	negaInfStr = "-Infinity"
)

func (s Score) MarshalJSON() ([]byte, error) {
	if s == posiInf {
		return []byte(posiInfExp), nil
	}
	if s == negaInf {
		return []byte(negaInfExp), nil
	}
	return []byte(strconv.FormatFloat(float64(s), 'g', -1, 64)), nil
}

func (s *Score) UnmarshalJSON(v []byte) error {
	var num float64
	numErr := json.Unmarshal(v, &num)
	if numErr == nil {
		*s = Score(num)
		return nil
	}
	var str string
	if err := json.Unmarshal(v, &str); err != nil {
		//NOTE: in this close, return numErr.
		// It excepts a number instead of a string, but it could not fulfilled.
		return numErr
	}
	if str == posiInfStr {
		*s = Score(math.Inf(0))
		return nil
	}
	if str == negaInfStr {
		*s = Score(math.Inf(-1))
		return nil
	}
	return fmt.Errorf("%#v could not be unmarshalled as elastic.Score: %w", v, numErr)
}

var _ json.Marshaler = Score(0)
var _ json.Marshaler = (*Score)(nil)

var _ json.Unmarshaler = (*Score)(nil)
