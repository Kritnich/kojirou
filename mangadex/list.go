package mangadex

import "sort"

type ChapterList []Chapter

type PathList []Path

type ImageList []Image

func (m ChapterList) CollapseBy(f func(ChapterInfo) interface{}) ChapterList {
	keys := make([]interface{}, 0)
	mapped := make(map[interface{}]Chapter)
	for _, val := range m {
		key := f(val.Info)
		if _, ok := mapped[key]; !ok {
			mapped[key] = val
			keys = append(keys, key)
		}
	}

	sorted := make(ChapterList, 0)
	for _, key := range keys {
		sorted = append(sorted, mapped[key])
	}

	return sorted
}

func (m ChapterList) FilterBy(f func(ChapterInfo) bool) ChapterList {
	sorted := make(ChapterList, 0)
	for _, val := range m {
		if f(val.Info) {
			sorted = append(sorted, val)
		}
	}

	return sorted
}

func (m ChapterList) SortBy(f func(ChapterInfo, ChapterInfo) bool) ChapterList {
	sorted := m
	sort.SliceStable(sorted, func(i, j int) bool {
		return f(sorted[i].Info, sorted[j].Info)
	})

	return sorted
}

func (m PathList) FilterBy(f func(Path) bool) PathList {
	sorted := make([]Path, 0)
	for _, val := range m {
		if f(val) {
			sorted = append(sorted, val)
		}
	}

	return sorted
}
