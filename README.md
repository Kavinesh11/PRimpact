# PR Impact Analyzer

A powerful tool that analyzes the potential impact of pull requests by examining code changes, affected modules, and dependencies. It provides insights for both technical and non-technical stakeholders.

---

## 🚀 Features

- 🔍 Analyzes files and functions changed in a PR  
- 🧩 Identifies affected modules and dependencies  
- ⚠️ Flags potentially risky changes  
- 👥 Suggests reviewers based on code ownership  
- 📝 Generates plain-English summaries for non-technical stakeholders  

---

## 🛠️ Installation

### ✅ Using Go

```bash
go install github.com/pr-impact-analyzer@latest
```

# Build the Docker image
docker build -t pr-analyzer .

# Run the container
docker run -it pr-analyzer


# Analyze a specific PR
pr-impact-analyzer analyze --owner <repo-owner> --repo <repo-name> --pr <pr-number>

# Get suggested reviewers
pr-impact-analyzer reviewers --owner <repo-owner> --repo <repo-name> --pr <pr-number>

# Generate summary for stakeholders
pr-impact-analyzer summarize --owner <repo-owner> --repo <repo-name> --pr <pr-number>


# Run the analyzer in Docker
docker run -it pr-analyzer

# The interactive menu will guide you through the available options:
 1. Test identifyModule
 2. Test isRiskyChange
 3.  3. Exit


GITHUB_TOKEN=your_github_token
