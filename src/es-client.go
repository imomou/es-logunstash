package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type esClient struct {
	Host string
}

func (esc *esClient) GetAllocations() ([]allocation, error) {
	resp, err := http.Get(esc.Host + "/_cat/allocation?format=json&bytes=b")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	allocationsRaw := make([]allocationRaw, 0)
	json.Unmarshal(body, &allocationsRaw)

	allocations := make([]allocation, len(allocationsRaw))
	for i, allocationRaw := range allocationsRaw {
		allocations[i] = *parseAllocation(&allocationRaw)
	}

	return allocations, nil
}
func (esc *esClient) GetIndices() ([]index, error) {
	resp, err := http.Get(esc.Host + "/_cat/indices?format=json&bytes=b")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	indicesRaw := make([]indexRaw, 0)
	json.Unmarshal(body, &indicesRaw)

	indices := make([]index, len(indicesRaw))
	for i, indexRaw := range indicesRaw {
		indices[i] = *parseIndex(&indexRaw)
	}

	return indices, nil
}
func (esc *esClient) DeleteIndex(indexName string) error {
	fmt.Println("Index to delete : ", indexName)
	//return nil
	req, err := http.NewRequest("DELETE", esc.Host+"/"+indexName, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Fetch Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
	return nil
}

type allocationRaw struct {
	Shards      string `json:"shards"`
	DiskIndices string `json:"disk.indices"`
	DiskUsed    string `json:"disk.used"`
	DiskAvail   string `json:"disk.avail"`
	DiskTotal   string `json:"disk.total"`
	DiskPercent string `json:"disk.percent"`
	Host        string `json:"host"`
	IP          string `json:"ip"`
	Node        string `json:"node"`
}
type allocation struct {
	Shards      int    `json:"shards"`
	DiskIndices int    `json:"disk.indices"`
	DiskUsed    int64  `json:"disk.used"`
	DiskAvail   int64  `json:"disk.avail"`
	DiskTotal   int64  `json:"disk.total"`
	DiskPercent int    `json:"disk.percent"`
	Host        string `json:"host"`
	IP          string `json:"ip"`
	Node        string `json:"node"`
}

func parseAllocation(ar *allocationRaw) *allocation {
	shards, _ := strconv.Atoi(ar.Shards)
	diskIndices, _ := strconv.Atoi(ar.DiskIndices)
	diskPercent, _ := strconv.Atoi(ar.DiskPercent)
	diskUsed, _ := strconv.ParseInt(ar.DiskUsed, 10, 64)
	diskAvail, _ := strconv.ParseInt(ar.DiskAvail, 10, 64)
	diskTotal, _ := strconv.ParseInt(ar.DiskTotal, 10, 64)

	return &allocation{
		Shards:      shards,
		DiskIndices: diskIndices,
		DiskUsed:    diskUsed,
		DiskAvail:   diskAvail,
		DiskTotal:   diskTotal,
		DiskPercent: diskPercent,
		Host:        ar.Host,
		IP:          ar.IP,
		Node:        ar.Node,
	}
}

type indexRaw struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	UUID         string `json:"uuid"`
	Pri          string `json:"pri"`
	Rep          string `json:"rep"`
	DocsCount    string `json:"docs.count"`
	DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PriStoreSize string `json:"pri.store.size"`
}

type index struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	UUID         string `json:"uuid"`
	Pri          int    `json:"pri"`
	Rep          int    `json:"rep"`
	DocsCount    int64  `json:"docs.count"`
	DocsDeleted  int64  `json:"docs.deleted"`
	StoreSize    int64  `json:"store.size"`
	PriStoreSize int64  `json:"pri.store.size"`
}

func parseIndex(ir *indexRaw) *index {
	pri, _ := strconv.Atoi(ir.Pri)
	rep, _ := strconv.Atoi(ir.Rep)
	docsCount, _ := strconv.ParseInt(ir.DocsCount, 10, 64)
	docsDeleted, _ := strconv.ParseInt(ir.DocsDeleted, 10, 64)
	storeSize, _ := strconv.ParseInt(ir.StoreSize, 10, 64)
	priStoreSize, _ := strconv.ParseInt(ir.PriStoreSize, 10, 64)

	return &index{
		Health:       ir.Health,
		Status:       ir.Status,
		Index:        ir.Index,
		UUID:         ir.UUID,
		Pri:          pri,
		Rep:          rep,
		DocsCount:    docsCount,
		DocsDeleted:  docsDeleted,
		StoreSize:    storeSize,
		PriStoreSize: priStoreSize,
	}
}
