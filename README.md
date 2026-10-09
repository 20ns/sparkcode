A small to-do app with a Go API and a TypeScript interface built with [Svelte](https://svelte.dev/).

The API lists and adds todos using in-memory storage. The frontend displays them and submits new entries. The OpenAPI specification is in `backend/`.

## Setup

If not already installed, please install the following:
1. Go ([install instructions](https://go.dev/doc/install))
2. NPM/NodeJS. We recommend using [NVM](https://github.com/nvm-sh/nvm)

We have tested this with Go 1.25 and Node 20. You may have issues if you try to use a different version.

A good starting point would be to look at the following files:
- `backend/main.go`
- `frontend/src/App.svelte`

## Running

Open two separate terminals - one for the Svelte app and one for the golang API.


### Golang API

1. In the first terminal, change to the backend directory (`cd backend`)
2. Run `go run main.go` to start the API server

This must be running for the frontend to work.

When you make a change, you must stop the server (`ctrl-c` in the terminal), and restart it with `go run main.go`.


### Svelte App

1. In the second terminal, change to the frontend directory (`cd frontend`)
2. Run `npm run dev` to start the Svelte app
3. If it doesn't open automatically, open [http://localhost:5173](http://localhost:5173) to view your website

Leave this running. It will automatically update when you make any changes.
