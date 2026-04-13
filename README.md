# vandex-gf 🚀
**Next-Gen Recon Filtering - Faster, Smarter, and Concurrent.**

**Vandex-gf** is a high-performance security tool built in Go. It is a heavily enhanced and optimized evolution of the classic ```gf``` tool by the legend **@tomnomnom**. While it respects the original logic, it breaks the limitations of the past by introducing true concurrency and massive quality-of-life improvements.

------------------------------------------
# 🔥 Why Vandex-gf? 
- **Built in Go:** Extremely fast and handles large datasets efficiently.
- **Zero Line Limits:** Forget running the tool 10 times for 10 patterns. In ```vandex-gf```, you can chain as many patterns as you want in **one single command**.
- **The ```-all``` Beast:** Unlike the original, ```vandex-gf``` can run **every single pattern** in your ```~/.gf``` folder simultaneously using Go's concurrency
- **Smart Auto-Output:** It doesn't just print to the screen; it automatically categorizes and saves findings into separate ```.txt``` files named after the patterns.
- **User-Friendly:** Direct feedback with colors and clear error messages.

----------------------------------------------------------
# 📥 Installation
1. **Prerequisites**: Make sure you have Go installed and your ```.gf``` patterns folder in your home directory (```~/.gf```), You can use these patterns [GFPattern](https://github.com/coffinxp/GFpattren)
2. **Installation**:
```
go install github.com/0xvandex/vandex-gf@latest
```

-----------------------------------------------
# 🚀 Usage
**First, ensure your patterns are located in ~/.gf.**

## 1.Help
```
vandex-gf -h
```
This will display help for the tool. Here are all the switches it supports.
```                                                                                                                            
Usage:
  cat targets.txt | vandex-gf [pattern_names] [flags]

Main Flags:
  -list    : List available patterns
  -all     : Run all patterns at once (Multi-threaded)
  -h       : Show this help menu

Examples:
  cat urls.txt | vandex-gf xss sqli
  cat urls.txt | vandex-gf -all
  cat urls.txt | vandex-gf lfi
```

## 2. List Available Patterns
```
vandex-gf -list
```

## 3. Filter for Specific Patterns
You can pass one or multiple patterns as arguments:
```
cat urls.txt | vandex-gf xss sqli lfi
```

## 4. Run Everything (The Beast Mode)
Run all patterns in your ```.gf``` folder against the input using Go routines:
```
cat urls.txt | vandex-gf -all
```

------------------------------------
# 📂 Output 
The tool intelligently manages output:
  - If matches are found for a pattern (e.g., ```xss```), it saves them to ```xss.txt```.
  - If no matches are found, it cleans up and doesn't create empty files.

-----------------------------------------
# 🤝 Credits & Acknowledgments
This tool is an improved wrapper and a tribute to the original ```gf``` tool created by [tomnomnom](https://github.com/tomnomnom). We stand on the shoulders of giants to build better tools for the community.

-----------------------
# 👨‍💻 Author
**Mohamed Magdy / vandex**

GitHub: [@0xvandex](https://github.com/0xvandex)

Medium: [@0xvandex](https://medium.com/@0xvandex) 
