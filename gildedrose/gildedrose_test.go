package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_UpdateQuality(t *testing.T) {
	tests := []struct {
		name     string
		items    []*gildedrose.Item
		expected []*gildedrose.Item
	}{
		{
			name: "Normal item",
			items: []*gildedrose.Item{
				{"foo", 10, 20},
			},
			expected: []*gildedrose.Item{
				{"foo", 9, 19},
			},
		},
		{
			name: "Aged Brie",
			items: []*gildedrose.Item{
				{"Aged Brie", 2, 0},
			},
			expected: []*gildedrose.Item{
				{"Aged Brie", 1, 1},
			},
		},
		{
			name: "Sulfuras",
			items: []*gildedrose.Item{
				{"Sulfuras, Hand of Ragnaros", 0, 80},
			},
			expected: []*gildedrose.Item{
				{"Sulfuras, Hand of Ragnaros", 0, 80},
			},
		},
		{
			name: "Backstage passes (long before concert)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			},
			expected: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
			},
		},
		{
			name: "Backstage passes (close to concert)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
			},
			expected: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
			},
		},
		{
			name: "Backstage passes (very close to concert)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
			},
			expected: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
			},
		},
		{
			name: "Backstage passes (after concert)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 0, 20},
			},
			expected: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
			},
		},
		{
			name: "Conjured item",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 3, 6},
			},
			expected: []*gildedrose.Item{
				{"Conjured Mana Cake", 2, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gildedrose.UpdateQuality(tt.items)

			for i, item := range tt.items {
				if item.Name != tt.expected[i].Name {
					t.Errorf("Name: Expected %s but got %s", tt.expected[i].Name, item.Name)
				}
				if item.SellIn != tt.expected[i].SellIn {
					t.Errorf("SellIn: Expected %d but got %d", tt.expected[i].SellIn, item.SellIn)
				}
				if item.Quality != tt.expected[i].Quality {
					t.Errorf("Quality: Expected %d but got %d", tt.expected[i].Quality, item.Quality)
				}
			}
		})
	}
}
