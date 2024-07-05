package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func IncreaseItemQuality(item *Item) {
	if item.Quality < 50 {
		item.Quality++
	}
}

func DecreaseItemQuality(item *Item) {
	if item.Quality > 0 && item.Name != "Sulfuras, Hand of Ragnaros" {
		item.Quality--
	}
}

func QualityModificationLogic(item *Item) {
	if item.Name != "Aged Brie" {
		DecreaseItemQuality(item)
		if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
			DecreaseItemQuality(item)
		} else {
			item.Quality = 0
		}
	} else {
		IncreaseItemQuality(item)
		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.SellIn < 11 {
				IncreaseItemQuality(item)
			}
			if item.SellIn < 6 {
				IncreaseItemQuality(item)
			}
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		if item.Name != "Sulfuras, Hand of Ragnaros" {
			item.SellIn--
		}

		QualityModificationLogic(item)

		if item.SellIn < 0 {
			QualityModificationLogic(item)
		}
	}

}
