package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bpradana/tars/llm"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/template"
)

// WeatherInfo represents structured weather data
type WeatherInfo struct {
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Description string  `json:"description"`
}

// Person represents structured person data
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	City    string   `json:"city"`
	Job     string   `json:"job"`
	Hobbies []string `json:"hobbies"`
}

// Product represents structured product data
type Product struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	InStock     bool    `json:"in_stock"`
	Rating      float64 `json:"rating"`
}

// AnalysisResult represents structured analysis data
type AnalysisResult struct {
	Sentiment       string   `json:"sentiment"`
	Confidence      float64  `json:"confidence"`
	KeyPoints       []string `json:"key_points"`
	Summary         string   `json:"summary"`
	Recommendations []string `json:"recommendations"`
}

func main() {
	fmt.Println("=== Structured Output Examples ===")

	// Example 1: Weather information
	fmt.Println("\n--- Weather Information Example ---")
	weatherExample()

	// Example 2: Person information
	fmt.Println("\n--- Person Information Example ---")
	personExample()

	// Example 3: Product information
	fmt.Println("\n--- Product Information Example ---")
	productExample()

	// Example 4: Text analysis
	fmt.Println("\n--- Text Analysis Example ---")
	analysisExample()

	// Example 5: Error handling with structured output
	fmt.Println("\n--- Error Handling Example ---")
	errorHandlingExample()
}

func weatherExample() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	template := template.From(
		message.FromSystem("You are a weather assistant. Provide weather information in the requested format."),
		message.FromUser("What's the weather like in Paris today? Provide the information in the specified JSON format."),
	)

	var weatherInfo WeatherInfo
	ctx := context.Background()
	response, err := provider.Invoke(ctx, template,
		llm.WithStructuredOutput(&weatherInfo),
		llm.WithTemperature(0.3), // Lower temperature for more consistent structured output
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// The response content will be JSON that matches the WeatherInfo struct
	fmt.Printf("Raw response: %s\n", response.GetContent())

	// The structured data is already parsed into the weatherInfo struct
	fmt.Printf("Structured data: Temperature=%.1f°C, Condition=%s, Humidity=%d%%, Wind Speed=%.1f km/h\n",
		weatherInfo.Temperature, weatherInfo.Condition, weatherInfo.Humidity, weatherInfo.WindSpeed)
	fmt.Printf("Description: %s\n", weatherInfo.Description)
}

func personExample() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	template := template.From(
		message.FromSystem("You are a data assistant. Create realistic person profiles in the specified format."),
		message.FromUser("Create a profile for a software engineer named Alice. Include realistic details."),
	)

	var person Person
	ctx := context.Background()
	response, err := provider.Invoke(ctx, template,
		llm.WithStructuredOutput(&person),
		llm.WithTemperature(0.5),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Raw response: %s\n", response.GetContent())
	fmt.Printf("Person: %s, Age: %d, City: %s, Job: %s\n",
		person.Name, person.Age, person.City, person.Job)
	fmt.Printf("Hobbies: %v\n", person.Hobbies)
}

func productExample() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	template := template.From(
		message.FromSystem("You are a product catalog assistant. Provide product information in the specified format."),
		message.FromUser("Create a product entry for a wireless Bluetooth headphones with noise cancellation."),
	)

	var product Product
	ctx := context.Background()
	response, err := provider.Invoke(ctx, template,
		llm.WithStructuredOutput(&product),
		llm.WithTemperature(0.3),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Raw response: %s\n", response.GetContent())
	fmt.Printf("Product: %s\n", product.Name)
	fmt.Printf("Price: $%.2f\n", product.Price)
	fmt.Printf("Category: %s\n", product.Category)
	fmt.Printf("In Stock: %t\n", product.InStock)
	fmt.Printf("Rating: %.1f/5\n", product.Rating)
	fmt.Printf("Description: %s\n", product.Description)
}

func analysisExample() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	textToAnalyze := "The new smartphone features an amazing camera and great battery life. However, the price is quite high and the software could be more intuitive. Overall, it's a solid device with room for improvement."

	template := template.From(
		message.FromSystem("You are a text analysis assistant. Analyze the given text and provide structured insights."),
		message.FromUser(fmt.Sprintf("Analyze this text: '%s'", textToAnalyze)),
	)

	var analysis AnalysisResult
	ctx := context.Background()
	response, err := provider.Invoke(ctx, template,
		llm.WithStructuredOutput(&analysis),
		llm.WithTemperature(0.2),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Raw response: %s\n", response.GetContent())
	fmt.Printf("Sentiment: %s (Confidence: %.2f)\n", analysis.Sentiment, analysis.Confidence)
	fmt.Printf("Key Points: %v\n", analysis.KeyPoints)
	fmt.Printf("Summary: %s\n", analysis.Summary)
	fmt.Printf("Recommendations: %v\n", analysis.Recommendations)
}

func errorHandlingExample() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Example with a template that might cause issues
	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser(""), // Empty message to demonstrate validation
	)

	var weatherInfo WeatherInfo
	ctx := context.Background()
	_, err := provider.Invoke(ctx, template,
		llm.WithStructuredOutput(&weatherInfo),
	)
	if err != nil {
		fmt.Printf("Expected error occurred: %v\n", err)
	} else {
		fmt.Println("Unexpected success with invalid template")
	}
}
