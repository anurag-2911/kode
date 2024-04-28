package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/anurag-2911/kode/app/chatcompletion/model"
)

const (
	apiURL = "https://api.openai.com/v1/chat/completions"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func CallOpenAI() {
	http.HandleFunc("/call-openai", handler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
        return
    }

    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        http.Error(w, "API key not set in environment variables", http.StatusInternalServerError)
        return
    }

    var userInput struct {
        Content string `json:"content"`
    }

    if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
        http.Error(w, "Error decoding request body", http.StatusBadRequest)
        return
    }

    messages := []Message{
        {
            Role:    "system",
            Content: "You are a helpful assistant.",
        },
        {
            Role:    "user",
            Content: userInput.Content,
        },
    }

    chatReq := struct {
        Model    string    `json:"model"`
        Messages []Message `json:"messages"`
    }{
        Model:    "gpt-4",
        Messages: messages,
    }

    requestBody, err := json.Marshal(chatReq)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error marshalling request: %v", err), http.StatusInternalServerError)
        return
    }

    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
    if err != nil {
        http.Error(w, fmt.Sprintf("Error creating request: %v", err), http.StatusInternalServerError)
        return
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error making request: %v", err), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var openAIResp model.OpenAIResponse
    if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
        http.Error(w, fmt.Sprintf("Error decoding response body: %v", err), http.StatusInternalServerError)
        return
    }

    if len(openAIResp.Choices) > 0 && len(openAIResp.Choices[0].Message.Content) > 0 {
        fmt.Fprint(w, openAIResp.Choices[0].Message.Content)
    } else {
        http.Error(w, "No content found in response", http.StatusInternalServerError)
    }
}

