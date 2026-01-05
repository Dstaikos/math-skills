# Math Skills — Statistical Analysis Tool 📊

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![GitHub Pages](https://img.shields.io/badge/GitHub%20Pages-Live%20Demo-brightgreen?style=for-the-badge&logo=github)](https://dstaikos.github.io/math-skills/)
[![Build Status](https://img.shields.io/badge/Build-Passing-success?style=for-the-badge)](https://github.com/dstaikos/math-skills)

A **Go**  statistical analysis tool designed for processing numerical datasets. Computes essential statistical measures including **Average**, **Median**, **Variance**, and **Standard Deviation** with precision and efficiency. The results are rounded. 

---

**[🌐Try the Live Demo Interactive Calculator](https://dstaikos.github.io/math-skills/)**

---

## 📈 Statistical Measures

| Function |
|----------|
| **[Average (Mean)](https://en.wikipedia.org/wiki/Average)** |
| **[Median](https://en.wikipedia.org/wiki/Median)** | 
| **[Variance](https://en.wikipedia.org/wiki/Variance)** | 
| **[Standard Deviation](https://en.wikipedia.org/wiki/Standard_deviation)** |


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


### Running Tests
```bash
go test -v ./...
```


## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for complete details.

---

<div align="center">

</div>