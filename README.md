# GitHub User Activity CLI
A small command-line tool to display the latest public activity of a GitHub user.
This project is intentionally kept simple: it fetches data from the GitHub public events API and 
renders it into human-readable activity lines.

## Features
- Fetches latest public events of a GitHub user
- Support multiple GitHub event types (Push, PR, Issues, Fork, Watch, etc.)
- Human-readable output
- Minimal dependencies (standard library only)

## Usage
```
github-activity <username>
```
Example:
```github-activity johndoe
```
Output:
```
johndoe activity:
1 - pushed a commit to johndoe/github-activity at 22/01/2026 17:18
```

## Disclaimer
This project is for learning and experimentation purposes for https://roadmap.sh/projects/github-user-activity

## License
MIT
