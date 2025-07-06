# ParamChunker

A simple and flexible CLI tool written in Go that reads a wordlist and outputs it as a list of HTTP parameters in the format:

```
param1=XNLV1&param2=XNLV2&...
```

## 🚀 Features

- ✅ Read input in **one of three modes**:
  - Piped from stdin
  - From a file (`-f`)
  - Pasted interactively (`-i`)
- ✅ Output formatted parameters:
  - All at once
  - In chunks (`-n`)
  - Optionally with interactive paging (`-oi`)
- ✅ Clean error handling and user-friendly prompts
- ⚡ Efficient for files up to 10MB

---

## 📦 Installation

```bash
git clone https://github.com/yourusername/wordlist-paramchunker.git
cd wordlist-paramchunker
go build -o paramchunker
```

---

## 🧠 Usage

### Basic

```bash
cat wordlist.txt | ./paramchunker
```

```bash
./paramchunker -f wordlist.txt
```

```bash
./paramchunker -i
```

> Paste your content, then press `Ctrl+D` to finish.

---

### Output Options

| Flag            | Description                                           |
|-----------------|-------------------------------------------------------|
| `-oi`           | Output each chunk interactively (press Enter to show next) |
| `-n <size>`     | Specify chunk size (e.g. 20 params per chunk)         |

> If `-n` is not set, the output is printed all at once regardless of `-oi`.

---

### Examples

#### 🔸 Output All at Once
```bash
./paramchunker -f wordlist.txt
```

#### 🔸 Output in Chunks of 10, Interactive
```bash
./paramchunker -f wordlist.txt -n 10 -oi
```

#### 🔸 Paste Mode, Output in Chunks
```bash
./paramchunker -i -n 5
```

---

## 📥 Input Format

Your input (file, stdin, or pasted) should be a list of parameter names, one per line:

```
username
password
token
session
```

---

## 🧾 Output Format

Transforms input to:

```
username=XNLV1&password=XNLV2&token=XNLV3&session=XNLV4
```

---

## ⚠️ Notes

- You must provide **exactly one** input method: stdin, `-f`, or `-i`.
- Maximum file size is tested up to 10MB.
- Empty lines are trimmed automatically.

---

## 📄 License

MIT License

---

## 🤝 Contributions

PRs are welcome! If you find a bug or want to add a feature, feel free to open an issue or pull request.
