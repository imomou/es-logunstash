package main

import (
	"fmt"
	"sort"
	"strings"
)

type indicesService struct {
	ec esClient
}

func (is *indicesService) ReportByIndexGrouping(prefixes []string) error {
	indices, _ := is.ec.GetIndices()

	totalTotal := int64(0)
	for _, prefix := range prefixes {

		filteredIndices := make([]index, 0)
		for _, index := range indices {
			if strings.HasPrefix(index.Index, prefix) {
				filteredIndices = append(filteredIndices, index)

			}
		}
		sort.Slice(filteredIndices, func(i, j int) bool {
			return filteredIndices[i].Index < filteredIndices[j].Index
		})
		// trimming first and last index to ensure we don't get unfinished indexes.
		totalPriStore := int64(0)
		for i, index := range filteredIndices {
			if i != 0 && i != len(filteredIndices)-1 {
				totalPriStore += index.PriStoreSize
			}
		}
		// for _, index := range filteredIndices {
		// 	fmt.Printf("%v - %v\n", index.Index, index.PriStoreSize/1000/1000)
		// }
		fmt.Printf("%vtotal (count: %v) - %vmb\n", prefix, len(filteredIndices), totalPriStore/1000/1000)
		fmt.Printf("%vaverage - %vmb per day\n\n", prefix, totalPriStore/1000/1000/int64(len(filteredIndices)-2))
		totalTotal += totalPriStore
	}

	fmt.Printf("Total across all prefixes - %v\n\n", totalTotal/1000/1000)

	return nil
}

func (is *indicesService) Purge(prefixToPurge string, numberToPreserve int) error {
	indices, _ := is.ec.GetIndices()

	filteredIndexes := make([]index, 0)
	for _, index := range indices {
		if strings.HasPrefix(index.Index, prefixToPurge) {
			filteredIndexes = append(filteredIndexes, index)
		}
	}

	sort.Slice(filteredIndexes, func(i, j int) bool {
		return filteredIndexes[i].Index < filteredIndexes[j].Index
	})
	filteredIndexLength := len(filteredIndexes)
	for i, index := range filteredIndexes {
		if i+numberToPreserve+1 > filteredIndexLength {
			break
		}

		is.ec.DeleteIndex(index.Index)
	}
	return nil
}
