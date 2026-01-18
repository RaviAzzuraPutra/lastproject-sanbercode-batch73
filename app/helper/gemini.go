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
	Decision string `json:"decision"`
	Insight  string `json:"insight"`
}

func GeminiInsight(itemName string, currentStock int, trxType string, trxQty int) (string, string, error) {
	ctx := context.Background()

	client, clientError := genai.NewClient(ctx, option.WithAPIKey(gemini_config.GEMINI_API_KEY))
	if clientError != nil {
		return "", "", NewInternalServerError(clientError.Error())
	}
	defer client.Close()

	geminiModel := client.GenerativeModel(os.Getenv("MODEL"))
	now := time.Now().Format("2006-01-02")

	prompt := fmt.Sprintf(`

		You are an expert inventory management AI And Position yourself as a PhD professor by offering critical, academic and theoretical advice, but still using everyday language...
		Your task is to provide deep, analytical, and critical advice, BUT you must explain complex academic concepts using very simple, everyday "Warung/UMKM" metaphors.

		Return STRICT JSON ONLY:
		{
		"decision": "SAFE | MONITOR | WARNING",
		"insight": "string"
		}

		Item Name: %s
		Current Stock: %d
		Transaction Type: %s
		Transaction Quantity: %d
		Date: %s

		RULES:
			1. Decision reflects risk level after transaction.
			2. ai_insight must be:
				- Written in Indonesian Language
				- Appreciation/Observation of the transaction.
				- Explain risk of stockout or overstock (in simple words)
				- Practical recommendation for the owner.
				- Conclusion
			3. Max 99 words
			4. No markdown, no explanation outside JSON
			5. No Jargon: Do NOT use terms like 'EOQ', 'Lead Time', 'Safety Stock', or 'FIFO' directly. 
           Instead, explain the logic behind them (e.g., instead of 'FIFO', say 'jual barang yang masuk duluan agar tidak basi').
		   6. Tone: Like a friendly professor explaining to a small business owner.
		   7. Analytical: Don't just say 'stock is enough'. Explain the 'Why' and 'What next'.
		   8. The conclusion must remain academic, critical and theoretical.
		`, itemName, currentStock, trxType, trxQty, now)

	resp, err := geminiModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", "", NewInternalServerError(err.Error())
	}

	if len(resp.Candidates) == 0 {
		return "", "", NewInternalServerError("no response from Gemini")
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
		return "", "", NewInternalServerError("invalid AI JSON response\nRaw: " + text)
	}

	jsonStr := strings.TrimSpace(text[start : end+1])
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	jsonStr = strings.ReplaceAll(jsonStr, "\r", "")
	jsonStr = strings.ReplaceAll(jsonStr, "\t", "")

	if !json.Valid([]byte(jsonStr)) {
		return "", "", NewInternalServerError("invalid JSON format\nResponse: " + text)
	}

	var result InventoryAiResponse
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return "", "", NewInternalServerError("failed to parse JSON: " + err.Error())
	}

	return result.Decision, result.Insight, nil
}
