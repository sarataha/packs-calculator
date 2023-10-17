package pack_calculator

import (
	"sort"

	"github.com/sarataha/packs-calculator/config"
)

// Pack represents a pack size and quantity
type Pack struct {
	Size     int
	Quantity int
}

// CalculatePacks calculates the number of packs needed for a given number of items
func CalculatePacks(items int) []Pack {
	if items == 0 || len(config.PackSizes) == 0 {
		return nil
	}

	// Sort pack sizes in descending order
	packSizes := config.PackSizes
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	smallestPack := packSizes[len(packSizes)-1]
	largestPack := packSizes[0]

	var neededPacks []Pack

	// If items are fewer than the smallest pack then use it
	if items < smallestPack {
		neededPacks = append(neededPacks, Pack{Size: smallestPack, Quantity: 1})

		return neededPacks
	}

	total := calculateTotal(packSizes, items, smallestPack, largestPack)

	return getOrderPacks(packSizes, neededPacks, total)
}

// calculateTotal calculates the total number of items that we can send in packs
func calculateTotal(packSizes []int, items, smallestPack, largestPack int) int {
	total := 0

	for items > 0 {
		// Check if items are dividable by the largest pack
		if items >= largestPack {
			packCount := items / largestPack
			total += packCount * largestPack
			items %= largestPack
		}

		// Try using smaller packs
		found := false

		for _, size := range packSizes {
			if items == size {
				total += size
				items = 0
				found = true
				break
			} else if items > size {
				count := items / size
				total += count * size
				items %= size
				found = true
			}
		}

		// Use smallest pack if nothing else is found
		if !found {
			total += smallestPack
			items = 0
		}
	}

	return total
}

// getOrderPacks returns packs that we need to send
func getOrderPacks(packSizes []int, neededPacks []Pack, total int) []Pack {
	var result []Pack

	for _, size := range packSizes {
		if total == size {
			result = append(result, Pack{Size: size, Quantity: 1})

			return append(neededPacks, result...)
		}

		if total > size {
			count := total / size
			mod := total % size

			result = append(result, Pack{Size: size, Quantity: count})
			total = mod
		}
	}

	return append(neededPacks, result...)
}
