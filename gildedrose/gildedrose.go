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

func QualityModificationLogic(item *Item, isMinusSellIn bool) {
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		DecreaseItemQuality(item)
	} else {
		IncreaseItemQuality(item)
		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			if !isMinusSellIn {
				if item.SellIn < 11 || item.SellIn < 6 {
					IncreaseItemQuality(item)
				}
			} else {
				item.Quality = 0
			}
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		QualityModificationLogic(item, false)

		if item.Name != "Sulfuras, Hand of Ragnaros" {
			item.SellIn--
		}

		if item.SellIn < 0 {
			QualityModificationLogic(item, true)
		}
	}

}
