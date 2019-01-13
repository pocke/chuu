Chuu (注)
===

注 convert `(注:something)` to markdown style notes.

Usage
---

```
# Both OK
$ chuu in.md out.md
$ chuu in.md > out.md
$ chuu < in.md > out.md
```


`in.md`

```
Foo bar(注:foobar)
Bar baz(注:barbaz)
```

`out.md`

```
Foo bar[^1]
Bar baz[^2]

[^1]: foobar
[^2]: barbaz
```
