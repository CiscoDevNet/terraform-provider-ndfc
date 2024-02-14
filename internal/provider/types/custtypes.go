package types

import "strconv"

const (
	ValuesDeeplyEqual = iota
	RequiresReplace
	RequiresUpdate
	ControlFlagUpdate
)

type Int64Custom int64

func (i *Int64Custom) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "\"\"" {
		*i = -9223372036854775808
	} else {
		ss := string(data)
		// If the string is quoted, remove the quotes
		ssUn, err := strconv.Unquote(ss)
		if err == nil {
			// Quote removed
			ss = ssUn
		}
		ii, _ := strconv.ParseInt(ss, 10, 64)
		*i = Int64Custom(ii)
	}

	return nil
}

func (i Int64Custom) MarshalJSON() ([]byte, error) {
	res := ""
	res = strconv.FormatInt(int64(i), 10)
	return []byte(strconv.Quote(res)), nil

}
