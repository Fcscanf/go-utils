package emailclient

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

type EmailClient struct {
	// SMTP协议的服务器地址
	serverAddr string
	serverHost string
	// 发件人账号
	senderAddr mail.Address
	// 发件人设置的授权密码
	authPassWord string
}

type Email struct {
	// 主题
	Subject string
	// 接收人
	To []string
	// 正文
	Body string
	// 附件
	AttachFiles []string
	// 邮件内容类型
	MailType MailType
}

type MailType string

const (
	boundary          = "boundary123"
	MailText MailType = "plain"
	MailHtml MailType = "html"
)

func (emailC *EmailClient) Init(serverAddr, authPassWord string, senderAddr mail.Address) error {
	if serverAddr == "" || authPassWord == "" || senderAddr.Address == "" {
		return errors.New("serverHost, serverSSLHost, authPassWord and senderAddr are required")
	}
	if host, _, err := net.SplitHostPort(serverAddr); err == nil {
		emailC.serverHost = host
	} else {
		return err
	}
	emailC.senderAddr = senderAddr
	emailC.serverAddr = serverAddr
	emailC.authPassWord = authPassWord
	return nil
}

func (emailC *EmailClient) auth() smtp.Auth {
	return smtp.PlainAuth("", emailC.senderAddr.Address, emailC.authPassWord, emailC.serverHost)
}

// Send 使用非加密方式发送
func (emailC *EmailClient) Send(email Email) error {
	msg, err := emailC.buildMsg(email)
	if err != nil {
		return err
	}
	// 普通发送
	return smtp.SendMail(emailC.serverAddr, emailC.auth(), emailC.senderAddr.Address, email.To, msg)
}

// SSend 使用SSL协议安全的发送，部分服务器会有此限制，普通发送不成功
func (emailC *EmailClient) SSend(email Email) error {
	msg, err := emailC.buildMsg(email)
	if err != nil {
		return err
	}
	// 通过SSL协议安全发送
	return emailC.sendMailUsingTLS(email.To, msg)
}

func (emailC *EmailClient) buildMsg(email Email) ([]byte, error) {
	if len(email.To) == 0 || email.MailType == "" {
		return nil, errors.New("email to and mail type are required")
	}
	buf := bytes.NewBuffer(nil)

	// 邮件头
	buf.WriteString(fmt.Sprintf("From: %s\r\n", emailC.senderAddr.String()))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ",")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", email.Subject))

	if len(email.AttachFiles) > 0 {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	}

	// 正文
	buf.WriteString("Content-Type: text/" + string(email.MailType) + "; charset=UTF-8\r\n\r\n")
	buf.WriteString(email.Body)
	buf.WriteString("\r\n")

	// 附件
	for _, file := range email.AttachFiles {
		data, _ := os.ReadFile(file)
		_, filename := filepath.Split(file)

		buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		buf.WriteString("Content-Type: application/octet-stream\r\n")
		buf.WriteString("Content-Transfer-Encoding: base64\r\n")
		buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", filename))
		// RFC 要求 Base64 编码的 MIME 部分 每行最多 76 个字符，不处理会导致附件发送失败
		encoded := base64.StdEncoding.EncodeToString(data)
		for i := 0; i < len(encoded); i += 76 {
			end := i + 76
			if end > len(encoded) {
				end = len(encoded)
			}
			buf.WriteString(encoded[i:end] + "\r\n")
		}
		buf.WriteString("\r\n")
	}

	if len(email.AttachFiles) > 0 {
		buf.WriteString(fmt.Sprintf("--%s--", boundary))
	}
	return buf.Bytes(), nil
}

// Dial return a smtp client
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Panicln("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// SendMailUsingTLS 参考net/smtp的func SendMail()
// 使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
// len(to)>1时,to[1]开始提示是密送
func (emailC *EmailClient) sendMailUsingTLS(to []string, msg []byte) (err error) {

	//create smtp client
	c, err := dial(emailC.serverAddr)
	if err != nil {
		return err
	}
	defer func(c *smtp.Client) {
		if closeErr := c.Close(); closeErr != nil {
			// 只记录非"连接已关闭"的错误
			if !strings.Contains(closeErr.Error(), "closed") {
				log.Println("Close Warning:", closeErr.Error())
			}
		}
	}(c)

	if emailC.auth() != nil {
		if ok, _ := c.Extension("AUTH"); !ok {
			return errors.New("email client smtp auth : server doesn't support AUTH")
		}
		if err = c.Auth(emailC.auth()); err != nil {
			return err
		}
	}

	if err = c.Mail(emailC.senderAddr.Address); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
