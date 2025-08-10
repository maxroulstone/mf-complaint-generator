package email

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/maxroulstone/mf-complaint-generator/pkg/person"
)

type EmailType int

const (
	ComplaintEmail EmailType = iota
	PasswordEmail
	ChaserEmail
)

func CreateComplaintMsg(p person.FakePerson, zipData []byte, includePasswords bool, pdfPassword, zipPassword string) string {
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700")
	zipBase64 := base64.StdEncoding.EncodeToString(zipData)
	
	body := fmt.Sprintf(`Dear Sir/Madam,

Please find attached my formal complaint regarding discretionary commission arrangements on my motor finance agreement.

I look forward to your prompt response and investigation into this matter.

Best regards,
%s`, p.FullName())

	if includePasswords {
		body += fmt.Sprintf(`

---
ATTACHMENT PASSWORDS:
PDF Password: %s
Zip Password: %s

Please use these passwords to access the attached documents.`, pdfPassword, zipPassword)
	}

	msgContent := fmt.Sprintf(`From: %s <%s>
To: complaints@motorfinance.com
Subject: Motor Finance Discretionary Commission Complaint - %s
Date: %s
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="boundary123"

--boundary123
Content-Type: text/plain; charset=UTF-8

%s

--boundary123
Content-Type: application/zip; name="complaint_documents.zip"
Content-Transfer-Encoding: base64
Content-Disposition: attachment; filename="complaint_documents.zip"

%s
--boundary123--
`, p.FullName(), p.Email, p.FullName(), timestamp, body, zipBase64)

	return msgContent
}

func CreatePasswordMsg(p person.FakePerson, pdfPassword, zipPassword string) string {
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700")
	
	body := fmt.Sprintf(`Dear Sir/Madam,

Further to my previous email containing my motor finance complaint, please find below the passwords required to access the attached documents:

PDF Password: %s
Zip Password: %s

Please use these passwords to extract and view my complaint documentation.

Best regards,
%s`, pdfPassword, zipPassword, p.FullName())

	msgContent := fmt.Sprintf(`From: %s <%s>
To: complaints@motorfinance.com
Subject: Re: Motor Finance Complaint - Access Passwords - %s
Date: %s
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8

%s
`, p.FullName(), p.Email, p.FullName(), timestamp, body)

	return msgContent
}

func CreateChaserMsg(p person.FakePerson, daysSince int) string {
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700")
	
	body := fmt.Sprintf(`Dear Sir/Madam,

I am writing to follow up on my motor finance discretionary commission complaint submitted %d days ago.

I have not yet received an acknowledgment or response to my complaint. According to FCA guidelines, complaints should be acknowledged within 3 business days and resolved within 8 weeks.

Could you please provide an update on the status of my complaint and confirm when I can expect a response?

My original complaint reference was regarding discretionary commission arrangements on my motor vehicle finance.

I look forward to your prompt response.

Best regards,
%s`, daysSince, p.FullName())

	msgContent := fmt.Sprintf(`From: %s <%s>
To: complaints@motorfinance.com
Subject: CHASER: Motor Finance Complaint Status Update Required - %s
Date: %s
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8

%s
`, p.FullName(), p.Email, p.FullName(), timestamp, body)

	return msgContent
}
