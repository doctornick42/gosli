### **Basic types**

We need to get gosli. All we need to do is run:
```
go get github.com/doctornick42/gosli
```
For all basic types we got a structure representing a slice of it. For example, slice of `int64` is wrapped
into `Int64Slice` structure. All methods are the same for all primitive types, so on this page
there's an example of interacting with `Int64Slice` but the idea will be exactly the same for all other types.

---
## **Methods**


* ### First
    Returns first item of a slice that is passed through a filter.

    If an item wasn't found, the method returns an error with "Not found" message.
    
    ```go
    sl := []int64{1, 2, 3}
    
    filter := func(f int64) bool {
        return f == 2
    }

    res, err := Int64Slice(sl).First(filter)

    //Result:
    //res = 2 
    ```

* ### FirstOrDefault
    Returns first item of a slice that is passed through a filter.

    If an item wasn't found, the result is the default value of undelying type
    (0 for `int64`).
    
    ```go
    sl := []int64{1, 2, 3}
    
    filter := func(f int64) bool {
        return f == 2
    }

    res := Int64Slice(sl).FirstOrDefault(filter)

    //Result:
    //res = 2 
    ```

* ### Where
    Returns all items of a slice that is passed through a filter.

    If items weren't found, the result is empty slice.
    
    ```go
    sl := []int64{1, 2, 3}
    
    filter := func(f int64) bool {
        return f > 1
    }
    res := Int64Slice(sl).Where(filter)

    //Result:
    //res = []int64{2, 3} 
    ```

* ### Select
    Applies a function to every item of a slice and returns slice of results.
    
    ```go
    sl := []int64{1, 2, 3}
    
    type tempType struct {
		Msg string
	}

    filter := func(f int64) interface{} {
        return &tempType{
            Msg: fmt.Sprintf("Value: %v", f),
        }
    }
    res, err := Int64Slice(sl).Select(f)

    //Result:
    //res = []interface{}{
    //    &tempType{
    //        Msg: "Value: 1",
    //    },
    //    &tempType{
    //        Msg: "Value: 2",
    //    },
    //    &tempType{
    //        Msg: "Value: 3",
    //    },
    //}
    ```

* ### Page
    Returns paginated slice according to given `number` (number of selected page) and `perPage` 
    (items per a page). `number` parameter should start with 1 (not with 0).
    
    ```go
    sl := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    pageNumber := 1
    perPage := 2
    res, err := Int64Slice(sl).Page(pageNumber, perPage)

    //Result:
    //res = []int64{1, 2, 3, 4, 5} 
    ```

* ### Any
    Returns `true` if any item of the slice is passed through a filter.

    ```go
    sl := []int64{1, 2, 3}
    
    filter := func(f int64) bool {
        return f == 2
    }
    res, err := Int64Slice(sl).Any(filter)

    //Result:
    //res = true
    ```

* ### Contains
    Returns `true` if a slice contains at least one item that is equal to the desired one.
    
    ```go
    sl := []int64{1, 2, 3}

    el := 2
    
    res, err := Int64Slice(sl).Contains(el)

    //Result:
    //res = true
    ```

* ### GetUnion
    Returns a slice that contains items that are contained in both given slices.
    
    ```go
    sl1 := []int64{1, 2, 3}
    sl2 := []int64{2, 3, 4}

    res, err := Int64Slice(sl1).GetUnion(sl2)

    //Result:
    //res = []int64{2, 3}
    ```

* ### InFirstOnly
    Returns elements that are contained only in a first slice and is not contained in a second one.
    
    ```go
    sl1 := []int64{1, 2, 3}
    sl2 := []int64{2, 3, 4}

    res, err := Int64Slice(sl1).InFirstOnly(sl2)

    //Result:
    //res = []int64{1} 
    ```

#### If the description looks unclear for you, please take a look at [unit tests for Int64Slice type](https://github.com/doctornick42/gosli/tree/master/int64_test.go).