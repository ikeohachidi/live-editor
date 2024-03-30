### Live Editor
Live editor is a web app that renders webpages in real-time(sort of). It's basically CodePen with less features.

### Supported Languages
- HTML
- CSS
- Sass
- JavaScript
- Typescript

[CodeMirror](https://codemirror.net/docs/guide/) is doing the heavy lifting, code editing, syntax highlighting, etc.

### Known issues
- Refreshing the page after writing Sass will return the compiled sass
- Refreshing the page after writing Ts will return the compiled typescript
- Resizing on the y axis downwards is janky, I've narrowed this down to the iframe used to render web pages.

#### Why
Curiosity, Boredom.