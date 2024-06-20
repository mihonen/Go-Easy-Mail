package mail



import (
  "crypto/tls"
  "bytes"
    "html/template"
  gomail "gopkg.in/mail.v2"
  "github.com/mihonen/mailer/utils"
)


const host = "smtp.zoho.eu"

func SendMail(receiver string) error{
  // receiver := "villemihonen@yahoo.fi"
  from        := utils.EnvVariableStr("EMAIL_ADDRESS") 
  password    := utils.EnvVariableStr("EMAIL_PW") 
  sender_name := utils.EnvVariableStr("EMAIL_SENDER_NAME") 
  subject     := utils.EnvVariableStr("EMAIL_SUBJECT") 

  name := sender_name + " <" + from + ">"

  m := gomail.NewMessage()

  m.SetHeader("From", name)
  m.SetHeader("To", receiver)
  m.SetHeader("Subject", subject)


  t := template.New("action")

  var templatePath string

  templatePath = "templates/hello.html"



  t, err := template.ParseFiles(templatePath)
  if err != nil {
      return err
  }



  var tpl bytes.Buffer
  if err := t.Execute(&tpl, nil); err != nil {
      return err
  }

  stringHtml := tpl.String()
  

  m.SetBody("text/html", stringHtml)

  d := gomail.NewDialer(host, 465, from, password)

  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  if err := d.DialAndSend(m); err != nil {
    return err
  }
  

  return nil
}

