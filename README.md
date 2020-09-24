## HTML Form Generator

Simple HTML form generator.

### How to use

Make sure you prepare form field and it type into
colon duplet.

`go run main.go <name field>:<type field>`

Example:

`go run main.go name:text secret:password`

The output on screen:

```html
<form action=""  method="" enctype="">
  <input name="nama" type="text">
  <input name="secret" type="password">
  <input type="submit">
</form>
```

To make it a file as :

`go run main.go name:text secret:password > form.html`


Roadmap:
- [ ] Support CLI argument as input form parameter
- [ ] Adding CLI param support for action, method and enctype params.
- [ ] Support JSON string as input form parameter
[WIP]

## LICENSE

[BSL-1.0](./LICENSE)