# Go Easy Mail

Welcome to the documentation for **Go Easy Mail**, a simple and efficient email sender application developed in Go. This application allows for sending emails either through direct input or by reading from a file. It supports customizable concurrency for enhanced performance.


## Configuration

### Using Environment Variables
For security and customization, the application uses environment variables to handle sensitive information and user-specific settings. You should create a .env file in the root directory of your application with the following variables:


`
EMAIL_ADDRESS=sender@email.com
EMAIL_PW=password
EMAIL_SENDER_NAME=Sender
EMAIL_SUBJECT=Hello
`
`
### Host
You can change the host name in mail/send.go, defaults to zoho.mail.eu


## Running the Application

### Input Mode

Run the application in input mode to manually enter email addresses interactively. In this mode, you can type each email address and press `Enter` to send an email immediately.

`bash
go run main.go
`

Simply start the application with the above command, enter an email address, and hit `Enter` to send.

### Batch Mode from File

To send emails from a predefined list, use the `-f` flag. The application expects a CSV file named `emails.csv` containing one email address per line.

`bash
go run main.go -f
`

Make sure that the `emails.csv` file is located in the same directory as your application or provide the relative path. Each email address should be on its own line without any additional characters or text.

### Changing Worker Count

To adjust the number of concurrent workers that handle sending emails, use the `-workers` flag followed by the number you wish to use.

`bash
go run main.go -workers 10
`

This flag can be combined with either mode to enhance performance based on your operational needs.

### Customizing Email Templates

The application uses a basic HTML template for the emails, which can be edited to suit your needs. You can find and modify this template in:

`
templates/hello.html
`

Edit this file to change the format or content of the emails sent by the application.
