package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/maxroulstone/mf-complaint-generator/pkg/email"
	"github.com/maxroulstone/mf-complaint-generator/pkg/pdf"
	"github.com/maxroulstone/mf-complaint-generator/pkg/person"
	"github.com/maxroulstone/mf-complaint-generator/pkg/zip"
)

// Version information (set by build flags)
var version = "dev"

type Options struct {
	PasswordInEmail   bool
	GenerateChaser    bool
	NumberOfCases     int
	ChaserDaysDelay   int
	WorkerCount       int
}

type CaseResult struct {
	CaseNum      int
	Person       person.FakePerson
	ComplaintFile string
	PasswordFile  string
	ChaserFile    string
	PDFPassword   string
	ZipPassword   string
	Error         error
}

func main() {
	fmt.Printf("Motor Finance Complaint Generator v%s\n", version)
	fmt.Println("==========================================")
	fmt.Println()

	options := getUserOptions()
	
	// Auto-configure workers based on CPU cores and number of cases
	maxWorkers := runtime.NumCPU()
	if options.NumberOfCases < maxWorkers {
		options.WorkerCount = options.NumberOfCases
	} else {
		options.WorkerCount = maxWorkers
	}
	
	fmt.Printf("Generating %d complaint case(s)...\n", options.NumberOfCases)
	
	startTime := time.Now()
	results := generateCases(options)
	duration := time.Since(startTime)
	
	// Display results summary
	successCount := 0
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("ERROR Case %d: %v\n", result.CaseNum, result.Error)
		} else {
			successCount++
		}
	}
	
	fmt.Printf("\nCompleted: %d/%d cases in %v\n", successCount, options.NumberOfCases, duration)
	fmt.Printf("Performance: %.2f cases/second\n", float64(successCount)/duration.Seconds())
}

func getUserOptions() Options {
	reader := bufio.NewReader(os.Stdin)
	
	// Password option
	fmt.Print("Include passwords in the main complaint email? (y/n) [y]: ")
	passwordInput, _ := reader.ReadString('\n')
	passwordInput = strings.TrimSpace(passwordInput)
	if passwordInput == "" {
		passwordInput = "y"
	}
	passwordInEmail := strings.ToLower(passwordInput) == "y"
	
	// Chaser email option
	fmt.Print("Generate chaser emails? (y/n) [n]: ")
	chaserInput, _ := reader.ReadString('\n')
	chaserInput = strings.TrimSpace(chaserInput)
	if chaserInput == "" {
		chaserInput = "n"
	}
	generateChaser := strings.ToLower(chaserInput) == "y"
	
	var chaserDays int
	if generateChaser {
		fmt.Print("Days delay for chaser email [14]: ")
		daysInput, _ := reader.ReadString('\n')
		daysInput = strings.TrimSpace(daysInput)
		if daysInput == "" {
			chaserDays = 14
		} else {
			var err error
			chaserDays, err = strconv.Atoi(daysInput)
			if err != nil {
				chaserDays = 14
			}
		}
	}
	
	// Number of cases
	fmt.Print("How many complaint cases to generate [1]: ")
	casesInput, _ := reader.ReadString('\n')
	casesInput = strings.TrimSpace(casesInput)
	numberOfCases := 1
	if casesInput != "" {
		var err error
		numberOfCases, err = strconv.Atoi(casesInput)
		if err != nil || numberOfCases < 1 {
			numberOfCases = 1
		}
	}
	
	return Options{
		PasswordInEmail: passwordInEmail,
		GenerateChaser:  generateChaser,
		NumberOfCases:   numberOfCases,
		ChaserDaysDelay: chaserDays,
		WorkerCount:     0, // Will be set automatically
	}
}

