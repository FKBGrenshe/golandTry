package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order := topoSort(prereqs)
	fmt.Println(order)

}

func topoSort(m map[string][]string) []string {

	var order []string
	seen := make(map[string]bool)

	// 方法1: 短变量声明（最常用）
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
				order = append(order, "->")
			}
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
