# tutorial-messages-go-rcs_text
Sample Go application demonstrating how to send RCS text messages with the Vonage Messages API. Includes JWT authentication implementation and environment configuration. Companion code for a step-by-step tutorial on building rich, interactive messaging capabilities in Go applications.


# How to Send RCS Text Messages with Go
A Go application that allows you to send iRch Communications Services(RCS) text messages.

> You can find full step-by-step instructions on the [Vonage Developer Blog](#). (Not published yet)



## Prerequisites
1. [Go 1.18+ installed](https://go.dev/dl/)
2. [ngrok installed for exposing your local server to the internet.](https://ngrok.com/downloads/mac-os)
3. [Vonage Developer Account](https://developer.vonage.com/sign-up)
4. A registered RCS Business Messaging (RBM) Agent.
5. A phone with RCS capabilities for testing.



## Instructions
1. Clone this repo
2. Initialize your Go application and install dependencies:
```
go mod init rcs-text-golang
go get github.com/joho/godotenv github.com/golang-jwt/jwt/v4
```
4. Rename the `.env.example` file to `.env`, and add your `VONAGE_APPLICATION_ID` and `RCS_SENDER_ID` values.
5. Add your `private.key` file in the root of the project directory.
6. Start your Node server:
```
go run main.go
```
7. Create a tunnel using ngrok:
```
ngrok http 3000
```
8. Test your app by sending an RCS suggested reply from the command line:
```
curl -X POST https://YOUR_NGROK_URL/send-rcs-carousel \
  -H "Content-Type: application/json" \
  -d '{"to": "YOUR_RCS_TEST_NUMBER"}'
```
