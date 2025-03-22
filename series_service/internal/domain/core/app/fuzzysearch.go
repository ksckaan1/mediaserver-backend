package app

import (
	"cmp"
	"series_service/internal/domain/core/models"
	"slices"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/samber/lo"
)

type fuzzySearch struct {
}

type wrapper struct {
	ms   *models.SeriesSearch
	rank int
}

func (f *fuzzySearch) Search(seriesSearchList []*models.SeriesSearch, search string) []*models.SeriesSearch {
	if len(seriesSearchList) == 0 {
		return []*models.SeriesSearch{}
	}
	var rankedResult []*wrapper
	for _, ms := range seriesSearchList {
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
	return lo.Map(rankedResult, func(w *wrapper, _ int) *models.SeriesSearch {
		return w.ms
	})
}
