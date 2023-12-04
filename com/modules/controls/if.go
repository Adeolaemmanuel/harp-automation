package controls

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

//Control
func (I If) IfControl(operator string, value1, value2 any) (bool, error) {
	switch operator {
	case IsEquals:
		return I.IsEquals(value1, value2)
	case IsStringContain:
		return I.IsStringContain(value1.(string), value2.(string))
	case IsStringNotContain:
		return I.IsStringNotContain(value1.(string), value2.(string))
	case IsStringEndWith:
		return I.IsStringEndWith(value1.(string), value2.(string))
	case IsStringStartWith:
		return I.IsStringStartWith(value1.(string), value2.(string))
	case IsLargerEqualNumber:
		return I.IsLargerEqualNumber(value1.(int), value2.(int))
	case IsLesserEqualNumber:
		return I.IsLesserEqualNumber(value1.(int), value2.(int))
	case IsLesserNumber:
		return I.IsLesserNumber(value1.(int), value2.(int))
	case IsLargerNumber:
		return I.IsLargerNumber(value1.(int), value2.(int))
	case IsStringEmpty:
		return I.isStringEmpty(value1.(string))
	case IsBetween:
		return I.IsBetweenValue(value1, value2)
	default:
		return false, errors.New("invalid operator found")
	}
}

// Check if sting is empty
func (I If) isStringEmpty(value string) (bool, error) {
	isEmpty := len(value) > 0

	if !isEmpty {
		return true, errors.New("string is empty")
	}

	return false, nil
}

// Check if value is string
func (I If) checkIfString(value any) bool {
	
	if I.checkIfDateTime(value) {
		return false
	}

	return reflect.TypeOf(value.(string)).Kind() == reflect.String
}

// Check if value is int
func (I If) checkIfInt(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int
}

// Check if value is date/time
func (I If) checkIfDateTime(value any) bool {

	if I.checkIfInt(value) {
		return I.checkIfInt(value)
	}

	_, err := time.Parse(time.RFC3339, value.(string))

	return err == nil
}

// Check if the value1 is equal value2
func (I If) IsEquals(value1 any, value2 any) (bool, error) {

	if I.checkIfInt(value1) && I.checkIfInt(value2) {

		return I.IsEqualsToNumber(value1.(int), value2.(int))

	} else if I.checkIfString(value1) && I.checkIfString(value2) {

		return I.IsEqualsToString(value1, value2)

	} else if I.checkIfDateTime(value1) && I.checkIfDateTime(value2) {

		return I.IsTimeEqual(value1, value2)

	}  else {

		return false, errors.New("invalid Equality comparison")

	}
}

// Check if the values is not equal to value2
func (I If) IsNotEquals(value1 any, value2 any) (bool, error) {

	if I.checkIfInt(value1) && I.checkIfInt(value2) {

		ok, err:= I.IsEqualsToNumber(value1.(int), value2.(int))

		return !ok, err

	} else if I.checkIfString(value1) && I.checkIfString(value2) {

		ok, err:=  I.IsEqualsToString(value1, value2)
		
		return !ok, err

	} else if I.checkIfDateTime(value1) && I.checkIfDateTime(value2) {

		ok, err:=  I.IsTimeEqual(value1, value2)
		
		return !ok, err

	}  else {

		return false, errors.New("invalid Equality comparison")

	}
}

// Check if the values are equal string
func (I If) IsEqualsToString(value1, value2 any) (bool, error) {

	if !I.checkIfString(value1) && !I.checkIfString(value2) {
		return false, errors.New("data type mut be a string")
	}

	isEmpty, err:= I.isStringEmpty(value2.(string))

	if isEmpty {
		return false, err
	}

	resp := strings.EqualFold(value1.(string), value2.(string))

	return resp, nil
}

// Check if the values are equal int
func (I If) IsEqualsToNumber(value1, value2 int) (bool, error) {

	if value1 == value2 {
		return true, nil
	}

	return false, errors.New("an error occurred while checking equality of a number")
}

