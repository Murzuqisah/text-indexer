# 📌 text-indexer : A Fast & Scalable Text Indexing System

TextIndex is a **high-performance, parallelized, and memory-efficient** text indexing tool written in Go.It allows for **fast chunk-based indexing and lookup** of large text files using **SimHash**, multi-threading, and fuzzy search with **BK-Trees**.

## Table of Contents

- [📌 text-indexer : A Fast & Scalable Text Indexing System](#-text-indexer--a-fast--scalable-text-indexing-system)
- [🚀 Features](#-features)
- [🌐 Web Interface](#-web-interface)
  - [Theme Support](#theme-support)
- [🔐 Authentication & Security](#-authentication--security)
- [📂 Project Structure](#-project-structure)
- [📥 Installation](#-installation)
  - [1️⃣ Install Go (if not installed)](#1️⃣-install-go-if-not-installed)
  - [2️⃣ Clone the Repository](#2️⃣-clone-the-repository)
- [🔧 Usage](#-usage)
  - [CLI Usage](#cli-usage)
    - [Index a File](#index-a-file)
    - [Look Up a SimHash](#look-up-a-simhash)
  - [Web Server](#web-server)
- [⚙️ Configuration](#️-configuration)
- [🛠 Future Improvements](#-future-improvements)
- [📜 License](#-license)
- [🙌 Contributors](#-contributors)

## 🚀 Features

✔ **Parallel Processing**: Uses worker pools for fast chunk hashing with configurable worker count.

✔ **Memory-Mapped File Chunking**: Efficient file reads without loading everything into RAM.

✔ **BoltDB Storage**: Persistent indexing instead of keeping everything in memory.

✔ **BK-Tree for Fuzzy Search**: Fast lookup of near-matching SimHash values.

✔ **CLI Interface**: Simple command-line tool for indexing and searching.

✔ **Web Interface**: Modern responsive UI with light/dark mode for easy file management.

✔ **User Authentication**: Secure login, registration, and profile management.

✔ **Security Features**: Password hashing, CSRF protection, rate limiting, and secure sessions.

✔ **Smooth Animations**: Enhanced user experience with smooth scrolling and transitions.

✔ **Responsive Design**: Works on desktop, tablet, and mobile devices.

## 🌐 Web Interface

The web interface provides a user-friendly way to interact with the text indexer:

- **Dashboard**: Overview of indexing statistics and recent activity
- **Upload**: Drag-and-drop file uploading with real-time progress tracking
- **Search**: Advanced search capabilities with filtering options
- **Index Management**: Tools to rebuild, optimize, or clear indexes
- **Profile**: User profile management
- **Settings**: Application configuration including:

- Light/Dark/System theme selection
- Parallel worker configuration
- Chunk size optimization
- Security settings





### Theme Support

The interface supports three theme modes:

- **Light Mode**: Clean, bright interface for daytime use
- **Dark Mode**: Eye-friendly dark theme for low-light environments
- **System Mode**: Automatically matches your system preferences


## 🔐 Authentication & Security

The system implements comprehensive security features:

- **Secure Authentication**: Email and password-based authentication
- **Password Security**:

- bcrypt hashing with appropriate cost factor
- Password strength meter
- Complexity requirements (length, special chars, numbers, uppercase)



- **Session Management**:

- Secure HTTP-only cookies
- Configurable session timeout
- Protection against session hijacking



- **CSRF Protection**: Cross-Site Request Forgery prevention
- **Rate Limiting**: Protection against brute force attacks
- **Input Validation**: Comprehensive validation for all user inputs
- **Security Headers**: Modern security headers to prevent common attacks
- **Two-Factor Authentication**: Optional 2FA for enhanced security


## 📂 Project Structure

Here's a comprehensive overview of the project's structure:

```plaintext
📂 textindex/
├── 📂 cmd/                # CLI Command Handling
│   ├── main.go            # Main entry point
│   ├── index.go           # Command to create an index
│   ├── lookup.go          # Command to perform lookups
│   ├── server.go          # Command to start the web server
│   ├── utils.go           # CLI helper functions
│
├── 📂 indexer/            # Indexing System
│   ├── index.go           # Index storage using BoltDB
│   ├── workerpool.go      # Parallel chunk hashing worker pool
│   ├── bktree.go          # Fuzzy search with BK-Tree
│   ├── simhash.go         # SimHash computation functions
│   ├── options.go         # Configurable indexing options
│   ├── progress.go        # Progress tracking and reporting
│
├── 📂 parser/             # File Processing
│   ├── chunker.go         # Memory-mapped chunking
│   ├── reader.go          # File reading utilities
│
├── 📂 auth/               # Authentication System
│   ├── user.go            # User model and management
│   ├── password.go        # Password hashing and validation
│   ├── session.go         # Session management
│   ├── token.go           # Token generation and validation
│
├── 📂 web/                # Web Server Components
│   ├── 📂 handlers/       # HTTP Request Handlers
│   │   ├── auth.go        # Authentication handlers (login/register)
│   │   ├── index.go       # Index management handlers
│   │   ├── search.go      # Search handlers
│   │   ├── profile.go     # User profile handlers
│   │   ├── settings.go    # Settings handlers
│   │
│   ├── 📂 middleware/     # HTTP Middleware
│   │   ├── auth.go        # Authentication middleware
│   │   ├── security.go    # Security headers, CSRF, etc.
│   │   ├── ratelimit.go   # Rate limiting
│   │   ├── logger.go      # Request logging
│   │
│   ├── 📂 templates/      # HTML Templates
│   │   ├── index.html     # Main application template
│   │   ├── 📂 partials/   # Partial templates
│   │       ├── header.html    # Header partial
│   │       ├── footer.html    # Footer partial
│   │       ├── nav.html       # Navigation partial
│   │
│   ├── 📂 static/         # Static Assets
│   │   ├── 📂 css/        # CSS files
│   │   ├── 📂 js/         # JavaScript files
│   │   ├── 📂 img/        # Image files
│   │
│   ├── server.go          # Web server setup and configuration
│   ├── router.go          # HTTP route definitions
│
├── 📂 storage/            # Data Storage
│   ├── db.go              # Database connection and management
│   ├── user_store.go      # User data storage
│   ├── session_store.go   # Session data storage
│   ├── index_store.go     # Index data storage
│
├── 📂 config/             # Configuration
│   ├── config.go          # Configuration loading and validation
│   ├── defaults.go        # Default configuration values
│
├── 📂 tests/              # Unit Tests
│   ├── index_test.go      # Tests for indexing
│   ├── lookup_test.go     # Tests for lookup functionality
│   ├── auth_test.go       # Tests for authentication
│   ├── web_test.go        # Tests for web server
│
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies checksum
└── README.md              # Project documentation
```

This structure ensures a clean separation of concerns, making the codebase easy to navigate and maintain.

## 📥 Installation

### **1️⃣ Install Go (if not installed)**

Ensure you have Go installed on your system.🔗 [Download Go](https://golang.org/dl/)

### **2️⃣ Clone the Repository**

```shellscript
git clone https://github.com/Murzuqisah/text-indexer.git
cd text-indexer
```

Install dependencies

```shellscript
go mod tidy
```

## 🔧 Usage

### CLI Usage

#### Index a File

This command indexes a text file by splitting it into fixed-size chunks and storing the SimHash fingerprints.

```shellscript
textindex -c index -i input.txt -s 4096 -o index.idx
```

| Flag | Description
|-----|-----
| -c | Command (index or lookup)
| -i | Input file to index
| -s | Chunk size (e.g., 1KB, 4KB, etc.)
| -o | Output index file


#### Look Up a SimHash

Retrieve the byte offset of a given SimHash fingerprint. This command searches the index file for the specified SimHash value and returns the corresponding byte offset in the original text file.

```shellscript
textindex -c lookup -i index.idx -h <simhash_value>
```

| Flag | Description
|-----|-----
| -c | Command (index or lookup)
| -i | Input index file to search
| -h | SimHash value to look up


Replace `<simhash_value>` with the actual SimHash fingerprint you want to search for. This feature is useful for quickly locating specific text chunks in large files based on their SimHash values.

### Web Server

Start the web server to access the user interface:

```shellscript
textindex -c server -p 8080
```

| Flag | Description
|-----|-----
| -c | Command (server)
| -p | Port to run the server on (default: 8080)
| -d | Database path (default: ./data)


Once started, access the web interface by navigating to `http://localhost:8080` in your browser.

## ⚙️ Configuration

The application can be configured through environment variables or a configuration file:

```shellscript
# Server configuration
TEXTINDEX_PORT=8080
TEXTINDEX_HOST=localhost
TEXTINDEX_DB_PATH=./data

# Security settings
TEXTINDEX_SESSION_SECRET=your-secret-key
TEXTINDEX_SESSION_TIMEOUT=30m
TEXTINDEX_ENABLE_HTTPS=false
TEXTINDEX_CERT_FILE=./cert.pem
TEXTINDEX_KEY_FILE=./key.pem

# Indexing defaults
TEXTINDEX_DEFAULT_CHUNK_SIZE=4096
TEXTINDEX_DEFAULT_WORKERS=4
TEXTINDEX_MAX_WORKERS=16
```

Alternatively, create a `config.yaml` file in the root directory:

```yaml
server:
  port: 8080
  host: localhost
  db_path: ./data

security:
  session_secret: your-secret-key
  session_timeout: 30m
  enable_https: false
  cert_file: ./cert.pem
  key_file: ./key.pem

indexing:
  default_chunk_size: 4096
  default_workers: 4
  max_workers: 16
```

## 🛠 Future Improvements

- Distributed Indexing using multiple nodes
- Real-time indexing via WebSockets
- Compression for reduced storage overhead
- OAuth integration for third-party login
- Advanced analytics dashboard
- API for programmatic access
- Mobile application


## 📜 License

This project is licensed under [GPL-3.0 license](https://github.com/Murzuqisah/text-inder/blob/versions/LICENSE)

## 🙌 Contributors

Contributions are welcome! 🎉

~ Joel Amos – Developer

This **README** is structured for clarity, covering **installation, usage, performance, and implementation details**. 🚀