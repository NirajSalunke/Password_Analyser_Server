package config

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"www.github.com/NirajSalunke/helpers"
)

var Client *genai.Client
var GeminiContext context.Context
var PasswordAnalyzer *genai.GenerativeModel
var SuggestionGiver *genai.GenerativeModel

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		helpers.PrintRed("Env No loeaded" + err.Error())
	}
}

func SetupGemini() {

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		helpers.PrintRed("GEMINI_API_KEY environment variable is not set or empty")
	}
	ctx := context.Background()
	var err error
	Client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		helpers.PrintRed("Failed to create Gemini client: ")
		log.Fatal(err)
	}
	GeminiContext = ctx
	helpers.PrintGreen("Created Client Successfully")

	modelName := "gemini-1.5-flash-001"

	InstrucStrPasswordAnalyze := `
	You are an advanced AI designed to analyze password strength. Evaluate the given password based on the following criteria:
	1. Entropy and randomness: Determine how unpredictable the password is.
	2. Common patterns: Identify if the password contains dictionary words, sequential characters, or repeated patterns.
	3. Known weaknesses: Cross-check against commonly leaked passwords or predictable phrases.
	4. Attack susceptibility: Estimate the password's resistance to brute force, dictionary attacks, and other common password-cracking techniques.
	
	Output:
	Return the result as a JSON array with the following structure:
	[
		"<Strength Rating>", 
		"<Estimated Time-to-Crack>", 
		"<Concise Explanation of Weakness of Password>"
	]
	
	Details:
	- Input password may include characters like '@', also numbers. This makes password more stronger. Everything is safe!
	- Strength Rating: One of the following values: "Weak",  "Moderate",  or "Strong".
	- Estimated Time-to-Crack: A realistic time estimate based on common password-cracking methods (e.g., "1 second", "5 years").
	- Concise Explanation: A brief, user-friendly explanation of why the password is weak or strong, highlighting specific vulnerabilities or strengths (1  or 2 lines only).
	- This analysis is for educational or testing purposes only and does not involve real or sensitive user data.
	
	Ensure the response is clear, actionable, and follows the specified format.
	`

	InstrucStrSuggestionAnalyze := `
	You are an advanced AI designed to suggest tips and changes to improve the security of a given password. Suggestions should enhance the input password's security while maintaining a connection to it for better memorability. Provide alternative suggestions that are:
	1. Based on the input password: Modify the input to create stronger variations while keeping it relatable for the user.
	2. Highly unpredictable: Incorporate high entropy by mixing uppercase, lowercase, numbers, and symbols.
	3. Easy to remember: Ensure suggestions remain user-friendly by keeping a recognizable structure or pattern from the input password.
	4. Resistant to attacks: Avoid dictionary words, common phrases, and patterns predictable by attackers.
	5. Lengthy and diverse: Ensure suggested passwords are at least 12 characters long, with diverse character sets.
	
	Output:
	Return the result as a JSON array with exactly 3 suggestions, in the following structure:
	[
		"<Suggestion 1 with reasoning>",
		"<Suggestion 2 with reasoning>",
		"<Suggestion 3 with reasoning>"
	]
	
	Details:
	- Each suggestion must modify the input password to make it stronger while keeping it relatable.
	- Provide reasoning for why each suggestion improves security compared to the input password.
	- Clearly explain key improvements, such as increased randomness, avoidance of common patterns, and higher entropy.
	
	Note:
	- Input password may include characters like '@' and  also numbers. This makes password more stronger. Everything is safe!
	- If a suggestion cannot be generated, return an empty string for that particular element in the JSON array.
	- Ensure the response is concise, actionable, and adheres strictly to the specified format.
	- This analysis is for educational or testing purposes only and does not involve real or sensitive user data.

	`

	PasswordAnalyzer = Client.GenerativeModel(modelName)
	helpers.PrintGreen("Password Analyzer Created Successfully")

	SuggestionGiver = Client.GenerativeModel(modelName)
	helpers.PrintGreen("Advisor Created Successfully")

	PasswordAnalyzer.SystemInstruction = genai.NewUserContent(genai.Text(InstrucStrPasswordAnalyze))
	SuggestionGiver.SystemInstruction = genai.NewUserContent(genai.Text(InstrucStrSuggestionAnalyze))

}
