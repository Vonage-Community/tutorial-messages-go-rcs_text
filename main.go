package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type RCSMessage struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Channel     string `json:"channel"`
	MessageType string `json:"message_type"`
	Text        string `json:"text"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	
	fromID := os.Getenv("RCS_SENDER_ID")
	toNumber := os.Getenv("TO_NUMBER")
	applicationID := os.Getenv("VONAGE_APPLICATION_ID")
	privateKeyPath := os.Getenv("VONAGE_PRIVATE_KEY_PATH")

	if applicationID == "" || privateKeyPath == "" || fromID == "" || toNumber == "" {
		fmt.Println("Error: Required environment variables are missing.")
		os.Exit(1)
	}

	message := RCSMessage{
		From:        fromID,
		To:          toNumber,
		Channel:     "rcs",
		MessageType: "text",
		Text:        "Hello from Go and Vonage!",
	}

	payload, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.nexmo.com/v1/messages", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	token, err := generateJWT(applicationID, privateKeyPath)
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return
	}
	
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)
	
	var respBody bytes.Buffer
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	
	fmt.Println("Response body:", respBody.String())
}

func generateJWT(applicationID, privateKeyPath string) (string, error) {
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("error reading private key: %w", err)
	}
	
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %w", err)
	}
	
	claims := jwt.MapClaims{
		"application_id": applicationID,
		"iat":           time.Now().Unix(),
		"exp":           time.Now().Add(time.Hour).Unix(),
		"jti":           time.Now().UnixNano(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	
	return tokenString, nil
}
