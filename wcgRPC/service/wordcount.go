package service

import (
	"sort"
	"strings"
	"wcgRPC/proto"
)

const limit = 99999

type Word interface {
	GetWords(text string) []*proto.WordCount
}

func WordCount() WordCount {
	return &WordCount{}
}

func (s *wordService) GetWords(text string) []*word_pb.WordCount {

	words := strings.Fields(text)

	maps := make(chan map[string]uint64)
	count := 1

	for {
		if len(words) > limit {
			go countWords(words[:limit], maps)
			words = words[limit:]
		} else {
			go countWords(words, maps)
			break
		}
		count++
	}

	m := make(map[string]uint64)
	for i := 0; i < count; i++ {
		for k, v := range <-maps {
			m[k] += v
		}
	}

	var wordCountList []*word_pb.WordCount

	for key, value := range m {
		wordCountList = append(wordCountList, &word_pb.WordCount{Word: key, Count: value})
	}

	sort.Slice(wordCountList, func(i, j int) bool {
		return wordCountList[i].Count > wordCountList[j].Count
	})

	if len(wordCountList) > 10 {
		return wordCountList[:10]
	}

	return wordCountList
}

func countWords(words []string, maps chan map[string]uint64) {

	m := make(map[string]uint64)

	for _, word := range words {
		m[word]++
	}

	maps <- m
}
