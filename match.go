package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cdipaolo/sentiment"
	"github.com/kljensen/snowball"
	"github.com/pkg/errors"
)

type Item struct {
	Name        string
	Description string
	Category    string
	Rating      float64
}

type User struct {
	Name     string
	Age      int
	Location string
	Keywords []string
}

func main() {
	// Define some sample items
	items := []Item{
		{Name: "iPhone", Description: "A smartphone made by Apple", Category: "Electronics", Rating: 4.5},
		{Name: "Samsung TV", Description: "A television made by Samsung", Category: "Electronics", Rating: 4.8},
		{Name: "Adidas shoes", Description: "A pair of shoes made by Adidas", Category: "Footwear", Rating: 4.2},
		{Name: "Nike sneakers", Description: "A pair of sneakers made by Nike", Category: "Footwear", Rating: 4.4},
		{Name: "Harry Potter book", Description: "A book in the Harry Potter series", Category: "Books", Rating: 4.9},
		{Name: "The Lord of the Rings book", Description: "A book in The Lord of the Rings series", Category: "Books", Rating: 4.7},
	}

	// Define some sample users
	users := []User{
		{Name: "Alice", Age: 25, Location: "New York", Keywords: []string{"phone", "television", "books"}},
		{Name: "Bob", Age: 32, Location: "Los Angeles", Keywords: []string{"shoes", "sneakers"}},
		{Name: "Charlie", Age: 41, Location: "Chicago", Keywords: []string{"books", "television"}},
	}

	// Loop through the users and match them with items
	for _, user := range users {
		matchedItems := matchItems(items, user)
		fmt.Printf("%s's matched items:\n", user.Name)
		for _, item := range matchedItems {
			fmt.Printf("- %s (%s, rating: %f)\n", item.Name, item.Category, item.Rating)
		}
	}
}

func matchItems(items []Item, user User) []Item {
	// Determine the sentiment of each keyword in the user's preferences
	var keywordSentiments []float64
	for _, keyword := range user.Keywords {
		keywordSentiments = append(keywordSentiments, getSentiment(keyword))
	}

	// Find items that match the user's preferences
	var matchedItems []Item
	for _, item := range items {
		// Check if the item's category matches any of the user's keywords
		if contains(user.Keywords, item.Category) {
			matchedItems = append(matchedItems, item)
			continue
		}

		// Check if the item's name or description matches any of the user's keywords
		if containsKeywords(item.Name, keywordSentiments) || containsKeywords(item.Description, keywordSentiments) {
			matchedItems = append(matchedItems, item)
			continue
		}

		// Check if the item's rating is above a certain threshold
		if item.Rating >= 4.5 {
			matchedItems = append(matchedItems, item)
			continue
		}
	}

	return matchedItems
}

func getSentiment(text string) float64 {
	// Normalize the text
	normalizedText := normalize(text)

	// Instantiate the sentiment analyzer
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	// Analyze the sentiment of the text
	score := model.SentimentAnalysis(normalizedText, sentiment.English)
	return score.Score
}

func normalize(text string) string {
	// Convert the text to lowercase
	text = strings.ToLower(text)

	// Remove punctuation
	text = regexp.MustCompile(`[^\p{L}\p{N}\s]+`).ReplaceAllString(text, "")

	// Tokenize the text
	words := strings.Split(text, " ")

	// Stem the words using the Snowball stemmer
	for i, word := range words {
		stemmed, err := snowball.Stem(word, "english", true)
		if err == nil {
			words[i] = stemmed
		}
	}

	// Join the stemmed words back into a string
	text = strings.Join(words, " ")
	return text
}

func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func containsKeywords(text string, sentiments []float64) bool {
	// Normalize the text
	normalizedText := normalize(text)

	// Tokenize the text
	words := strings.Split(normalizedText, " ")

	// Check if any of the words match the keywords
	for i, word := range words {
		if i >= len(sentiments) {
			break
		}
		if sentiments[i] != 0 && strings.Contains(normalizedText, word) {
			return true
		}
	}

	return false
}


func matchItems(items []Item, user User) []Item {
	// Determine the sentiment of each keyword in the user's preferences
	var keywordSentiments []float64
	for _, keyword := range user.Keywords {
		keywordSentiments = append(keywordSentiments, getSentiment(keyword))
	}

	// Find items that match the user's preferences
	var matchedItems []Item
	for _, item := range items {
		// Check if the item's category matches any of the user's keywords
		if contains(user.Keywords, item.Category) {
			matchedItems = append(matchedItems, item)
			continue
		}

		// Check if the item's name or description matches any of the user's
if containsKeywords(item.Name, keywordSentiments) || containsKeywords(item.Description, keywordSentiments) {
matchedItems = append(matchedItems, item)
}
}
  
  
  return matchedItems
}


//The `matchItems` function takes a slice of `Item` structs and a `User` struct as input, and returns a slice of `Item` structs that match the user's preferences. It first determines the sentiment of each keyword in the user's preferences using the `getSentiment` function. It then iterates through each item in the `items` slice, and checks if the item's category matches any of the user's keywords using the `contains` function. If the category matches, the item is added to the `matchedItems` slice and the iteration continues with the next item. If the category does not match, the function checks if the item's name or description matches any of the user's keywords using the `containsKeywords` function. If the name or description matches, the item is added to the `matchedItems` slice.

//Overall, this code demonstrates how to use AI techniques to match items with a user's preferences. By using sentiment analysis to determine the sentiment of the user's keywords, we can more accurately identify items that match the user's preferences.

  
    
    
    
    
    
    
    






  
  
  
  
  
  
  
  
  
  
