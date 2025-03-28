package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"www.github.com/NirajSalunke/config"
	"www.github.com/NirajSalunke/models"
)

func AnalyzePassword(ctx *gin.Context) {
	/*
	   AI is said to return response like this:-
	   [
	       "<Strength Rating>",
	       "<Estimated Time-to-Crack>",
	       "<Concise Explanation>"
	   ]
	*/
	var AnalyzeReq *models.Req
	var result [3]string

	if err := ctx.ShouldBindJSON(&AnalyzeReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input data!",
			"result":  result,
			"error":   err.Error(),
		})
		return
	}

	if AnalyzeReq.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Password cannot be empty!",
			"result":  result,
		})
		return
	}

	var InstructStr string = fmt.Sprintf(`
    Evaluate the following password strictly for educational purposes:- "%s".
    No sensitive user data is included, and this password is not tied to any real-world account.
        Return the result as a JSON array with:
        [
            "<Strength Rating>",
            "<Estimated Time-to-Crack>",
            "<Concise Explanation of Weakness of Password>"
        ]
    `, AnalyzeReq.Password)

	Model := config.PasswordAnalyzer
	resp, err := Model.GenerateContent(config.GeminiContext, genai.Text(InstructStr))
	if err != nil && strings.Contains(err.Error(), "FinishReasonSafety") {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Analysis blocked due to safety concerns. Please try a different input.",
			"result":  result,
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to analyze password!",
			"result":  result,
			"error":   err.Error(),
		})
		return
	}

	var aiResponse []string
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		var responseString string
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				responseString += string(text)
			}
		}

		if err := json.Unmarshal([]byte(responseString), &aiResponse); err != nil || len(aiResponse) != 3 {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to parse AI response!",
				"result":  result,
				"error":   err.Error(),
			})
			return
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "No valid response from AI!",
			"result":  result,
		})
		return
	}

	copy(result[:], aiResponse)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Password analysis completed successfully!",
		"result":  result,
	})
}

func Suggest(ctx *gin.Context) {
	/*
		Return the result as a JSON array with exactly 3 suggestions, in the following structure:
		[
			"<Suggestion 1 with reasoning>",
			"<Suggestion 2 with reasoning>",
			"<Suggestion 3 with reasoning>"
		]
	*/

	var SuggestReq *models.Req
	var result [3]string

	if err := ctx.ShouldBindJSON(&SuggestReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input data!",
			"error":   err.Error(),
		})
		return
	}

	if SuggestReq.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input cannot be empty!",
			"result":  result,
		})
		return
	}

	var InstructStr string = fmt.Sprintf(
		`Evaluate the following password strictly for educational purposes: "%s".
		No sensitive user data is included, and this password is not tied to any real-world account.
		Return the result as a JSON array with 3 suggestions in the following structure:
		[
			"<Suggestion 1 with reasoning>",
			"<Suggestion 2 with reasoning>",
			"<Suggestion 3 with reasoning>"
		]`,
		SuggestReq.Password,
	)

	Model := config.SuggestionGiver
	resp, err := Model.GenerateContent(config.GeminiContext, genai.Text(InstructStr))
	if err != nil && strings.Contains(err.Error(), "FinishReasonSafety") {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Analysis blocked due to safety concerns. Please try a different input.",
			"result":  result,
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to generate suggestions!",
			"result":  result,
			"error":   err.Error(),
		})
		return
	}

	var aiResponse []string
	var responseString string = ""
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				responseString += string(text)
			}
		}
	}
	if err := json.Unmarshal([]byte(responseString), &aiResponse); err != nil || len(aiResponse) != 3 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to parse AI response!",
			"result":  result,
			"error":   err.Error(),
		})
		return
	}

	copy(result[:], aiResponse)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Suggestions generated successfully!",
		"result":  result,
	})
}