func generateCases(options Options) []CaseResult {
	// Create channels for work distribution
	jobs := make(chan int, options.NumberOfCases)
	results := make(chan CaseResult, options.NumberOfCases)
	progress := make(chan int, options.NumberOfCases)
	
	// Start progress bar goroutine
	go showProgressBar(progress, options.NumberOfCases)
	
	// Start worker goroutines
	var wg sync.WaitGroup
	for w := 1; w <= options.WorkerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, progress, options, &wg)
	}
	
	// Send jobs
	for i := 1; i <= options.NumberOfCases; i++ {
		jobs <- i
	}
	close(jobs)
	
	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
		close(progress)
	}()
	
	// Collect results
	var allResults []CaseResult
	for result := range results {
		allResults = append(allResults, result)
	}
	
	// Sort results by case number for consistent output
	for i := 0; i < len(allResults)-1; i++ {
		for j := i + 1; j < len(allResults); j++ {
			if allResults[i].CaseNum > allResults[j].CaseNum {
				allResults[i], allResults[j] = allResults[j], allResults[i]
			}
		}
	}
	
	return allResults
}

func showProgressBar(progress <-chan int, total int) {
	completed := 0
	for range progress {
		completed++
		percent := float64(completed) / float64(total) * 100
		bar := strings.Repeat("=", int(percent/2)) + strings.Repeat(" ", 50-int(percent/2))
		fmt.Printf("\rProgress: [%s] %d/%d (%.1f%%)", bar, completed, total, percent)
	}
	fmt.Println() // New line after completion
}

func worker(id int, jobs <-chan int, results chan<- CaseResult, progress chan<- int, options Options, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for caseNum := range jobs {
		result := generateCase(caseNum, options)
		results <- result
		progress <- 1 // Signal progress
	}
}

func generateCase(caseNum int, options Options) CaseResult {
	result := CaseResult{CaseNum: caseNum}
	
	// Generate fake person
	p := person.Generate()
	result.Person = p
	
	// Generate passwords
	pdfPassword := zip.GeneratePassword(8)
	zipPassword := zip.GeneratePassword(8)
	result.PDFPassword = pdfPassword
	result.ZipPassword = zipPassword
	
	// Generate complaint PDF with password protection
	protectedPDF, err := pdf.GenerateComplaintPDFWithPassword(p, pdfPassword)
	if err != nil {
		result.Error = fmt.Errorf("failed to generate PDF: %v", err)
		return result
	}
	
	// Create attachments
	attachments := []zip.AttachmentData{
		{Filename: "motor_finance_complaint.pdf", Content: protectedPDF},
		{Filename: "supporting_documents.txt", Content: []byte("Additional documentation related to motor finance complaint and evidence supporting the claim for discretionary commission investigation.")},
	}
	
	// Create password-protected zip
	zipData, err := zip.CreatePasswordProtected(attachments, zipPassword)
	if err != nil {
		result.Error = fmt.Errorf("failed to create zip: %v", err)
		return result
	}
	
	// Generate main complaint email
	msgContent := email.CreateComplaintMsg(p, zipData, options.PasswordInEmail, pdfPassword, zipPassword)
	filename := fmt.Sprintf("case_%02d_complaint_%s_%s_%d.eml", caseNum, p.FirstName, p.LastName, time.Now().UnixNano())
	
	err = os.WriteFile(filename, []byte(msgContent), 0644)
	if err != nil {
		result.Error = fmt.Errorf("failed to write complaint EML file: %v", err)
		return result
	}
	result.ComplaintFile = filename
	
	// Generate separate password email if needed
	if !options.PasswordInEmail {
		passwordMsgContent := email.CreatePasswordMsg(p, pdfPassword, zipPassword)
		passwordFilename := fmt.Sprintf("case_%02d_passwords_%s_%s_%d.eml", caseNum, p.FirstName, p.LastName, time.Now().UnixNano())
		
		err = os.WriteFile(passwordFilename, []byte(passwordMsgContent), 0644)
		if err != nil {
			result.Error = fmt.Errorf("failed to write password EML file: %v", err)
			return result
		}
		result.PasswordFile = passwordFilename
	}
	
	// Generate chaser email if needed
	if options.GenerateChaser {
		chaserMsgContent := email.CreateChaserMsg(p, options.ChaserDaysDelay)
		chaserFilename := fmt.Sprintf("case_%02d_chaser_%s_%s_%d.eml", caseNum, p.FirstName, p.LastName, time.Now().UnixNano())
		
		err = os.WriteFile(chaserFilename, []byte(chaserMsgContent), 0644)
		if err != nil {
			result.Error = fmt.Errorf("failed to write chaser EML file: %v", err)
			return result
		}
		result.ChaserFile = chaserFilename
	}
	
	return result
}
