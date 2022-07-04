package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []struct {
		nama   string
		harga  int
		rating float32
		likes  int
	}{
		{"indomie", 3000, 5, 150},
		{"laptop", 4000000, 4.5, 123},
		{"aqua", 3000, 4, 250},
		{"smart tv", 4000000, 4.5, 42},
		{"headphone", 4000000, 3.5, 90},
		{"very smart tv", 4000000, 3.5, 87},
	}

	sort.Slice(data, func(a, b int) bool {
		if data[a].harga == data[b].harga {
			if data[a].rating > data[b].rating {
				return data[a].likes > data[b].likes
			} else {
				return data[a].rating > data[b].rating
			}
		} else {
			if data[a].rating < data[b].rating {
				return data[a].likes > data[b].likes
			} else {
				return data[a].rating > data[b].rating
			}
		}
	})
	fmt.Println(data)

}
