package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"last-project/app/config/gemini_config"
	"os"
	"strings"
	"time"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type InventoryAiResponse struct {
	EOQRecomendation int    `json:"eoq_recomendation"`
	Decision         string `json:"decision"`
	Insight          string `json:"insight"`
}

func Gemini(itemName string, currentStock int, rop int, eoq int, avgDailyDemand int, leadTimeDays int) (int, string, string, error) {
	ctx := context.Background()

	client, clientError := genai.NewClient(ctx, option.WithAPIKey(gemini_config.GEMINI_API_KEY))

	if clientError != nil {
		return 0, "", "", NewInternalServerError(clientError.Error())
	}

	defer client.Close()

	geminiModel := client.GenerativeModel(os.Getenv("MODEL"))

	now := time.Now().Format("2006-01-02")

	prompt := fmt.Sprintf(`
	You are an expert inventory management AI specializing in supply chain, EOQ, and ROP, position yourself as a PhD professor.

	Return STRICT JSON ONLY in the format below:
	{
		"eoq_recommendation": number,
		"decision": "REORDER NOW | MONITOR | SAFE",
		"ai_insight": "string"
	}

	Item Name: %s
	Current Stock: %d
	Reorder Point (ROP): %d
	EOQ (Economic Order Quantity): %d
	Average Daily Demand: %d
	Lead Time (Days): %d
	Today: %s

	RULES:
	1. If current stock <= ROP → decision MUST be "REORDER NOW"
	2. If stock is slightly above ROP → "MONITOR"
	3. If stock is far above ROP → "SAFE"
	4. EOQ Recommendation may adjust from provided EOQ if AI believes smarter value exists based on risk reasoning.
	5. ai_insight must be:
		- Written in indonesian language
		- Contextual, analytical
		- Explains urgency, risk of stockout, potential operational impact
		- Provide recommended strategy
	6. Use \\n for line breaks
	7. Maximum 400 words
	8. NO markdown, no explanation outside JSON
	`, itemName, currentStock, rop, eoq, avgDailyDemand, leadTimeDays, now)

	resp, err := geminiModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return 0, "", "", NewInternalServerError(err.Error())
	}

	if len(resp.Candidates) == 0 {
		return 0, "", "", NewInternalServerError("no response from Gemini")
	}

	var text string
	for _, part := range resp.Candidates[0].Content.Parts {
		if content, ok := part.(genai.Text); ok {
			text += string(content)
		}
	}

	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")

	if start == -1 || end == -1 {
		return 0, "", "", NewInternalServerError("invalid AI JSON response\nRaw: " + text)
	}

	jsonStr := text[start : end+1]

	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	jsonStr = strings.ReplaceAll(jsonStr, "\r", "")
	jsonStr = strings.ReplaceAll(jsonStr, "\t", "")
	jsonStr = strings.ReplaceAll(jsonStr, "```json", "")
	jsonStr = strings.ReplaceAll(jsonStr, "```", "")
	jsonStr = strings.TrimSpace(jsonStr)

	if !json.Valid([]byte(jsonStr)) {
		return 0, "", "", NewInternalServerError("invalid JSON format\nResponse: " + text)
	}

	var result InventoryAiResponse

	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return 0, "", "", NewInternalServerError("failed to parse JSON: " + err.Error() + text)
	}

	return result.EOQRecomendation, result.Decision, result.Insight, nil
}
