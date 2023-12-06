package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/mail"
	"net/smtp"
	"os"
	"sendemail/pkg/env"
	"strconv"
	"strings"
	"time"
)

type smtpServer struct {
	host string
	port int
}

type sender struct {
	email    string
	password string
	smtpServer
}

func main() {
	smtpPort, err := strconv.Atoi(env.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	sender := newSender(env.Getenv("SENDER_EMAIL"), env.Getenv("SENDER_PASSWORD"), env.Getenv("SMTP_HOST"), smtpPort)

	emails := getInput(reader, ">>> Recipient email (separate with ; if more than one):\n")
	recipientEmails := strings.Split(emails, ";")
	for _, recipientEmail := range recipientEmails {
		if !isValidEmail(recipientEmail) {
			log.Fatalln("Invalid email:", recipientEmail)
		}
	}

	subject := getInput(reader, "\n>>> Subject:\n")
	body := getMultilineInput(reader, "\n>>> Body (enter :s to save):\n")

	fmt.Println("\nSending email...")
	if err := sender.sendEmail(recipientEmails, subject, body); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Email sent successfully")
}

func (s *sender) sendEmail(to []string, subject, body string) error {
	serverAddress := fmt.Sprintf("%s:%d", s.host, s.port)
	auth := smtp.PlainAuth("", s.email, s.password, s.host)

	body = s.writeEmail(to, subject, body)
	if err := smtp.SendMail(serverAddress, auth, s.email, to, []byte(body)); err != nil {
		return err
	}
	return nil
}

func (s *sender) writeEmail(to []string, subject, body string) string {
	var message string
	var encodedBody bytes.Buffer
	delimiter := fmt.Sprintf("**=mail%d", time.Now().UnixNano())

	encodeMessage(body, &encodedBody)

	message += fmt.Sprintf("From: %s\r\n", s.email)
	message += fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))

	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "MIME-Version: 1.0\r\n"
	message += fmt.Sprintf("Content-Type: multipart/alternative; boundary=\"%s\"\r\n", delimiter)

	message += fmt.Sprintf("--%s\r\n", delimiter)
	message += "Content-Transfer-Encoding: quoted-printable\r\n"
	message += fmt.Sprintf("Content-Type: %s; charset=\"utf-8\"\r\n", "text/plain")
	message += "Content-Disposition: inline\r\n"
	message += fmt.Sprintf("%s\r\n", encodedBody.String())

	return message
}

func getInput(reader *bufio.Reader, prompt string) string {
	if len(prompt) != 0 {
		fmt.Print(prompt)
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimSpace(input)
}

func getMultilineInput(reader *bufio.Reader, prompt string) string {
	if len(prompt) != 0 {
		fmt.Print(prompt)
	}

	var lines []string
	for {
		line := getInput(reader, "")
		if strings.ToLower(line) == ":s" {
			fmt.Println("Saved!")
			break
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func newSender(email, password, host string, port int) *sender {
	return &sender{
		email:    email,
		password: password,
		smtpServer: smtpServer{
			host: host,
			port: port,
		},
	}
}

func encodeMessage(body string, encodedBody *bytes.Buffer) {
	writer := quotedprintable.NewWriter(encodedBody)
	if _, err := writer.Write([]byte(body)); err != nil {
		log.Fatalln(err)
	}
	defer writer.Close()
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
