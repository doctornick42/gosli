package experiment

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FakeTypeTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestFactory(t *testing.T) {
	suite.Run(t, new(FakeTypeTestSuite))
}

func (ts *FakeTypeTestSuite) TestFakeTypeFirstOrDefault() {
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
			res := FakeTypeSlice().FirstOrDefault(tc.sl, tc.filter)

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

func (ts *FakeTypeTestSuite) TestFakeTypeFirst() {
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
			res, err := FakeTypeSlice().First(tc.sl, tc.filter)

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

func (ts *FakeTypeTestSuite) TestFakeTypeWhere() {
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
			res := FakeTypeSlice().Where(tc.sl, tc.filter)

			assert.EqualValues(t, tc.expectedRes, res)
		})
	}
}

func (ts *FakeTypeTestSuite) TestFakeTypeSelect() {
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
			res := FakeTypeSlice().Select(tc.sl, tc.filter)

			assert.EqualValues(t, tc.expectedRes, res)
		})
	}
}

func (ts *FakeTypeTestSuite) initDependencies() {
	ts.mockCtrl = gomock.NewController(ts.T())
}
