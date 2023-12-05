package controls

import (
	"testing"
)

var ifControl = If{}

func TestCheckIfString(t *testing.T) {
	ok := ifControl.checkIfString("Hello world!")

	if ok {
		t.Log("valid")
	} else {
		t.Error("Invalid")
	}
}

func TestCheckIfInt(t *testing.T) {
	ok := ifControl.checkIfInt(3)

	if ok {
		t.Log("valid")
	} else {
		t.Error("Invalid")
	}
}

func TestCheckIfDateTime(t *testing.T) {
	ok := ifControl.checkIfDateTime("2023-11-23T10:28:16.980Z")

	if ok {
		t.Log("valid")
	} else {
		t.Error("Invalid")
	}
}

func TestStringEquality(t *testing.T) {
	ok, err := ifControl.IfControl(IsEquals, "Hello", "Hello")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("string not equal")
	} else {
		t.Log("string is equal")
	}
}

func TestIntEquality(t *testing.T) {
	ok, err := ifControl.IfControl(IsEquals, 2, 2)

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("int not equal")
	} else {
		t.Log("int is equal")
	}
}

func TestTimeEquality(t *testing.T) {
	ok, err := ifControl.IfControl(IsEquals, "2023-11-23T10:28:16.980Z", "2023-11-23T10:28:16.980Z")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("time not equal")
	} else {
		t.Log("time is equal")
	}
}

func TestIsStringContain(t *testing.T) {
	ok, err := ifControl.IfControl(IsStringContain, "Hello world", "world")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 does not contain value2")
	} else {
		t.Log("value1 contain value2")
	}
}

func TestIsStringNotContain(t *testing.T) {
	ok, err := ifControl.IfControl(IsStringNotContain, "Hello", "world")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 does not contain value2")
	} else {
		t.Log("value1 contain value2")
	}
}

func TestIsStringEndWith(t *testing.T) {
	ok, err := ifControl.IfControl(IsStringEndWith, "Hello", "o")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 does not end with value2")
	} else {
		t.Log("value1 end with value2")
	}
}

func TestIsStringStartWith(t *testing.T) {
	ok, err := ifControl.IfControl(IsStringStartWith, "Hello", "H")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 does not end with value2")
	} else {
		t.Log("value1 end with value2")
	}
}

func TestIsStringEmpty(t *testing.T) {
	ok, err := ifControl.IfControl(IsStringEmpty, "H", "")

	if err != nil {
		t.Error(err)
	} else if ok {
		t.Error("string empty")
	} else {
		t.Log("string is not empty")
	}
}

func TestIsLargerNumber(t *testing.T) {
	ok, err := ifControl.IfControl(IsLargerNumber, 2, 1)

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 is not larger than value2")
	} else {
		t.Log("value1 is larger value2")
	}
}

func TestIsLesserNumber(t *testing.T) {
	ok, err := ifControl.IfControl(IsLesserNumber, 1, 3)

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 is not lesser than value2")
	} else {
		t.Log("value1 is lesser value2")
	}
}

func TestIsLesserEqualNumber(t *testing.T) {
	ok, err := ifControl.IfControl(IsLesserEqualNumber, 1, 1)

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 is not lesser/equal than value2")
	} else {
		t.Log("value1 is lesser/equal value2")
	}
}

func TestIsLargerEqualNumber(t *testing.T) {
	ok, err := ifControl.IfControl(IsLargerEqualNumber, 1, 1)

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value1 is not lesser/equal than value2")
	} else {
		t.Log("value1 is lesser/equal value2")
	}
}

func TestIsBetweenValue(t *testing.T) {
	ok, err := ifControl.IfControl(IsBetween, "1954-10-03,1960-06-06", "1924-10-03")

	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Error("value2 is not in between value1")
	} else {
		t.Log("value2 is between value2")
	}
}