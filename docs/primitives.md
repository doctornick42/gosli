### **Basic types**

We need to get gosli. All we need to do is run:
```
go get github.com/doctornick42/gosli
```
For all basic types we got a structure representing a slice of it. For example, slice of `int64` is wrapped
into `Int64Slice` structure. All methods are the same for all primitive types, so during this page
I'll show a way of interacting with `Int64Slice` but the idea will be exactly the same for all other types.

---
## **Methods**


* ### First
    Returns first item of a slice that is passed through a filter.

    If an item wasn't found, the method returns an error.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A == 2
    }
    res, err := FakeTypeSlice(sl).First(filter)

    //res = &FakeType{
    //    A: 2,
    //    B: "two",
    //} 
    ```

* ### FirstOrDefault
    Returns first item of a slice that is passed through a filter.

    If an item wasn't found, the result is nil.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A == 2
    }
    res, err := FakeTypeSlice(sl).FirstOrDefault(filter)

    //res = &FakeType{
    //    A: 2,
    //    B: "two",
    //} 
    ```

* ### Where
    Returns all items of a slice that is passed through a filter.

    If items weren't found, the result is empty slice.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A >= 2
    }
    res, err := FakeTypeSlice(sl).Where(filter)

    //res = []*FakeType{
	//    &FakeType{
    //        A: 2,
    //        B: "two",
    //    },
    //    &FakeType{
    //        A: 3,
    //        B: "three",
    //    },
    //} 
    ```

* ### Select
    Applies a function to every item of a slice and returns slice of results.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    f := func(t *FakeType) interface{} {
        return struct {
            Msg string
        }{
            Msg: t.B,
        }
    }
    res, err := FakeTypeSlice(sl).Select(f)

    //res = []struct {
	//    Msg string
	//}{
    //    {
    //        Msg: "one",
    //    },
    //    {
    //        Msg: "two",
    //    },
    //    {
    //        Msg: "three",
    //    },
	//} 
    ```

* ### Page
    Returns paginated slice according to given `number` (number of selected page) and `perPage` 
    (items per a page). `number` parameter should start with 1 (not 0).
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    res, err := FakeTypeSlice(sl).PerPage(1, 2)

    //res = []*FakeType
    //    &FakeType{
    //        A: 1,
    //        B: "one",
	//    },
    //    &FakeType{
    //        A: 2,
    //        B: "two",
    //    },
	//} 
    ```

* ### Any
    Returns `true` if any item of the slice is passed through a filter.

    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A == 2
    }
    res, err := FakeTypeSlice(sl).Any(filter)

    //res = true
    ```

* ### Contains
    Returns `true` if a slice contains at least one item that is equal to the desired one.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }

    el := &FakeType{
        A: 2,
        B: "two",
    }
    
    res, err := FakeTypeSlice(sl).Contains(el)

    //res = true
    ```

* ### GetUnion
    Returns a slice that contains items that are contained in both given slices.
    
    ```go
    sl1 := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }

    sl2 := []*FakeType{
	    &FakeType{
            A: 2,
            B: "two",
	    },
        &FakeType{
            A: 3,
            B: "three",
        },
        &FakeType{
            A: 4,
            B: "four",
        },
    }

    res, err := FakeTypeSlice(sl1).GetUnion(sl2)

    //res = []*FakeType{
	//    &FakeType{
    //        A: 2,
    //        B: "two",
	//    },
    //    &FakeType{
    //        A: 3,
    //        B: "three",
    //    },
    //} 
    ```

* ### InFirstOnly
    Returns elements that are contained only in a first slice and is not contained in a second one.
    
    ```go
    sl1 := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }

    sl2 := []*FakeType{
	    &FakeType{
            A: 2,
            B: "two",
	    },
        &FakeType{
            A: 3,
            B: "three",
        },
        &FakeType{
            A: 4,
            B: "four",
        },
    }

    res, err := FakeTypeSlice(sl1).InFirstOnly(sl2)

    //res = []*FakeType{
	//    &FakeType{
    //        A: 1,
    //        B: "one",
	//    },
    //} 
    ```

#### If the description looks unclear for you, please take a look at [`experiment` folder](https://github.com/doctornick42/gosli/tree/master/experiment). You can find there unit test, benchmarks and some generated code that could describe the essent of the library much better than my poor English :)