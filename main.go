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

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

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
	Result   string
	Index   int
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
		results := searcher.Search(strings.ToLower(query[0]))
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
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

func (s *Searcher) Search(query string) []searchResult {
	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	sort.Ints(idxs)
	results := []searchResult{}
	for _, idx := range idxs {
		lowerBound := int(math.Max(0, float64(idx-100)))
		for i := lowerBound; i > 0; i-- {
			if s.CompleteWorks[i-1] == '\n' {
				lowerBound = i
				break
			}
		  }
		upperBound := int(math.Min(float64(len(s.CompleteWorks)), float64(idx+100)))
		for i := upperBound; i <= int(len(s.CompleteWorks)); i++ {
			if s.CompleteWorks[i] == '\n' {
				upperBound = i
				break
			}
		  }
		boldedResult := s.CompleteWorks[lowerBound: idx] + "<mark>" +  s.CompleteWorks[idx: idx+len(query)] + "</mark>" + s.CompleteWorks[idx+len(query): upperBound]
		finalResult := searchResult{Result: boldedResult, Index: idx}
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
	upperBound := int(math.Min(float64(len(s.CompleteWorks)), float64(idx+500)))
	for i := upperBound; i <= int(len(s.CompleteWorks)); i++ {
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
