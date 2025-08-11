# Motor Finance Complaint Generator

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/maxroulstone/mf-complaint-generator)](https://goreportcard.com/report/github.com/maxroulstone/mf-complaint-generator)

A high-performance Go application that generates realistic motor finance discretionary commission complaint emails with password-protected PDF attachments. Perfect for testing email systems, compliance training, or generating synthetic data for development purposes.

## Features

- **Modular Architecture**: Clean package structure following Go best practices
- **Realistic Data Generation**: Creates authentic UK customer data with proper postcodes
- **Professional PDFs**: Motor finance complaint documents with proper formatting
- **Dual Password Protection**: Password-protected PDFs inside password-protected ZIP files
- **Multiple Email Types**: Complaint, password, and chaser emails
- **High Performance**: 500+ cases/second with parallel processing
- **Auto-Configuration**: Intelligent worker allocation based on CPU cores
- **Progress Tracking**: Real-time progress bars and performance metrics
- **Interactive CLI**: User-friendly prompts with sensible defaults

## Project Structure

```
mf-complaint-generator/
├── cmd/
│   └── generator/          # Main application
│       └── main.go
├── pkg/
│   ├── person/            # Fake person generation
│   ├── pdf/               # PDF creation and protection
│   ├── zip/               # ZIP file creation with passwords
│   └── email/             # Email generation (complaint, password, chaser)
├── go.mod
└── README.md
```

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/maxroulstone/mf-complaint-generator.git
cd mf-complaint-generator

# Install dependencies
go mod tidy

# Build the application
make build

# Or run directly
go run cmd/generator/main.go
```

### Basic Usage

```bash
# Generate a single complaint case
go run cmd/generator/main.go

# Use Makefile shortcuts
make demo      # Generate 3 demo cases
make benchmark # Test with 50 cases
make stress    # Stress test with 100 cases
```

## Interactive Configuration

The application provides user-friendly prompts for:

1. **Password Location**: Include passwords in main email or send separately
2. **Chaser Emails**: Generate follow-up emails after specified days
3. **Batch Size**: Number of complaint cases to generate

```
Motor Finance Complaint Generator
=================================

Include passwords in the main complaint email? (y/n) [y]: n
Generate chaser emails? (y/n) [n]: y
Days delay for chaser email [14]: 7
How many complaint cases to generate [1]: 10

Generating 10 complaint cases using 8 workers...
Progress: [==================================================] 10/10 (100.0%)
Completed: 10/10 cases in 45.2ms
Performance: 221.24 cases/second
```

```go
config := EmailConfig{
    SMTPHost:  "smtp.gmail.com",           // Your SMTP server
    SMTPPort:  "587",                      // SMTP port
    Username:  "your-email@gmail.com",     // Your email
    Password:  "your-app-password",        // Your password/app password
    FromEmail: "your-email@gmail.com",     // From email
    FromName:  "Your Name",                // From name
}
```

### Gmail Setup

For Gmail, you need to:

1. Enable 2-factor authentication
2. Generate an app password: [Google App Passwords](https://myaccount.google.com/apppasswords)
3. Use the app password instead of your regular password

## Usage

### Basic Usage

```bash
go run main.go
```

### Demo Mode

If you haven't configured SMTP credentials, the program runs in demo mode and creates a test zip file locally.

## Code Structure

- `EmailConfig`: SMTP configuration
- `EmailData`: Email content and attachments
- `AttachmentData`: File data for attachments
- `EmailGenerator`: Main class handling email generation and sending

## Security Features

- Random password generation using cryptographically secure random numbers
- Password-protected zip files with standard encryption
- Passwords included in email body for recipient access

## Example Output

```
Email sent successfully to: recipient@example.com
Zip password: aB3$xY9#mK2P
```

## Customization

You can customize:

- Password length (currently 12 characters)
- Character set for password generation
- Email templates
- Attachment file types and content
- SMTP configuration

## Generated Files

Each case generates up to 3 files:

```
case_01_complaint_John_Smith_1234567890.eml    # Main complaint with ZIP attachment
case_01_passwords_John_Smith_1234567890.eml    # Separate passwords (optional)
case_01_chaser_John_Smith_1234567890.eml       # Follow-up email (optional)
```

### EML Format Compatibility

The generated `.eml` files use standard RFC 2822 email format with MIME multipart structure, making them compatible with:

- **Windows**: Outlook, Windows Mail, Thunderbird
- **macOS**: Mail.app, Outlook, Thunderbird
- **Linux**: Thunderbird, Evolution, KMail
- **Web**: Most webmail clients can import EML files
- **Email servers**: Can be directly imported into most email systems

The files contain proper email headers (From, To, Subject, Date) and MIME boundaries for attachments, ensuring broad compatibility across email clients and platforms.

## Architecture

```
mf-complaint-generator/
├── cmd/
│   └── generator/          # Main CLI application
├── pkg/
│   ├── person/            # UK person data generation
│   ├── pdf/               # PDF creation & password protection
│   ├── zip/               # Password-protected ZIP creation
│   └── email/             # Email generation (3 types)
├── go.mod                 # Go module definition
├── Makefile              # Build & demo targets
└── README.md
```

## Performance

- **Speed**: 500+ cases/second on modern hardware
- **Parallel Processing**: Auto-configures workers based on CPU cores
- **Memory Efficient**: Streaming file generation
- **Scalable**: Tested with 1000+ concurrent cases

## Security Features

- **PDF Password Protection**: Each PDF protected with unique 8-character password
- **ZIP Encryption**: Additional layer of protection for email attachments
- **Unique Passwords**: Every case uses different randomly generated passwords
- **No Persistent Storage**: Passwords only exist in generated email content

## Use Cases

- **Email System Testing**: Generate realistic test data for email processing systems
- **Compliance Training**: Create training scenarios for financial services
- **Development Testing**: Synthetic data for complaint management systems
- **Performance Testing**: High-volume email generation for load testing
- **Security Testing**: Test password-protected attachment handling

## Makefile Targets

```bash
make build          # Build binary to bin/complaint-generator
make demo           # Generate 3 demo cases with all features
make benchmark      # Performance test with 50 cases
make stress         # Stress test with 100 cases
make clean          # Remove generated files
make test           # Run unit tests
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This tool is designed for testing and development purposes. Ensure compliance with your organization's data protection and privacy policies when using generated data.

## Acknowledgments

- Built with Go's excellent standard library
- PDF generation powered by gofpdf
- Password protection via pdfcpu
- ZIP encryption using alexmullins/zip

---

**Star this repository if you find it helpful!**
