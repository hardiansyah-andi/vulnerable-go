# vulnerable-go

This repository contains source codes for vulnerability testing. Codes in this repository contains bad practices, bugs, and vulnerabilities. These are intended for **demonstration purposes only**.

# how to


1. Run through the setup flow in the security tab to enable code scanning. Commit directly to main. 
2. Go to actions, and confirm that the analysis has run successfully.
3. Edit main.go and add this code snippet. Create a PR. Wait for validation. 


```go

func endsWith(x, y string) bool {
  index := strings.LastIndex(x, y)
  return index != -1 && index == len(x) - len(y)
}

```
