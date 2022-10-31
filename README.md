# 🩳 Shorts

Simple CLI fuzzy search & bookmarks for `*Unix`

[![asciicast](https://asciinema.org/a/Qu7Ouu0PhIzfhP5RC01JK2txM.svg)](https://asciinema.org/a/Qu7Ouu0PhIzfhP5RC01JK2txM)

Store your favorite sites
---
```json
// `$HOME/.shorts/paths.json`
{
  "duck": "https://duckduckgo.com",
  "social": {
    "github": "https://github.com",
    "youtube": "https://youtube.com",
    "linkedin": "https://linkedin.com",
    "hacker news": "https://news.ycombinator.com/"
  }
}
```
Configure how it all works
---
```json
// `$HOME/.shorts/config.json`
{
  // command and args to open sites   
  "BrowserCommand": ["xdg-open"],
  // dedicated bin directory   
  "BinaryDirName": "bin",
  // dedicated CD dir
  "DirectoriesDirName": "directories"
}
```

Install
---

1. download and build
2. `alias shorts=$(~/bin/shorts)`
3. 🚀