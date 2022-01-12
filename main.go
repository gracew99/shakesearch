package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"math"
	"strconv"
	"strings"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./client/build"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))
	http.HandleFunc("/read", handleRead(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

type searchResult struct {
	LowerBound   int
	UpperBound int
	Index   int
	Query string
}


type mergedSearchResult struct {
	LowerBound   int
	UpperBound int
	Index   []int
	Query []string
}

type finalSearchResult struct {
	Result string
	Index[] int
	Query[] string
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		fmt.Println("SEARCHING")
		initialResult := searcher.Search(strings.ToLower(query[0]))
		// fmt.Println(len(strings.Fields(query[0])))
		additionalResults := []searchResult{}
		if len(strings.Fields(query[0])) > 1 {
			for i := 0; i < len(strings.Fields(query[0])); i++ {
				fmt.Println(strings.ToLower(strings.Fields(string(query[0]))[i]))
				newResult := searcher.Search(strings.ToLower(strings.Fields(string(query[0]))[i]))
				additionalResults = append(additionalResults, newResult...)
	
			}
		}
		// fmt.Println(additionalResults[0:2])
		sort.Slice(additionalResults, func(i, j int) bool {
			return additionalResults[i].Index < additionalResults[j].Index
		})
		results := append(initialResult, additionalResults...)
		fmt.Println(len(results))

		finalResults := []finalSearchResult{}
		if len(results) > 0 {

			sort.Slice(results, func(i, j int) bool {
				if results[i].Index == results[j].Index {
					return len(results[i].Query) > len(results[j].Query)
				}
				return results[i].Index < results[j].Index
			})
	
			// merge overlap
			mergedResults := []mergedSearchResult{}
			// special case: first item
			var firstIndexArr []int
			firstIndexArr = append(firstIndexArr, results[0].Index)
			var firstQueryArr []string
			firstQueryArr = append(firstQueryArr, results[0].Query)
			firstMergedResult := mergedSearchResult{
				LowerBound: results[0].LowerBound,
				UpperBound: results[0].UpperBound,
				Index: firstIndexArr,
				Query: firstQueryArr,
				
			}
	
			fmt.Println("Results")
			// fmt.Println(results)
	
			mergedResults = append(mergedResults, firstMergedResult)
			// fmt.Println(mergedResults)
			// current := mergedSearchResult {}
			for i := 1; i < len(results); i++ {
				anchor := mergedResults[len(mergedResults)-1]
				compare := results[i]
				// fmt.Println("Anchor then compare")
				// fmt.Printf("%+v\n", anchor)
				// fmt.Printf("%+v\n", compare)
	
				// overlap with window
				if compare.Index <= anchor.UpperBound {
					// extend window
					anchor.UpperBound = compare.UpperBound
					// overlap with highlight: discard
					anchorIndex := anchor.Index[len(anchor.Index)-1]
					anchorQuery := anchor.Query[len(anchor.Query)-1]
					if compare.Index >= anchorIndex && compare.Index <= anchorIndex + len(anchorQuery) {
						fmt.Println("Swallow")
						continue
					} else if compare.Index > anchorIndex + len(anchorQuery) { 				// no overlap: append index/query
						// fmt.Println("append to same sublist")
						// fmt.Println(mergedResults)
						anchor.Index = append(anchor.Index, compare.Index)
						anchor.Query = append(anchor.Query, compare.Query)
						mergedResults[len(mergedResults)-1] = anchor
						// fmt.Println(mergedResults)
	
					} else {
						fmt.Println(compare.Index)
						fmt.Println(anchorIndex)
						fmt.Println(len(anchorQuery))
						fmt.Println("SHOULD NOT REACH HERE")
					}
				} else {
					// fmt.Println("append new list")
					var indexArr []int
					indexArr = append(indexArr, results[i].Index)
					var queryArr []string
					queryArr = append(queryArr, results[i].Query)
	
					mergedResult := mergedSearchResult{
						LowerBound: results[i].LowerBound,
						UpperBound: results[i].UpperBound,
						Index: indexArr,
						Query: queryArr,
						
					}
					mergedResults = append(mergedResults, mergedResult)
				}
			}
	
			// fmt.Printf("%+v\n", mergedResults)
			// fmt.Println("NOW")
			// iterate through mergedResults and highlight
	
			for i := 0; i < len(mergedResults); i++ {
				// fmt.Println("loop")
				result := searcher.CompleteWorks[mergedResults[i].LowerBound: mergedResults[i].Index[0]]
				j := 0
				// fmt.Println(strings.Fields(result))
				// fmt.Println("inner loop")
				for j = 0; j < len(mergedResults[i].Index)-1; j++ {
					result += "<mark>"
					result += searcher.CompleteWorks[mergedResults[i].Index[j]: mergedResults[i].Index[j] + len(mergedResults[i].Query[j])]
					result += "</mark>"
					result += searcher.CompleteWorks[mergedResults[i].Index[j] + len(mergedResults[i].Query[j]): mergedResults[i].Index[j+1]]
	
					// fmt.Println("temp")
					// fmt.Println(strings.Fields(result))
				}
				// fmt.Println(j)
				result += "<mark>"
				result += searcher.CompleteWorks[mergedResults[i].Index[j]: mergedResults[i].Index[j] + len(mergedResults[i].Query[j])]
				result += "</mark>"
				if mergedResults[i].Index[j] + len(mergedResults[i].Query[j]) < mergedResults[i].UpperBound {
					// fmt.Println(mergedResults[i].Index[j] + len(mergedResults[i].Query[j]))
					// fmt.Println(mergedResults[i].UpperBound)
					result += searcher.CompleteWorks[mergedResults[i].Index[j] + len(mergedResults[i].Query[j]): mergedResults[i].UpperBound]
				}
				// fmt.Println(strings.Fields(result))
				finalResult := finalSearchResult {
					Result: result,
					Index: mergedResults[i].Index,
					Query: mergedResults[i].Query,
				}
				finalResults = append(finalResults, finalResult)
			}
			fmt.Println(len(finalResults))
			// fmt.Println(finalResults)
		}
		
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(finalResults)
		// b, err := json.Marshal(results)
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
    	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
		// if c != nil {
		// 	fmt.Println(c)
		// } else {
		// 	fmt.Println(a)
		// }
	}
}


func handleRead(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		index, ok := r.URL.Query()["index"]
		if !ok || len(index) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing index in URL params"))
			return
		}
		query, ok1 := r.URL.Query()["query"]
		if !ok1 || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing index in URL params"))
			return
		}
		highlight, ok2 := r.URL.Query()["highlight"]
		if !ok2 || len(highlight[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing highlight in URL params"))
			return
		}



		n, _ := strconv.Atoi(index[0])
		highlightBool, _ := strconv.ParseBool(highlight[0])

		results := searcher.Read(n, query[0], highlightBool)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		// user := &User{Name: "Frank"}
		// b, err := json.Marshal(results)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("HUH")
		// fmt.Println(string(b))
		// fmt.Println("HUH")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
    	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		bytes := buf.Bytes()
		w.Write(bytes)
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	lowercaseStr := strings.ToLower(string(dat))
	lowercaseArr := []byte(lowercaseStr)
	s.SuffixArray = suffixarray.New(lowercaseArr)
	return nil
}

