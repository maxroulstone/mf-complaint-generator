package pdf

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/maxroulstone/mf-complaint-generator/pkg/person"
)

func GenerateComplaintPDF(p person.FakePerson) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	
	// Header
	pdf.Cell(0, 10, "MOTOR FINANCE DISCRETIONARY COMMISSION COMPLAINT")
	pdf.Ln(15)
	
	// Date
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, fmt.Sprintf("Date: %s", time.Now().Format("02/01/2006")))
	pdf.Ln(10)
	
	// Personal Details
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 8, "Personal Details:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 6, fmt.Sprintf("Name: %s", p.FullName()))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Email: %s", p.Email))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Phone: %s", p.Phone))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Address: %s", p.FullAddress()))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Date of Birth: %s", p.DateOfBirth))
	pdf.Ln(15)
	
	// Complaint Details
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 8, "Complaint Details:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	
	complaint := `I am writing to formally complain about the discretionary commission arrangements that were in place when I purchased my motor vehicle finance. I believe I was not properly informed about the commission structure and how it may have affected the interest rate I was charged.

Key Issues:
1. I was not informed that the finance broker/dealer could set my interest rate within a range
2. I was not told that they would receive higher commission for setting a higher rate
3. The discretionary commission arrangement created a conflict of interest
4. I believe I may have been charged a higher interest rate than necessary

I understand that the Financial Conduct Authority has found these arrangements problematic and I request:
- Full disclosure of all commissions paid
- Calculation of any overcharged interest
- Appropriate compensation for financial detriment
- Explanation of how my interest rate was determined

This complaint relates to motor finance taken out approximately between 2007-2021 when discretionary commission arrangements were common in the industry.

I look forward to your investigation and resolution of this matter.`

	// Split text into lines that fit
	lines := pdf.SplitText(complaint, 180)
	for _, line := range lines {
		pdf.Cell(0, 6, line)
		pdf.Ln(6)
	}
	
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 6, "Yours sincerely,")
	pdf.Ln(8)
	pdf.Cell(0, 6, p.FullName())
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	
	return buf.Bytes(), nil
}

func GenerateComplaintPDFWithPassword(p person.FakePerson, password string) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	// Set password protection before adding content
	if password != "" {
		// In gofpdf, the SetProtection parameters are: (privFlag, userPass, ownerPass)
		// We want to require the password to open the document, so we set it as userPass
		pdf.SetProtection(gofpdf.CnProtectPrint|gofpdf.CnProtectCopy, password, "")
	}
	
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	
	// Header
	pdf.Cell(0, 10, "MOTOR FINANCE DISCRETIONARY COMMISSION COMPLAINT")
	pdf.Ln(15)
	
	// Date
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, fmt.Sprintf("Date: %s", time.Now().Format("02/01/2006")))
	pdf.Ln(10)
	
	// Personal Details
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 8, "Personal Details:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 6, fmt.Sprintf("Name: %s", p.FullName()))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Email: %s", p.Email))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Phone: %s", p.Phone))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Address: %s", p.FullAddress()))
	pdf.Ln(6)
	pdf.Cell(0, 6, fmt.Sprintf("Date of Birth: %s", p.DateOfBirth))
	pdf.Ln(15)
	
	// Complaint Details
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 8, "Complaint Details:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	
	complaint := `I am writing to formally complain about the discretionary commission arrangements that were in place when I purchased my motor vehicle finance. I believe I was not properly informed about the commission structure and how it may have affected the interest rate I was charged.

Key Issues:
1. I was not informed that the finance broker/dealer could set my interest rate within a range
2. I was not told that they would receive higher commission for setting a higher rate
3. The discretionary commission arrangement created a conflict of interest
4. I believe I may have been charged a higher interest rate than necessary

I understand that the Financial Conduct Authority has found these arrangements problematic and I request:
- Full disclosure of all commissions paid
- Calculation of any overcharged interest
- Appropriate compensation for financial detriment
- Explanation of how my interest rate was determined

This complaint relates to motor finance taken out approximately between 2007-2021 when discretionary commission arrangements were common in the industry.

I look forward to your investigation and resolution of this matter.`

	// Split text into lines that fit
	lines := pdf.SplitText(complaint, 180)
	for _, line := range lines {
		pdf.Cell(0, 6, line)
		pdf.Ln(6)
	}
	
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 6, "Yours sincerely,")
	pdf.Ln(8)
	pdf.Cell(0, 6, p.FullName())
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	
	return buf.Bytes(), nil
}

func PasswordProtect(pdfData []byte, password string) ([]byte, error) {
	// This function is now deprecated in favor of GenerateComplaintPDFWithPassword
	// For existing PDFs, we'll just return the original data since the new approach
	// embeds protection during generation
	return pdfData, nil
}
