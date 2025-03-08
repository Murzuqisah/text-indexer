# 📌 text-indexer : A Fast & Scalable Text Indexing System

TextIndex is a **high-performance, parallelized, and memory-efficient** text indexing tool written in Go.  
It allows for **fast chunk-based indexing and lookup** of large text files using **SimHash**, multi-threading, and fuzzy search with **BK-Trees**.

## Table of Contents

- [📌 TextIndex: A Fast & Scalable Text Indexing System](#-textindex-a-fast--scalable-text-indexing-system)
- [🚀 Features](#-features)
- [📂 Project Structure](#-project-structure)
- [📥 Installation](#-installation)
    - [1️⃣ Install Go (if not installed)](#1️⃣-install-go-if-not-installed)
    - [2️⃣ Clone the Repository](#2️⃣-clone-the-repository)
- [🔧 Usage](#-usage)
    - [Index a File](#index-a-file)
    - [Look Up a SimHash](#look-up-a-simhash)
- [🛠 Future Improvements](#-future-improvements)
- [📜 License](#-license)
- [🙌 Contributors](#-contributors)

## 🚀 Features

✔ **Parallel Processing**: Uses worker pools for fast chunk hashing.  
✔ **Memory-Mapped File Chunking**: Efficient file reads without loading everything into RAM.  
✔ **BoltDB Storage**: Persistent indexing instead of keeping everything in memory.  
✔ **BK-Tree for Fuzzy Search**: Fast lookup of near-matching SimHash values.  
✔ **CLI Interface**: Simple command-line tool for indexing and searching.

## 📂 Project Structure

Here's a quick overview of the project's structure:

``` txt
📂 textindex/
├── 📂 cmd/                # CLI Command Handling
│   ├── main.go            # Main entry point
│   ├── index.go           # Command to create an index
│   ├── lookup.go          # Command to perform lookups
│   ├── utils.go           # CLI helper functions
│
├── 📂 indexer/            # Indexing System
│   ├── index.go           # Index storage using BoltDB
│   ├── workerpool.go      # Parallel chunk hashing worker pool
│   ├── bktree.go          # Fuzzy search with BK-Tree
│   ├── simhash.go         # SimHash computation functions
│
├── 📂 parser/             # File Processing
│   ├── chunker.go         # Memory-mapped chunking
│   ├── reader.go          # File reading utilities
│
├── 📂 tests/              # Unit Tests
│   ├── index_test.go      # Tests for indexing
│   ├── lookup_test.go     # Tests for lookup functionality
│
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies checksum
└── README.md              # Project documentation
```

This structure ensures a clean separation of concerns, making the codebase easy to navigate and maintain.

## 📥 Installation

### **1️⃣ Install Go (if not installed)**

Ensure you have Go installed on your system.  
🔗 [Download Go](https://golang.org/dl/)

### **2️⃣ Clone the Repository**

```sh

git clone https://github.com/Murzuqisah/text-indexer.git
cd text-indexer
```

Install dependencies

```sh

go mod tidy
```

## 🔧 Usage

### Index a File

This command indexes a text file by splitting it into fixed-size chunks and storing the SimHash fingerprints.

```sh
textindex -c index -i input.txt -s 4096 -o index.idx
```

| Flag | Description |
|------|-------------|
| -c   | Command (index or lookup) |
| -i   | Input file to index |
| -s   | Chunk size (e.g., 1KB, 4KB, etc.) |
| -o   | Output index file |

### Look Up a SimHash

Retrieve the byte offset of a given SimHash fingerprint. This command searches the index file for the specified SimHash value and returns the corresponding byte offset in the original text file.

```sh
textindex -c lookup -i index.idx -h <simhash_value>
```

| Flag | Description |
|------|-------------|
| -c   | Command (index or lookup) |
| -i   | Input index file to search |
| -h   | SimHash value to look up |

Replace `<simhash_value>` with the actual SimHash fingerprint you want to search for. This feature is useful for quickly locating specific text chunks in large files based on their SimHash values.

## 🛠 Future Improvements

- Distributed Indexing using multiple nodes
- Real-time indexing via WebSockets
- Compression for reduced storage overhead

## 📜 License

This project is licensed under [GPL-3.0 license](https://github.com/Murzuqisah/text-inder/blob/versions/LICENSE)

## 🙌 Contributors

Contributions are welcome! 🎉

~ Joel Amos – Developer

This **README** is structured for clarity, covering **installation, usage, performance, and implementation details**. 🚀  
Would you like any refinements? 😊