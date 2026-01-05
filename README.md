# Math Skills — CLI & Demo Calculator 🚀

A small, fast **Go** command-line tool that reads numbers (one per line) and computes basic statistics: **Average**, **Median**, **Variance**, and **Standard Deviation**. Includes a browser demo you can publish with **GitHub Pages**.

---

[![Try The Live Demo](https://img.shields.io/badge/demo-live-brightgreen)](https://<Dstaikos>.github.io/math-skills/)

## ✅ Features

- 🧮 Compute: **Average**, **Median**, **Variance**, **Standard Deviation**
- ⚡ Fast, single-file CLI written in Go
- 🌐 Static browser demo (in `docs/`) for easy sharing via GitHub Pages

---

## 🔧 Quick Start

### Prerequisites
- Go 1.18+ installed

### Run locally (CLI)

1. Change into the `cmd` directory:

```bash
cd cmd
```

2. Ensure your input file (one number per line) is available, e.g. `input.txt`:

```text
1.2
3.4
5
2
```

3. Run the program:

```bash
go run main.go < input.txt >
```

> The program reads from stdin (one number per line) and prints the computed statistics.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

