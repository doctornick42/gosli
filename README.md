# gosli

### **Tedious intro**:
The goal of this library is to make easier to work with slices in Go. It's inspired by .NET LINQ query language.

Since Go doesn't have generic methods, we have to write all these `for` loops for every type we work with. And this routine takes a huge amount of time and it's not very fun to wright almost the same loops again and again. 

gosli doesn't use `reflect` package at all because performance really matters for the most of golang applications, and `reflect` is considered to be slow.

So, after all, the good way to avoid using reflection is to implement some code generator. And here comes gosli.

Gosli now has two ways for using:

- If you want to use gosli to work with slices of basic golang type, you can do it just out-of-the box.
[Read about it here](docs/primitives.md) 
- Or you can generate a wrapper for any of your custom type to work with slices of it. [Read about it here](docs/custom.md)
