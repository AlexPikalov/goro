package files

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetRegExp(t *testing.T) {
	//
	// * pattern
	//

	pattern := "abc/*"
	testPath1 := "abc"
	testPath2 := "abc/slkdfjslkdfj"
	testPath3 := "abc/slkdfjslkdfj/xxx"
	testPath4 := "abc/slkdfjslkdfj/xxx/"
	testPath5 := "abcd/slkdfjslkdfj"

	r := getRegExp(pattern)

	if m, err := regexp.MatchString(r, testPath1); err != nil {
		t.Error(err.Error())
	} else if m {
		t.Error("should not match")
	}

	if m, err := regexp.MatchString(r, testPath2); err != nil {
		t.Error(err.Error())
	} else if !m {
		t.Error("should match")
	}

	if m, err := regexp.MatchString(r, testPath3); err != nil {
		t.Error(err.Error())
	} else if m {
		t.Error("should not match")
	}

	if m, err := regexp.MatchString(r, testPath4); err != nil {
		t.Error(err.Error())
	} else if m {
		t.Error("should not match")
	}

	if m, err := regexp.MatchString(r, testPath5); err != nil {
		t.Error(err.Error())
	} else if m {
		t.Error("should not match")
	}

	//
	// ** pattern
	//

	pattern2 := "abc/**"
	testPath21 := "abc"
	testPath22 := "abc/slkdfjslkdfj"
	testPath23 := "abc/slkdfjslkdfj/xxx"
	testPath24 := "abc/slkdfjslkdfj/xxx/"

	r2 := getRegExp(pattern2)
	fmt.Println(r2)

	if m, err := regexp.MatchString(r2, testPath21); err != nil {
		t.Error(err.Error())
	} else if m {
		t.Error("should not match")
	}

	if m, err := regexp.MatchString(r2, testPath22); err != nil {
		t.Error(err.Error())
	} else if !m {
		t.Error("should match")
	}

	if m, err := regexp.MatchString(r2, testPath23); err != nil {
		t.Error(err.Error())
	} else if !m {
		t.Error("should match")
	}

	if m, err := regexp.MatchString(r2, testPath24); err != nil {
		t.Error(err.Error())
	} else if !m {
		t.Error("should match")
	}

	pattern3 := "!abc/**"
}
