package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Int64TestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestInt64Factory(t *testing.T) {
	suite.Run(t, new(Int64TestSuite))
}

func (ts *Int64TestSuite) TestFirstOrDefault() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl          []int64
		filter      func(int64) bool
		expectedRes int64
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 2
			},
			expectedRes: 2,
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 123
			},
			expectedRes: 0,
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := Int64Slice(tc.sl).FirstOrDefault(tc.filter)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func (ts *Int64TestSuite) TestFirst() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl          []int64
		filter      func(int64) bool
		expectedRes int64
		expectedErr error
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 2
			},
			expectedRes: 2,
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 123
			},
			expectedErr: errors.New("Not found"),
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := Int64Slice(tc.sl).First(tc.filter)

			if tc.expectedErr == nil {
				assert.Equal(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *Int64TestSuite) TestWhere() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl          []int64
		filter      func(int64) bool
		expectedRes []int64
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(f int64) bool {
				return f > 1
			},
			expectedRes: []int64{2, 3},
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 123
			},
			expectedRes: []int64{},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := Int64Slice(tc.sl).Where(tc.filter)
			assert.EqualValues(t, tc.expectedRes, res)
		})
	}
}

func (ts *Int64TestSuite) TestSelect() {
	sl := []int64{1, 2, 3}

	type tempType struct {
		Msg string
	}

	testCases := []struct {
		name        string
		sl          []int64
		filter      func(int64) interface{}
		expectedRes []interface{}
	}{
		{
			name: "everything ok",
			sl:   sl,
			filter: func(f int64) interface{} {
				return &tempType{
					Msg: fmt.Sprintf("Value: %v", f),
				}
			},
			expectedRes: []interface{}{
				&tempType{
					Msg: "Value: 1",
				},
				&tempType{
					Msg: "Value: 2",
				},
				&tempType{
					Msg: "Value: 3",
				},
			},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := Int64Slice(tc.sl).Select(tc.filter)
			assert.EqualValues(t, tc.expectedRes, res)
		})
	}
}

func (ts *Int64TestSuite) TestPage() {
	sl := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testCases := []struct {
		name        string
		sl          []int64
		pageNumber  int64
		perPage     int64
		expectedRes []int64
		expectedErr error
	}{
		{
			name:        "10 items, per page - 5, page 1",
			sl:          sl,
			pageNumber:  1,
			perPage:     5,
			expectedRes: []int64{1, 2, 3, 4, 5},
		},
		{
			name:        "10 items, per page - 5, page 2",
			sl:          sl,
			pageNumber:  2,
			perPage:     5,
			expectedRes: []int64{6, 7, 8, 9, 10},
		},
		{
			name:        "10 items, per page - 5, page 3",
			sl:          sl,
			pageNumber:  3,
			perPage:     5,
			expectedRes: []int64{},
		},
		{
			name:        "10 items, per page - 7, page 2",
			sl:          sl,
			pageNumber:  2,
			perPage:     7,
			expectedRes: []int64{8, 9, 10},
		},
		{
			name:        "10 items, per page - 12, page 1",
			sl:          sl,
			pageNumber:  1,
			perPage:     12,
			expectedRes: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
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
			expectedRes: []int64{},
		},
		{
			name:        "10 items, per page - 5, page 15",
			sl:          sl,
			pageNumber:  15,
			perPage:     5,
			expectedRes: []int64{},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := Int64Slice(tc.sl).Page(tc.pageNumber, tc.perPage)

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

func (ts *Int64TestSuite) TestAny() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl          []int64
		filter      func(int64) bool
		expectedRes bool
	}{
		{
			name: "found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 2
			},
			expectedRes: true,
		},
		{
			name: "not found",
			sl:   sl,
			filter: func(f int64) bool {
				return f == 123
			},
			expectedRes: false,
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res := Int64Slice(tc.sl).Any(tc.filter)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func (ts *Int64TestSuite) TestContains() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl          []int64
		el          int64
		expectedRes bool
		expectedErr error
	}{
		{
			name:        "contains",
			sl:          sl,
			el:          2,
			expectedRes: true,
		},
		{
			name:        "doesn't contain",
			sl:          sl,
			el:          7000,
			expectedRes: false,
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := Int64Slice(tc.sl).Contains(tc.el)

			if tc.expectedErr == nil {
				assert.Equal(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *Int64TestSuite) TestGetUnion() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl1         []int64
		sl2         []int64
		expectedRes []int64
		expectedErr error
	}{
		{
			name:        "union exists",
			sl1:         sl,
			sl2:         []int64{2, 3, 4},
			expectedRes: []int64{2, 3},
		},
		{
			name:        "no union",
			sl1:         sl,
			sl2:         []int64{5, 4},
			expectedRes: []int64{},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := Int64Slice(tc.sl1).GetUnion(tc.sl2)

			if tc.expectedErr == nil {
				assert.EqualValues(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *Int64TestSuite) TestInFirstOnly() {
	sl := []int64{1, 2, 3}

	testCases := []struct {
		name        string
		sl1         []int64
		sl2         []int64
		expectedRes []int64
		expectedErr error
	}{
		{
			name:        "union exists",
			sl1:         sl,
			sl2:         []int64{2, 3, 4},
			expectedRes: []int64{1},
		},
		{
			name:        "no union",
			sl1:         sl,
			sl2:         []int64{5, 4},
			expectedRes: []int64{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		ts.initDependencies()

		ts.T().Run(tc.name, func(t *testing.T) {
			res, err := Int64Slice(tc.sl1).InFirstOnly(tc.sl2)

			if tc.expectedErr == nil {
				assert.EqualValues(t, tc.expectedRes, res)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			}
		})
	}
}

func (ts *Int64TestSuite) initDependencies() {
	ts.mockCtrl = gomock.NewController(ts.T())
}
