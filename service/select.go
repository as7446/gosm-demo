package service

import (
	"sort"
	"strings"
	"time"
)

type storageObj struct {
	GenericObjList  []StorageCell
	dataSelectQuery *DataSelectQuery
}
type StorageCell interface {
	GetCreation() time.Time
	GetName(string) string
}

type DataSelectQuery struct {
	FilterQuery   *FilerQuery
	PaginateQuery *PaginateQuery
}

type FilerQuery struct {
	Name string
	Type string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

func (s *storageObj) Len() int {
	return len(s.GenericObjList)
}

func (s *storageObj) Swap(i, j int) {
	s.GenericObjList[i], s.GenericObjList[j] = s.GenericObjList[j], s.GenericObjList[i]
}

func (s *storageObj) Less(i, j int) bool {
	a := s.GenericObjList[i].GetCreation()
	b := s.GenericObjList[j].GetCreation()
	return a.Before(b)
}
func (s *storageObj) Sort() *storageObj {
	sort.Sort(s)
	return s
}

// 过滤
func (s *storageObj) Filter() *storageObj {
	if s.dataSelectQuery.FilterQuery.Name == "" {
		return s
	}
	storageCells := []StorageCell{}
	for _, v := range s.GenericObjList {
		matches := true
		if !strings.Contains(v.GetName(s.dataSelectQuery.FilterQuery.Type), s.dataSelectQuery.FilterQuery.Name) {
			matches = false
			continue
		}
		if matches {
			storageCells = append(storageCells, v)
		}
	}
	s.GenericObjList = storageCells
	return s
}

// 分页
func (s *storageObj) Paginate() *storageObj {
	limit := s.dataSelectQuery.PaginateQuery.Limit
	page := s.dataSelectQuery.PaginateQuery.Page

	startIndex := limit * (page - 1)
	endIndex := limit * page
	if len(s.GenericObjList) < startIndex {
		startIndex = 0
	}
	// 25 10 2
	if len(s.GenericObjList) > endIndex {
		endIndex = len(s.GenericObjList)
	}
	s.GenericObjList = s.GenericObjList[startIndex:endIndex]
	return s
}
