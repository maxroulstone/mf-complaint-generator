# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed

- **CRITICAL**: Fixed PDF password protection - PDFs are now properly encrypted with the generated passwords
- Resolved issue where `api.EncryptFile` was being called with `nil` configuration

### Changed

- Enhanced CI/CD pipeline with automatic release creation
- Auto-increment patch version on successful tests in main branch
- Automated build and asset upload for new releases

### Added

- Initial release of Motor Finance Complaint Generator
- Parallel processing with goroutines for high performance
- Interactive CLI with user-friendly prompts
- Auto-configuration of worker count based on CPU cores
- Real-time progress bar with completion percentage
- Support for three email types: complaint, password, and chaser
- Dual password protection (PDF + ZIP)
- Realistic UK person data generation
- Professional PDF complaint document generation
- Comprehensive Makefile with demo, benchmark, and stress-test targets
- Cross-platform builds (Windows, macOS, Linux)
- MIT license for open source distribution

### Performance

- Achieves 500+ cases/second on modern hardware
- Efficient memory usage with streaming file generation
- Scales automatically based on available CPU cores

## [1.0.0] - 2025-08-10

### Added

- Initial public release
- Complete motor finance complaint generation system
- Password-protected PDF and ZIP file creation
- Multi-threaded parallel processing
- Interactive command-line interface
- Comprehensive documentation and examples