func repeatedSlice(value, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = value
	}
	return arr
}

func (s *Searcher) Search(query string) []searchResult {
	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	sort.Ints(idxs)
	results := []searchResult{}
	for _, idx := range idxs {
		lowerBound := int(math.Max(0, float64(idx-150)))
		for i := lowerBound; i > 0; i-- {
			if s.CompleteWorks[i-1] == '\n' {
				lowerBound = i
				break
			}
		  }
		
		upperBound := int(math.Min(float64(len(s.CompleteWorks)-1), float64(idx+150)))
		for i := upperBound; i >= idx; i-- {
			if s.CompleteWorks[i] == '\n' {
				upperBound = i
				break
			}
		}
		
		// boldedResult := s.CompleteWorks[lowerBound: idx] + "<mark>" +  s.CompleteWorks[idx: idx+len(query)] + "</mark>" 
		// if idx+len(query) < len(s.CompleteWorks)-1 {
		// 	boldedResult = boldedResult + s.CompleteWorks[idx+len(query): upperBound]
		// }
		
		finalResult := searchResult{LowerBound: lowerBound, UpperBound: upperBound, Index: idx, Query: query}
		results = append(results, finalResult)
	}
	return results
}

func (s *Searcher) Read(idx int, query string, highlight bool) string {
	lowerBound := int(math.Max(0, float64(idx-500)))
	for i := lowerBound; i > 0; i-- {
		if s.CompleteWorks[i-1] == '\n' {
			lowerBound = i
			break
		}
	  }
	upperBound := int(math.Min(float64(len(s.CompleteWorks)-1), float64(idx+500)))
	for i := upperBound; i >= idx; i-- {
		if s.CompleteWorks[i] == '\n' {
			upperBound = i
			break
		}
	  }
	highlightedResult := s.CompleteWorks[lowerBound: idx] + "<mark>" +  s.CompleteWorks[idx: idx+len(query)] + "</mark>" + s.CompleteWorks[idx+len(query): upperBound]
	if highlight {
		return highlightedResult 
	}
	return s.CompleteWorks[lowerBound: upperBound]
}
