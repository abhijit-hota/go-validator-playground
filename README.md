<div align="center">
	<h1>
        <a href="https://apps.abhijithota.me/go-validator-playground">
            Go Validator Playground
        </a>
    </h1>
	<p>
		<b>
            Iterate & test your <a href="https://pkg.go.dev/github.com/go-playground/validator/v10"><code>go-playground/validator</code></a> struct tags
        </b>
        <br>
        <br>
        <b>
            <a href="https://apps.abhijithota.me/go-validator-playground">Go to the app!</a>
        </b>
        <br>
	</p>
</div>

![Screenshot of the webapp](.github/readme.png)

## What is this?

This is a simple web app that allows you to test your [`go-playground/validator`](https://pkg.go.dev/github.com/go-playground/validator/v10) struct tags. You can use it at https://apps.abhijithota.me/go-validator-playground.

## Purpose

[`go-playground/validator`](https://pkg.go.dev/github.com/go-playground/validator/v10) is probably the most popular Go validation library. It uses (or *exploits*, dare I say) struct tags to validate the particular struct fields. 

It's pretty flexible, but it's also a PITA when you misspell something or forget the syntax. This tool is not a silver bullet but it is here to help you find out what's wrong with the code more quickly.

## How does it work?

It uses [WebAssembly](https://webassembly.org/) to run the [`go-playground/validator`](https://pkg.go.dev/github.com/go-playground/validator/v10) library in the browser. 

- User enters the struct type definition.
- User enters the corresponding JSON data.
- These two (`string`) values are passed to the WebAssembly module.
- The struct definition is parsed and the struct type is created dynamically using some [`reflect`](https://golang.org/pkg/reflect/) magic.
- The JSON is unmarshalled into the struct.
- The actual validation is performed.
- Errors and invalid fields/tags are returned back to the JS code.
  

## Todo

- [ ] Tests
- [ ] Add highlights in the `textarea` for invalid fields/tags
- [ ] Syntax highlighting maybe?
- [ ] Minimum responsiveness
- [ ] Write blog post — "*A primer on Go reflection*"
- [ ] Maybe buy a domain if enough people use it

## A word by the author

I mostly coded this as a way to explore WebAssembly and out of frustration with not knowing the unpredictable behaviour of struct tags.

It is mostly a proof of concept. I've tested in as much as possible *manually* but it might still fail, in which case, please reload the page and try again.

Thanks to https://gotemplate.io/ for the inspiration.

