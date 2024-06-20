package main


import (
  "log"
  "github.com/mihonen/mailer/mail"
  "bufio"
  "fmt"
  "os"
  "strings"
  "flag"
  "encoding/csv"
  "sync"
)




func main(){

  fromFile := flag.Bool("f", false, "Read emails from emails.csv")
  numWorkers := flag.Int("workers", 5, "Number of worker goroutines")
  flag.Parse()

  if *fromFile{

    file, err := os.Open("emails.csv")
    if err != nil {
        log.Fatal("Unable to read input file emails.csv", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    emails := make(chan string, 100) // Buffered channel for emails
    var wg sync.WaitGroup

    // Start a pool of worker goroutines
    for i := 0; i < *numWorkers; i++ { // Adjust number of workers based on your needs
        wg.Add(1)
        go func() {
            defer wg.Done()
            for email := range emails {
                if err := mail.SendMail(email); err != nil {
                    log.Printf("Failed to send mail to %s: %v", email, err)
                } else {
                    log.Printf("Sent mail to %s", email)
                }
            }
        }()
    }

    // Read emails from CSV and send to channel
    for {
        record, err := reader.Read()
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            log.Fatal("An error occurred while reading the CSV file.", err)
        }
        emails <- record[0] // Send email to the channel
    }

    close(emails) // Close channel after all emails are sent
    wg.Wait()     // Wait for all goroutines to finish
    return
    }



  for {
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Enter email address to send to (or type 'exit' to quit): ")
      toEmail, err := reader.ReadString('\n')
      if err != nil {
          log.Fatalf("Failed to read from input: %v", err)
      }

      toEmail = strings.TrimSpace(toEmail) // Remove newline or any extra space

      if toEmail == "exit" {
          fmt.Println("Exiting...")
          break
      }

      // Send mail using a separate function for clarity
      if err := mail.SendMail(toEmail); err != nil {
          log.Printf("Failed to send mail to %s: %v", toEmail, err)
      } else {
          log.Printf("Sent mail to %s", toEmail)
      }
  }
  
}



