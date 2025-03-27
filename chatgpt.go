package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
)

type Message struct {
        Role    string `json:"role"`
        Content string `json:"content"`
}

type Request struct {
        Model    string    `json:"model"`
        Messages []Message `json:"messages"`
}

type Choice struct {
        Message Message `json:"message"`
}

type Response struct {
        Choices []Choice `json:"choices"`
}

func main() {
        apiKey := os.Getenv("OPENAI_API_KEY") // Securely get the API key from environment variables.
        if apiKey == "" {
                fmt.Println("Error: OPENAI_API_KEY environment variable not set.")
                return
        }

        apiUrl := "https://api.openai.com/v1/chat/completions" //ChatGPT API Endpoint

        prompt := "Write a short poem about the moon."

        requestBody := Request{
                Model: "gpt-3.5-turbo", // or "gpt-4", etc.
                Messages: []Message{
                        {
                                Role:    "user",
                                Content: prompt,
                        },
                },
        }

        jsonValue, _ := json.Marshal(requestBody)

        req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonValue))
        if err != nil {
                fmt.Printf("HTTP request error: %s\n", err)
                return
        }

        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+apiKey)

        client := &http.Client{}
        response, err := client.Do(req)

        if err != nil {
                fmt.Printf("HTTP request failed: %s\n", err)
                return
        }

        defer response.Body.Close()

        data, err := ioutil.ReadAll(response.Body)
        if err != nil {
                fmt.Printf("Error reading response body: %s\n", err)
                return
        }

        var result Response
        err = json.Unmarshal(data, &result)
        if err != nil {
                fmt.Printf("Error unmarshalling JSON: %s\n", err)
                return
        }

        if len(result.Choices) > 0 {
                fmt.Println(result.Choices[0].Message.Content)
        } else {
                fmt.Println("No response generated.")
        }
}
