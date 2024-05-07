### Live Editor
Live editor is a web app that renders webpages in real-time(sort of). It's basically CodePen with less features.

### Supported Languages
- HTML
- CSS
- Sass
- JavaScript
- Typescript

[CodeMirror](https://codemirror.net/docs/guide/) is doing the heavy lifting, code editing, syntax highlighting, etc.

### How to run
- Clone this repo
#### With docker
- If you have docker insalled run the command `docker compose up -d` in your terminal
#### Without docker
- Make a duplicate copy of the `.env.example` file in the `web` folder, save the duplicate as `.env`
- Provide a value for the environment variable `VITE_API`. This is the url of the backend, assuming you didn't
change the port it should be running on `localhost:8000`.
- The api requires the environment variable `API` to be set, you can do this by running `export API="localhost:8000"`. This
should be the url where the backend is running on.
- Run `npm run dev` from the `web` folder to get the frontend running
- Run `go run main.go` from the `api` folder to get the backend running

### Known issues
- Refreshing the page after writing Sass will return the compiled sass
- Refreshing the page after writing Ts will return the compiled typescript
- Resizing on the y axis downwards is janky, I've narrowed this down to the iframe used to render web pages.

#### Why
Curiosity, Boredom.