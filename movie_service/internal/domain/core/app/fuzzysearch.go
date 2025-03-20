package app

import (
	"cmp"
	"movie_service/internal/domain/core/models"
	"slices"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/samber/lo"
)

type fuzzySearch struct {
}

type wrapper struct {
	ms   *models.MovieSearch
	rank int
}

func (f *fuzzySearch) Search(moviesSearchList []*models.MovieSearch, search string) []*models.MovieSearch {
	if len(moviesSearchList) == 0 {
		return []*models.MovieSearch{}
	}
	var rankedResult []*wrapper
	for _, ms := range moviesSearchList {
		rank := fuzzy.RankMatchFold(search, ms.Title)
		if rank == -1 {
			continue
		}
		rankedResult = append(rankedResult, &wrapper{
			ms:   ms,
			rank: rank,
		})
	}
	slices.SortFunc(rankedResult, func(a, b *wrapper) int {
		return cmp.Compare(a.rank, b.rank)
	})
	return lo.Map(rankedResult, func(w *wrapper, _ int) *models.MovieSearch {
		return w.ms
	})
}