// Check if the value2 is contained in value1
func (I If) IsStringContain(value1 string, value2 string) (bool, error) {

	if !I.checkIfString(value1) && !I.checkIfString(value2) {
		return false, errors.New("datatype mut be a string")
	}

	isEmpty, err := I.isStringEmpty(value2)

	if isEmpty {
		return false, err
	}

	return strings.Contains(value1, value2), nil
}

// Check if the value2 is not contained in value1
func (I If) IsStringNotContain(value1, value2 string) (bool, error) {


	if !I.checkIfString(value1) && !I.checkIfString(value2) {
		return false, errors.New("datatype mut be a string")
	}

	isEmpty, err := I.isStringEmpty(value2)

	if isEmpty {
		return false, err
	}

	return !strings.Contains(value1, value2), nil
}

// Check if the value2 ends with value1
func (I If) IsStringEndWith(value1, value2 string) (bool, error) {


	if !I.checkIfString(value1) && !I.checkIfString(value2) {
		return false, errors.New("datatype mut be a string")
	}

	isEmpty, err := I.isStringEmpty(value2)

	if isEmpty {
		return false, err
	}

	return strings.HasSuffix(value1, value2), nil
}

// Check if the value2 starts with value1
func (I If) IsStringStartWith(value1, value2 string) (bool, error) {

	if !I.checkIfString(value1) && !I.checkIfString(value2) {
		return false, errors.New("datatype mut be a string")
	}

	isEmpty, err := I.isStringEmpty(value2)

	if isEmpty {
		return false, err
	}

	return strings.HasPrefix(value1, value2), nil
}

// Check if the value2 is lager with value1
func (I If) IsLargerNumber(value1, value2 int)  (bool, error) {

	if !I.checkIfInt(value1) && !I.checkIfInt(value2) {
		return false, errors.New("datatype must be an int")
	}

	return value1 > value2, nil
}

// Check if the value2 is lesser with value1
func (I If) IsLesserNumber(value1, value2 int)  (bool, error) {

	if !I.checkIfInt(value1) && !I.checkIfInt(value2) {
		return false, errors.New("datatype must be an int")
	}

	return value1 < value2, nil
}

// Check if the value2 is lesser/equal with value1
func (I If) IsLesserEqualNumber(value1, value2 int)  (bool, error) {

	if !I.checkIfInt(value1) && !I.checkIfInt(value2) {
		return false, errors.New("datatype must be an int")
	}

	return value1 <= value2, nil
}

// Check if the value2 is lager/equal with value1
func (I If) IsLargerEqualNumber(value1, value2 int)  (bool, error) {

	if !I.checkIfInt(value1) && !I.checkIfInt(value2) {
		return false, errors.New("datatype must be an int")
	}

	return value1 >= value2, nil
}

// Check if time of value2 is equal to value1
func (I If) IsTimeEqual(value1, value2 any) (bool, error) {
	time1, err1 := time.Parse(time.RFC3339, value1.(string))
	time2, err2 := time.Parse(time.RFC3339, value2.(string))

	if err1 != nil {
		return false, err1
	}

	if err2 != nil {
		return false, err2
	}

	return time1.Year() == time2.Year() && time1.YearDay() == time2.YearDay(),  nil
}

func (I If) IsBetweenValue(value1, value2 any) (bool, error) {

	value := strings.Split(value1.(string), ",")

	if I.checkIfDateTime(value1) && I.checkIfDateTime(value2) {
		time1, err1 := time.Parse(time.RFC3339, value[0])
		time2, err2 := time.Parse(time.RFC3339, value[1])
		time3, err3 := time.Parse(time.RFC3339, value2.(string))

		if err1 != nil {
			return false, err1
		}
	
		if err2 != nil {
			return false, err2
		}

		if err3 != nil {
			return false, err3
		}
		
		return I.TimeIsBetween(time3, time1, time2), nil
	}

	if I.checkIfInt(value1) && I.checkIfInt(value2) {
		return value1.(int) >= value2.(int) && value1.(int) <= value2.(int), nil
	}

	return false, errors.New("invalid values for between, supports only int and date")

}

func (I If) TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}