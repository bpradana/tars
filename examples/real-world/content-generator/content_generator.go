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

// ContentGenerator provides templates for different content types
type ContentGenerator struct {
	provider llm.BaseProvider
}

// NewContentGenerator creates a new content generator
func NewContentGenerator(provider llm.BaseProvider) *ContentGenerator {
	return &ContentGenerator{
		provider: provider,
	}
}

// GenerateBlogPost generates a blog post based on topic and style
func (cg *ContentGenerator) GenerateBlogPost(ctx context.Context, topic, style, tone string) (string, error) {
	template := template.From(
		message.FromSystem("You are a professional content writer. Write engaging, well-structured blog posts."),
		message.FromUser("Write a {{.Style}} blog post about {{.Topic}} in a {{.Tone}} tone. Include an introduction, main points, and conclusion."),
	)

	invokedTemplate := template.Invoke(struct {
		Topic string
		Style string
		Tone  string
	}{
		Topic: topic,
		Style: style,
		Tone:  tone,
	})

	response, err := cg.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to generate blog post: %v", err)
	}

	return response.GetContent(), nil
}

// GenerateEmail generates a professional email
func (cg *ContentGenerator) GenerateEmail(ctx context.Context, recipient, purpose, context string) (string, error) {
	template := template.From(
		message.FromSystem("You are a professional email writer. Write clear, concise, and professional emails."),
		message.FromUser("Write a professional email to {{.Recipient}} for the purpose of {{.Purpose}}. Context: {{.Context}}. Include a proper greeting and closing."),
	)

	invokedTemplate := template.Invoke(struct {
		Recipient string
		Purpose   string
		Context   string
	}{
		Recipient: recipient,
		Purpose:   purpose,
		Context:   context,
	})

	response, err := cg.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to generate email: %v", err)
	}

	return response.GetContent(), nil
}

// GenerateSocialMediaPost generates a social media post
func (cg *ContentGenerator) GenerateSocialMediaPost(ctx context.Context, platform, topic, hashtags string) (string, error) {
	template := template.From(
		message.FromSystem("You are a social media expert. Create engaging posts optimized for different platforms."),
		message.FromUser("Create a {{.Platform}} post about {{.Topic}}. Include relevant hashtags: {{.Hashtags}}. Keep it engaging and platform-appropriate."),
	)

	invokedTemplate := template.Invoke(struct {
		Platform string
		Topic    string
		Hashtags string
	}{
		Platform: platform,
		Topic:    topic,
		Hashtags: hashtags,
	})

	response, err := cg.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to generate social media post: %v", err)
	}

	return response.GetContent(), nil
}

// GenerateProductDescription generates a product description
func (cg *ContentGenerator) GenerateProductDescription(ctx context.Context, product, features, targetAudience string) (string, error) {
	template := template.From(
		message.FromSystem("You are a marketing copywriter. Write compelling product descriptions that convert."),
		message.FromUser("Write a compelling product description for {{.Product}}. Key features: {{.Features}}. Target audience: {{.TargetAudience}}. Focus on benefits and value proposition."),
	)

	invokedTemplate := template.Invoke(struct {
		Product        string
		Features       string
		TargetAudience string
	}{
		Product:        product,
		Features:       features,
		TargetAudience: targetAudience,
	})

	response, err := cg.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to generate product description: %v", err)
	}

	return response.GetContent(), nil
}

func main() {
	fmt.Println("=== Content Generator Example ===")

	// Create provider
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Create content generator
	generator := NewContentGenerator(provider)
	ctx := context.Background()

	// Example 1: Generate a blog post
	fmt.Println("\n--- Blog Post Generation ---")
	blogPost, err := generator.GenerateBlogPost(ctx,
		"artificial intelligence in healthcare",
		"informative",
		"professional")
	if err != nil {
		log.Printf("Error generating blog post: %v", err)
	} else {
		fmt.Printf("Blog Post:\n%s\n", blogPost)
	}

	// Example 2: Generate an email
	fmt.Println("\n--- Email Generation ---")
	email, err := generator.GenerateEmail(ctx,
		"john.doe@company.com",
		"follow-up on project proposal",
		"We discussed a new software project last week and I'd like to schedule a follow-up meeting")
	if err != nil {
		log.Printf("Error generating email: %v", err)
	} else {
		fmt.Printf("Email:\n%s\n", email)
	}

	// Example 3: Generate a social media post
	fmt.Println("\n--- Social Media Post Generation ---")
	socialPost, err := generator.GenerateSocialMediaPost(ctx,
		"LinkedIn",
		"remote work productivity tips",
		"#remotework #productivity #workfromhome")
	if err != nil {
		log.Printf("Error generating social media post: %v", err)
	} else {
		fmt.Printf("Social Media Post:\n%s\n", socialPost)
	}

	// Example 4: Generate a product description
	fmt.Println("\n--- Product Description Generation ---")
	productDesc, err := generator.GenerateProductDescription(ctx,
		"Smart Home Security Camera",
		"1080p HD video, night vision, motion detection, two-way audio, cloud storage",
		"homeowners and small business owners")
	if err != nil {
		log.Printf("Error generating product description: %v", err)
	} else {
		fmt.Printf("Product Description:\n%s\n", productDesc)
	}
}
