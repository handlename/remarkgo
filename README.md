# remarkgo

remarkgo is a simple slide viewer for [remark](https://github.com/gnab/remark).

## Install

```
$ go get github.com/handlename/remarkgo/cmd/remarkgo
```

## Usage

Place your markdownfile as `index.md` and execute remarkgo.

```
$ remarkgo
```

Or set option `-s path/to/markdownfile`.

```
$ remarkgo -s path/to/markdownfile
```

Then, you can look slides at `http://localhost:8080`.

Some options are supported.
See help by `remakgo -help`.

## Shortcut keys

Those are availabled in browser.

- `r`: Reload slides from markdown file.

## Author

NAGATA Hiroaki ([@handlename](https://github.com/handlename))

## License

MIT
