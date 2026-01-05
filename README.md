# Math Skills — Statistical Analysis CLI Tool 📊

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![GitHub Pages](https://img.shields.io/badge/GitHub%20Pages-Live%20Demo-brightgreen?style=for-the-badge&logo=github)](https://dstaikos.github.io/math-skills/)
[![Build Status](https://img.shields.io/badge/Build-Passing-success?style=for-the-badge)](https://github.com/dstaikos/math-skills)

A high-performance **Go** command-line statistical analysis tool designed for processing numerical datasets. Computes essential statistical measures including **Average**, **Median**, **Variance**, and **Standard Deviation** with precision and efficiency. Features an interactive web-based calculator for real-time statistical analysis.

---

## 🌐 Live Demo

**[Try the Interactive Calculator →](https://dstaikos.github.io/math-skills/)**

---

## 📈 Statistical Functions

| Function | Description | Use Case |
|----------|-------------|----------|
| **Average (Mean)** | Arithmetic mean of all values | Central tendency analysis |
| **Median** | Middle value in sorted dataset | Robust central measure, outlier-resistant |
| **Variance** | Average squared deviation from mean | Data spread quantification |
| **Standard Deviation** | Square root of variance | Dispersion in original units |

## ⚡ Performance Features

- **Zero Dependencies**: Built entirely with Go standard library
- **Memory Efficient**: Optimized for large datasets
- **Fast Processing**: Single-pass algorithms where possible
- **Cross-Platform**: Runs on Windows, macOS, and Linux
- **Production Ready**: Comprehensive error handling and validation

## 🏗️ Architecture

```
math-skills/
├── cmd/
│   └── main.go           # CLI entry point
├── Operations/
│   ├── Average.go        # Mean calculation
│   ├── Median.go         # Median computation
│   ├── Variance.go       # Variance analysis
│   └── Sdeviation.go     # Standard deviation
├── docs/
│   ├── index.html        # Web calculator interface
│   └── app.js           # Statistical computation logic
└── go.mod               # Go module definition
```

## 🚀 Installation & Usage

### Prerequisites
- **Go 1.18+** installed on your system
- Basic understanding of statistical concepts

### Command Line Interface

1. **Clone the repository:**
```bash
git clone https://github.com/dstaikos/math-skills.git
cd math-skills
```

2. **Prepare your dataset** (one number per line):
```text
# data.txt
189.5
113.2
121.0
114.8
145.3
110.1
```

3. **Execute statistical analysis:**
```bash
cd cmd
go run main.go < ../data.txt
```

### Expected Output Format
```
Average: 132
Median: 117
Variance: 785
Standard Deviation: 28
```

## 🧮 Web Calculator

The included web interface provides:
- **Real-time computation** as you type
- **Input validation** with error handling
- **Responsive design** for mobile and desktop
- **Professional UI** with modern styling
- **GitHub Pages deployment** ready


### Running Tests
```bash
go test -v ./...
```


## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for complete details.

---

<div align="center">

**Built with ❤️ using Go**

[Report Bug](https://github.com/dstaikos/math-skills/issues) • [Request Feature](https://github.com/dstaikos/math-skills/issues) • [Documentation](https://github.com/dstaikos/math-skills/wiki)

</div>