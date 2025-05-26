package service

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"time"

	"backend/config"
	"backend/database"
)

// 浅浅写一个负责邮件的发送和验证吧

type SmtpService struct{}

func NewSmtpService() *SmtpService {
	return &SmtpService{}
}

// 发送邮件
func (ss *SmtpService) SendMail(to, subject, body string) error {
	smtpConfig := config.GetSMTPConfig()
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)
	msg := "From: " + smtpConfig.Username + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body
	addr := smtpConfig.Host + ":" + smtpConfig.Port

	// Establishing TLS connection
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpConfig.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		log.Printf("Error establishing TLS connection: %v", err)
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, smtpConfig.Host)
	if err != nil {
		log.Printf("Error creating SMTP client: %v", err)
		return err
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		log.Printf("Error during SMTP authentication: %v", err)
		return err
	}

	if err = client.Mail(smtpConfig.Username); err != nil {
		log.Printf("Error setting sender: %v", err)
		return err
	}

	if err = client.Rcpt(to); err != nil {
		log.Printf("Error setting recipient: %v", err)
		return err
	}

	w, err := client.Data()
	if err != nil {
		log.Printf("Error getting data writer: %v", err)
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Printf("Error writing message: %v", err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Printf("Error closing writer: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}

// 验证验证码
func (ss *SmtpService) VerifyCode(email, code string) (bool, error) {
	redisClient := database.GetRedisClient()
	storedCode, err := redisClient.GetValue(email)
	if err != nil {
		log.Printf("Error retrieving code from Redis: %v", err)
		return false, err
	}
	if storedCode != code {
		return false, nil
	}

	err = redisClient.DeleteValue(email)
	if err != nil {
		log.Printf("Error deleting code from Redis: %v", err)
		return false, err
	}

	return true, nil
}

// 发送验证码
func (ss *SmtpService) SendCode(email, code string) error {
	redisClient := database.GetRedisClient()
	err := redisClient.SetValue(email, code, 5*time.Minute)
	if err != nil {
		log.Printf("Error saving code to Redis: %v", err)
		return err
	}

	subject := "面试者挑战验证码"
	body := "你的验证码是 " + code
	if err := ss.SendMail(email, subject, body); err != nil {
		redisClient.DeleteValue(email)
		log.Printf("Error sending email: %v", err)
		return err
	}

	return nil
}
