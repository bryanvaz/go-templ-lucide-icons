
# Lucide Icons for Templ

<div align="center">

Lucide icon library for Golang Templ applications.

<p align="center">
  <a href="https://pkg.go.dev/github.com/bryanvaz/go-templ-lucide-icons">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7">
  </a>
  <a href="https://goreportcard.com/report/github.com/bryanvaz/go-templ-lucide-icons">
    <img src="https://goreportcard.com/badge/github.com/bryanvaz/go-templ-lucide-icons">
  </a>
  <a href="https://github.com/bryanvaz/go-templ-lucide-icons/actions/workflows/test.yml">
    <img src="https://github.com/bryanvaz/go-templ-lucide-icons/actions/workflows/test.yml/badge.svg?branch=main">
  </a>
</p>

[About Lucide](https://lucide.dev/guide/)
.
[Icons](https://lucide.dev/icons/)
.
[About Templ](https://templ.guide/)

</div>

## Installation

`go get -u github.com/bryanvaz/go-templ-lucide-icons`

## Quick start

```templ
package main

import "github.com/bryanvaz/go-templ-lucide-icons"

templ Index() {
  <!DOCTYPE html>
  <html lang="en">
    <head>
      <meta charset="UTF-8"/>
      <title>Lucide Test Server</title>
    </head>
    <body>
      <h1>Pen Icon</h1>
      <div>@icons.Pen()</div>
    </body>
  </html>
}
```

## Usage

All icons are available as individual templ components in the `icons` package
so only the icons referenced by your project will be included in your build. 

Each icon can be used as a normal templ component which will render an inline
SVG element. Attributes can be passed to the icon component as `templ.Attributes`, which 
will be attached to the root SVG element.

All attribute values must be strings to be properly processed by templ.

**Example**

```templ
package templates

import "github.com/bryanvaz/go-templ-lucide-icons"

templ RenderAPen() {
  <div>
    @icons.Pen(templ.Attributes{
      "size":         "24",
      "color":        "red",
      "stroke-width": "4",
    })
  </div>
}
```

## Props

|  name                   |   type    |  default     |
| ----------------------- | --------- | ------------ |
| `size`                  | *number*  | 24           |
| `color`                 | *string*  | currentColor |
| `stroke-width`          | *number*  | 2            |
| `absolute-stroke-width` | *boolean* | false        |
| `default-class`         | *string*  | lucide-icon  |

### Adding classes to the icon

To add classes to the icon, pass a `templ.Attributes` object with the `class` key
to the component.

```templ
<div>
  @icons.Pen(templ.Attributes{ "class": "my-icon" })
</div>
```

### Passing children to the icon

To pass children to the icon, pass children how would normally be passed to a
templ component.

```templ
<div>
  @icons.Pen() {
    <text>Hello, world!</text>
  }
</div>
```

### Using icon aliases

The `icons` package also exports all the aliases for an icon if relevant.
These are usually legacy names for an icon. For example, the `Pen` icon is
is aliased to `Edit2` in the `icons` package (the previous name for `pen` was
`edit-2`.)

This package does not export prefixed or suffixed aliases as golang automatically
namespaces imports.

These aliases are also mentioned in the godoc.

## Issues and Contributions

This is a release repository for templ implementations of `bryanvaz/go-lucide`.

Please file issues and pull requests against the upstream repository at 
[bryanvaz/go-lucide](https://github.com/bryanvaz/go-lucide).

## License

Lucide is licensed under the ISC license. See [LICENSE](https://lucide.dev/license).

This wrapper is licensed under the MIT license.

## Sponsors

This library is currently supported on a best-effort basis. 
If you would like to sponsor this project, please reach out via Gihub.

