package experiment

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FakeTypePTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func (ts *FakeTypePTestSuite) TestFakeTypeFirstOrDefault() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		filter      func(*FakeType) bool
		expectedRes *FakeType
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 2
			},
			expectedRes: &FakeType{
				A: 2,
				B: "two",
			},
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 123
			},
			expectedRes: nil,
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := FakeTypePSlice(tc.sl).FirstOrDefault(tc.filter)

			if tc.expectedRes == nil {
				assert.Nil(t, res)
				return
			}

			isEqualToExpected, err := res.Equal(tc.expectedRes)
			assert.Nil(t, err)
			assert.True(t, isEqualToExpected)
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeFirst() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		filter      func(*FakeType) bool
		expectedRes *FakeType
		expectedErr error
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 2
			},
			expectedRes: &FakeType{
				A: 2,
				B: "two",
			},
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 123
			},
			expectedErr: errors.New("Not found"),
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := FakeTypePSlice(tc.sl).First(tc.filter)

			if tc.expectedErr == nil {
				isEqualToExpected, err := res.Equal(tc.expectedRes)
				assert.Nil(t, err)
				assert.True(t, isEqualToExpected)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeWhere() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		filter      func(*FakeType) bool
		expectedRes []*FakeType
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A > 1
			},
			expectedRes: []*FakeType{
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
			},
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 123
			},
			expectedRes: []*FakeType{},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := FakeTypePSlice(tc.sl).Where(tc.filter)

			assert.EqualValues(t, tc.expectedRes, res)
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeSelect() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	type tempType struct {
		Msg string
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		filter      func(*FakeType) interface{}
		expectedRes []interface{}
	}{
		{
			name: "everything ok",
			sl:   sl,
			filter: func(ft *FakeType) interface{} {
				return &tempType{
					Msg: fmt.Sprintf("%v-%s", ft.A, ft.B),
				}
			},
			expectedRes: []interface{}{
				&tempType{
					Msg: "1-one",
				},
				&tempType{
					Msg: "2-two",
				},
				&tempType{
					Msg: "3-three",
				},
			},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := FakeTypePSlice(tc.sl).Select(tc.filter)

			assert.EqualValues(t, tc.expectedRes, res)
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypePage() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
		&FakeType{
			A: 4,
			B: "four",
		},
		&FakeType{
			A: 5,
			B: "five",
		},
		&FakeType{
			A: 6,
			B: "six",
		},
		&FakeType{
			A: 7,
			B: "seven",
		},
		&FakeType{
			A: 8,
			B: "eight",
		},
		&FakeType{
			A: 9,
			B: "nine",
		},
		&FakeType{
			A: 10,
			B: "ten",
		},
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		pageNumber  int64
		perPage     int64
		expectedRes []*FakeType
		expectedErr error
	}{
		{
			name:       "10 items, per page - 5, page 1",
			sl:         sl,
			pageNumber: 1,
			perPage:    5,
			expectedRes: []*FakeType{
				&FakeType{
					A: 1,
					B: "one",
				},
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
				&FakeType{
					A: 4,
					B: "four",
				},
				&FakeType{
					A: 5,
					B: "five",
				},
			},
		},
		{
			name:       "10 items, per page - 5, page 2",
			sl:         sl,
			pageNumber: 2,
			perPage:    5,
			expectedRes: []*FakeType{
				&FakeType{
					A: 6,
					B: "six",
				},
				&FakeType{
					A: 7,
					B: "seven",
				},
				&FakeType{
					A: 8,
					B: "eight",
				},
				&FakeType{
					A: 9,
					B: "nine",
				},
				&FakeType{
					A: 10,
					B: "ten",
				},
			},
		},
		{
			name:        "10 items, per page - 5, page 3",
			sl:          sl,
			pageNumber:  3,
			perPage:     5,
			expectedRes: []*FakeType{},
		},
		{
			name:       "10 items, per page - 7, page 2",
			sl:         sl,
			pageNumber: 2,
			perPage:    7,
			expectedRes: []*FakeType{
				&FakeType{
					A: 8,
					B: "eight",
				},
				&FakeType{
					A: 9,
					B: "nine",
				},
				&FakeType{
					A: 10,
					B: "ten",
				},
			},
		},
		{
			name:       "10 items, per page - 12, page 1",
			sl:         sl,
			pageNumber: 1,
			perPage:    12,
			expectedRes: []*FakeType{
				&FakeType{
					A: 1,
					B: "one",
				},
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
				&FakeType{
					A: 4,
					B: "four",
				},
				&FakeType{
					A: 5,
					B: "five",
				},
				&FakeType{
					A: 6,
					B: "six",
				},
				&FakeType{
					A: 7,
					B: "seven",
				},
				&FakeType{
					A: 8,
					B: "eight",
				},
				&FakeType{
					A: 9,
					B: "nine",
				},
				&FakeType{
					A: 10,
					B: "ten",
				},
			},
		},
		{
			name:        "10 items, per page - 5, page 0",
			sl:          sl,
			pageNumber:  0,
			perPage:     5,
			expectedErr: errors.New("Page number should start with 1"),
		},
		{
			name:        "10 items, per page - 5, page -1",
			sl:          sl,
			pageNumber:  -1,
			perPage:     5,
			expectedErr: errors.New("Page number should start with 1"),
		},
		{
			name:        "10 items, per page - 0, page 1",
			sl:          sl,
			pageNumber:  1,
			perPage:     0,
			expectedRes: []*FakeType{},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := FakeTypePSlice(tc.sl).Page(tc.pageNumber, tc.perPage)

			if tc.expectedErr == nil {
				assert.Nil(t, err)
				assert.EqualValues(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeAny() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		filter      func(*FakeType) bool
		expectedRes bool
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 2
			},
			expectedRes: true,
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(ft *FakeType) bool {
				return ft.A == 123
			},
			expectedRes: false,
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := FakeTypePSlice(tc.sl).Any(tc.filter)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeContains() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl          []*FakeType
		el          *FakeType
		expectedRes bool
		expectedErr error
	}{
		{
			name: "contains",
			sl:   sl,
			el: &FakeType{
				A: 2,
				B: "two",
			},
			expectedRes: true,
		},
		{
			name: "doesn't contain",
			sl:   sl,
			el: &FakeType{
				A: 7000,
				B: "seven thousands",
			},
			expectedRes: false,
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := FakeTypePSlice(tc.sl).Contains(tc.el)

			if tc.expectedErr == nil {
				assert.Equal(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeGetUnion() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl1         []*FakeType
		sl2         []*FakeType
		expectedRes []*FakeType
		expectedErr error
	}{
		{
			name: "union exists",
			sl1:  sl,
			sl2: []*FakeType{
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
				&FakeType{
					A: 4,
					B: "four",
				},
			},
			expectedRes: []*FakeType{
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
			},
		},
		{
			name: "no union",
			sl1:  sl,
			sl2: []*FakeType{
				&FakeType{
					A: 5,
					B: "five",
				},
				&FakeType{
					A: 4,
					B: "four",
				},
			},
			expectedRes: []*FakeType{},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := FakeTypePSlice(tc.sl1).GetUnion(tc.sl2)

			if tc.expectedErr == nil {
				assert.EqualValues(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *FakeTypePTestSuite) TestFakeTypeInFirstOnly() {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	testCases := []struct {
		name        string
		sl1         []*FakeType
		sl2         []*FakeType
		expectedRes []*FakeType
		expectedErr error
	}{
		{
			name: "union exists",
			sl1:  sl,
			sl2: []*FakeType{
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
				&FakeType{
					A: 4,
					B: "four",
				},
			},
			expectedRes: []*FakeType{
				&FakeType{
					A: 1,
					B: "one",
				},
			},
		},
		{
			name: "no union",
			sl1:  sl,
			sl2: []*FakeType{
				&FakeType{
					A: 5,
					B: "five",
				},
				&FakeType{
					A: 4,
					B: "four",
				},
			},
			expectedRes: []*FakeType{
				&FakeType{
					A: 1,
					B: "one",
				},
				&FakeType{
					A: 2,
					B: "two",
				},
				&FakeType{
					A: 3,
					B: "three",
				},
			},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := FakeTypePSlice(tc.sl1).InFirstOnly(tc.sl2)

			if tc.expectedErr == nil {
				assert.EqualValues(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *FakeTypePTestSuite) initDependencies() {
	ts.mockCtrl = gomock.NewController(ts.T())
}
