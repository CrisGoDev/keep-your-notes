package meta

import (
	"reflect"
	"testing"
)

func TestMeta(t *testing.T) {

	testCases := []struct {
		Name     string
		page     int
		perPage  int
		total    int
		Expected *Meta
	}{
		{
			Name:    "Get Pagination correctly",
			page:    3,
			perPage: 10,
			total:   21,
			Expected: &Meta{
				TotalCount: 21,
				PerPage:    10,
				PageCount:  3,
				Page:       3,
			},
		},

		{
			Name:    "Incorrect Last Page",
			page:    4,
			perPage: 10,
			total:   21,
			Expected: &Meta{
				TotalCount: 21,
				PerPage:    10,
				PageCount:  3,
				Page:       3,
			},
		},

		{
			Name:    "Incorrect Negative Page",
			page:    -1,
			perPage: 10,
			total:   21,
			Expected: &Meta{
				TotalCount: 21,
				PerPage:    10,
				PageCount:  3,
				Page:       1,
			},
		},

		{
			Name:    "Total Page is Zero",
			page:    -1,
			perPage: 10,
			total:   0,
			Expected: &Meta{
				TotalCount: 0,
				PerPage:    10,
				PageCount:  0,
				Page:       1,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {
			result, _ := New(tc.page, tc.perPage, tc.total)

			if !reflect.DeepEqual(result, tc.Expected) {
				t.Errorf("The result was incorrect, got :%d, want :%d", result, tc.Expected)
			}
		})
	}

}
