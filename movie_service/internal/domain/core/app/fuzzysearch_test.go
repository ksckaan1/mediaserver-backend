package app

import (
	"movie_service/internal/domain/core/models"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuzzySearch_Search(t *testing.T) {
	tests := []struct {
		search     string
		givenList  []*models.MovieSearch
		wantedList []*models.MovieSearch
	}{
		{
			search: "john wick",
			givenList: []*models.MovieSearch{
				{
					Title: "John Wick 4",
				},
				{
					Title: "Transporter 4",
				},
				{
					Title: "The Matrix 4",
				},
			},
			wantedList: []*models.MovieSearch{
				{
					Title: "John Wick 4",
				},
			},
		},
		{
			search: "ted",
			givenList: []*models.MovieSearch{
				{
					Title: "Teddy",
				},
				{
					Title: "Ted",
				},
				{
					Title: "Terry Diamond",
				},
			},
			wantedList: []*models.MovieSearch{
				{
					Title: "Ted",
				},
				{
					Title: "Teddy",
				},
				{
					Title: "Terry Diamond",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			gotList := (&fuzzySearch{}).Search(tt.givenList, tt.search)
			require.Equal(t, tt.wantedList, gotList)
		})
	}
}

func TestSelam(t *testing.T) {
	selam()
}

func selam() {
	pc, _, _, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	println(functionName)
}
