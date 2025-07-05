# README

## About

This is a [WAILS](https://wails.io) project template that uses HTMX for interactivity, Tailwind CS for styling, daisyUI for CSS components, and the Golang standard library for routing and HMTL templating. See the links below for official documentation on the libraries used.

- Interactivity: [HTMX](https://htmx.org)
- CSS Styling: [Tailwind CSS](https://tailwindcss.com/)
- CSS Components: [daisyUI](https://daisyui.com/docs/intro/)
- HTML Templating: [Golang html/template](https://pkg.go.dev/html/template)
- Routing: [Golang net/http ServeMux](https://pkg.go.dev/net/http#ServeMux)

## Setup Instructions

Start a new wails project using the template

```shell
wails init -n new-project-name -t https://github.com/ltcovalt/wails-htmx-tailwind-daisyui-template
```

Open the new project directory and start wails in live developmen mode or build the redistributable by following the instructions below.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will do the following things:

- Build your application and run it
- Bind your Go code to the frontend so it can be called from JavaScript
- Use [Vite](https://vitejs.dev/) to watch for modifications in your Go files and rebuild/re-run on change
- Set up a [webserver](http://localhost:34115/) to serve your application, allowing your favorite browser and extrensions to be used
- See the [official docs](https://wails.io/docs/gettingstarted/development) for details.

## Building

To build a redistributable, production mode package, use `wails build`.
