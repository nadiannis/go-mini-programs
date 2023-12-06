# Send Email

Send email from the command-line.

## Usage

- Make a copy of `env.example` file & rename it to `.env`.

  Write the SMTP server host, SMTP server port, sender email, & sender password. I use a Gmail account to test sending the email.

  ```
  SMTP_HOST=smtp.gmail.com
  SMTP_PORT=587

  SENDER_EMAIL=example@gmail.com
  SENDER_PASSWORD=abcd efgh ijkl mnop
  ```

  Note:

  > If you use Gmail & provide your account password, there will be an error. Since May 2022, Google doesn’t support less secure app. So you have to enable 2-step verification & use an app password for authentication. If you already generate the app password, you should use it instead of your regular Gmail password. Learn how to create & use app password [here](https://support.google.com/accounts/answer/185833).

- Compile & run the program.

  ```sh
  go run main.go
  ```

  Then provide the recipient email, email subject, & email body.

  ```
  >>> Recipient email (separate with ; if more than one):
  annisanadianeyla@gmail.com;neylanadiaannisa@gmail.com

  >>> Subject:
  Example Email

  >>> Body (enter :s to save):
  Hello, world!

  I'm sending email through CLI.
  :s
  Saved!

  Sending email...
  Email sent successfully
  ```
