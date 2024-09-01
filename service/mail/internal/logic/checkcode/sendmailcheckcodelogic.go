package checkcodelogic

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"math/rand"
	"net/smtp"
	"time"
	"zerooj/service/mail/pb/mail"
	"zerooj/service/mail/static"

	"zerooj/service/mail/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMailCheckCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMailCheckCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMailCheckCodeLogic {
	return &SendMailCheckCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送邮箱验证码
func (l *SendMailCheckCodeLogic) SendMailCheckCode(in *mail.SendMailCheckCodeReq) (*mail.SendMailCheckCodeResp, error) {
	// 准备数据
	smtpCfg := l.svcCtx.Config.SMTP
	auth := smtp.PlainAuth("", smtpCfg.Username, smtpCfg.Password, smtpCfg.Host)
	sender := smtpCfg.Sender
	recipient := in.Email
	addr := fmt.Sprintf("%s:%d", smtpCfg.Host, smtpCfg.Port)
	checkCode := fmt.Sprintf("%06d", rand.Int()%1000000)
	subject := "ZeroOJ - 验证码获取"

	// 准备邮件内容
	var buf bytes.Buffer
	// 先准备邮件头
	buf.WriteString(fmt.Sprintf("To: %s\r\n", recipient))
	buf.WriteString(fmt.Sprintf("From: ZeroOJ <%s>\r\n", sender))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	buf.WriteString("\r\n")
	// 再准备邮件体
	tmpl := template.Must(template.ParseFS(static.MailTemplateFS, "template/checkcode.html"))
	err := tmpl.Execute(&buf, map[string]string{
		"Subject":       subject,
		"UserOperation": in.UserOperation,
		"CheckCode":     checkCode,
	})
	if err != nil {
		return nil, err
	}

	key := in.RedisKey
	// 将验证码写入Redis，有效期5min
	rdb := l.svcCtx.RDB
	err = rdb.SetEx(context.Background(), key, checkCode, 5*60*time.Second).Err()
	if err != nil {
		return nil, err
	}

	// 发送邮件
	err = smtp.SendMail(addr, auth, sender, []string{in.Email}, buf.Bytes())
	if err != nil {
		return nil, err
	}

	return &mail.SendMailCheckCodeResp{
		CheckCode: checkCode,
	}, nil
}
