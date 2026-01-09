package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var filtered []BrainrotMeme
	for _, meme := range memes {
		if meme.Views > minViews {
			filtered = append(filtered, meme)
		}
	}

	if len(filtered) < 2 {
		return filtered
	}

	for i := 0; i < len(filtered)-1; i++ {
		for j := 0; j < len(filtered)-i-1; j++ {
			if filtered[j].TrendLevel < filtered[j+1].TrendLevel {
				filtered[j], filtered[j+1] = filtered[j+1], filtered[j]
			}
		}
	}

	return filtered
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)
	for _, meme := range memes {
		impact[meme.Category] += meme.Views
	}
	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var names []string
	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			names = append(names, meme.Name)
		}
	}
	return names
}

func main() {
	memes := []BrainrotMeme{
		{Name: "мем1", TrendLevel: 8, Category: "Subo Bratik", Views: 45.5},
		{Name: "мем2", TrendLevel: 10, Category: "TUNTUNTUNSAHUR", Views: 120.0},
		{Name: "мем3", TrendLevel: 6, Category: "Sigma", Views: 60.2},
		{Name: "мем4", TrendLevel: 9, Category: "Skibidi", Views: 30.0},
		{Name: "мем5", TrendLevel: 7, Category: "Mewing", Views: 55.5},
		{Name: "мем6", TrendLevel: 5, Category: "Sigma", Views: 40.0},
		{Name: "мем7", TrendLevel: 4, Category: "Other", Views: 25.3},
		{Name: "мем8", TrendLevel: 9, Category: "Sigma", Views: 70.1},
	}

	topTrending := FindTopTrending(memes, 50.0)
	fmt.Println("модные молодежные (views > 50):")
	for _, meme := range topTrending {
		fmt.Printf("%s (TrendLevel: %d, Views: %.1f)\n", meme.Name, meme.TrendLevel, meme.Views)
	}

	impact := CalculateCategoryImpact(memes)
	fmt.Println("\nпросмотры категории:")
	for cat, views := range impact {
		fmt.Printf("%s: %.1f\n", cat, views)
	}

	filteredNames := FilterByComplexCondition(memes)
	fmt.Println("\nотфильтрованные по условию:")
	for _, name := range filteredNames {
		fmt.Println(name)
	}
}