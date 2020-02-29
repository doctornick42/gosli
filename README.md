# gosli

The goal of this library is to make easier to work with slices in Go. It was inspired by .NET LINQ query language.

Since Go doesn't have generic methods, sometimes it's hard to choose a nice-looking way to find some element in a slice or how to filter it. It's pretty easy to use `for` loop for it, but we can face with a compromise between the DRY (Don't Repeat Yourself) principle and using `interface{}` (the thing that I try to avoid if possible as a big fan of strong typing).

Also, there was an idea to avoid using the `reflect` package, because performance really matters for the most of golang applications, and `reflect` might be slow.

So, after all, maybe a good way to deal with is to implement some code generator. And here comes gosli.

Gosli now has two ways for using:

- If you want to use gosli to work with slices of basic golang type, for example to find or filter something in a slice of strings, you can do it just out-of-the box. [Read more](docs/primitives.md) 
- Or you can generate a wrapper for any of your custom type to work with slices of it. In a few words, the generator can create for your `MyType` structure two wrappers: `MyTypeSlice` and `MyTypePSlice` (where `P` is for 'pointer') [Read more](docs/custom.md)
