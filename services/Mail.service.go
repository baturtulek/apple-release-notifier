package services

import (
	"log"
	"os"
	"strconv"

	"github.com/baturtulek/apple-release-notifier/types"
	mail "github.com/xhit/go-simple-mail/v2"
)

// SendMail Sends Mail to Clients About New Release Informations
func SendMail(newReleasesNotExistsInOldReleases []types.Release) {
	smtpClient := getSmtpClient()
	email := prepareMail(newReleasesNotExistsInOldReleases)
	err := email.Send(smtpClient)
	if err != nil {
		log.Fatal("ERROR - SendMail: ", err)
	}
	log.Print("Mail sent to contacts.")
}

func getSmtpServer() *mail.SMTPServer {
	SENDER_MAIL_SMTP_HOST,
		SENDER_MAIL_SMTP_PORT,
		SENDER_MAIL_ADDRESS,
		SENDER_MAIL_PASSWORD := getMailServerConfig()

	smtpServer := mail.NewSMTPClient()
	smtpServer.Host = SENDER_MAIL_SMTP_HOST
	smtpServer.Port = SENDER_MAIL_SMTP_PORT
	smtpServer.Username = SENDER_MAIL_ADDRESS
	smtpServer.Password = SENDER_MAIL_PASSWORD
	smtpServer.Encryption = mail.EncryptionTLS

	return smtpServer
}

func getSmtpClient() *mail.SMTPClient {
	server := getSmtpServer()
	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal("ERROR - getSmtpClient - SMPT Connect: ", err)
	}
	return smtpClient
}

func getMailServerConfig() (string, int, string, string) {
	var SENDER_MAIL_SMTP_HOST = os.Getenv("SENDER_MAIL_SMTP_HOST")
	var SENDER_MAIL_SMTP_PORT, _ = strconv.Atoi(os.Getenv("SENDER_MAIL_SMTP_PORT"))
	var SENDER_MAIL_ADDRESS = os.Getenv("SENDER_MAIL_ADDRESS")
	var SENDER_MAIL_PASSWORD = os.Getenv("SENDER_MAIL_PASSWORD")
	return SENDER_MAIL_SMTP_HOST, SENDER_MAIL_SMTP_PORT, SENDER_MAIL_ADDRESS, SENDER_MAIL_PASSWORD
}

func prepareMail(newReleasesNotExistsInOldReleases []types.Release) *mail.Email {
	htmlBody := prepareNewReleaseMailBody(newReleasesNotExistsInOldReleases)
	htmlMailBody := getHtmlMailBody(htmlBody)
	email := mail.NewMSG()

	email.SetFrom(os.Getenv("SENDER_MAIL_NAME") + "<" + os.Getenv("SENDER_MAIL_ADDRESS") + ">")

	mailContactsArr := ReadMailContactsFromFile()
	for _, mailAddress := range mailContactsArr {
		email.AddTo(mailAddress)
	}

	email.SetSubject("New Apple Release")
	email.SetBody(mail.TextHTML, htmlMailBody)

	return email
}

func prepareNewReleaseMailBody(newReleasesNotExistsInOldReleases []types.Release) string {
	var message string = ""
	message += "<p>Today Apple announces <b>" + strconv.Itoa(len(newReleasesNotExistsInOldReleases)) + "</b> new release(s).</p>"
	for i, newRelease := range newReleasesNotExistsInOldReleases {
		message += "<p>"
		message += "<b style='display: inline-block;'>" + strconv.Itoa(i+1) + " - </b>"
		if newRelease.Platform != "" {
			message += "<span style='padding-left: 5px'>"
			message += "<b>Platform: </b>" + newRelease.Platform
			message += "</span>"
		}
		if newRelease.Version != "" {
			message += "<br>"
			message += "<span style='padding-left: 26px'>"
			message += "<b>Version: </b>" + newRelease.Version
			message += "</span>"
		}
		if newRelease.Code != "" {
			message += "<br>"
			message += "<span style='padding-left: 26px'>"
			message += "<b>Code: </b>" + newRelease.Code
			message += "</span>"
		}
		message += "</p>"
	}
	return message
}

func getHtmlMailBody(mailBody string) string {
	return `
	<html>
		<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
			<title>New Releases Available</title>
		</head>
	<body>` + mailBody + "</body></html>"
}
