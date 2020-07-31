# arange
arange

How do I do something like numpy's arange in Go?

Numpy's arange function returns a list of evenly spaced values within a given interval with float steps. Is there a short and simple way to create a slice like that in Go?

Here is the implementation of the answer which comes from https://stackoverflow.com/a/19908214

Put it here for any one's convenience


# Usage

```
arange(0., 10., 0.5)
will output

[0 0.5 1 1.5 2 2.5 3 3.5 4 4.5 5 5.5 6 6.5 7 7.5 8 8.5 9 9.5]


Corresponding to the Python:

np.arange(0., 10., 0.5)

which output 
array([ 0. ,  0.5,  1. ,  1.5,  2. ,  2.5,  3. ,  3.5,  4. ,  4.5,  5. ,
    5.5,  6. ,  6.5,  7. ,  7.5,  8. ,  8.5,  9. ,  9.5])
```

# MIT license
