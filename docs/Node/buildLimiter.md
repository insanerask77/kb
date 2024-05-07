You need to control the max memory size flags (all sizes are taken in MB).

The [recommended amounts for a "low memory device" are](https://github.com/v8/v8/blob/2b17c752d117bf0ededd23206bb38b69e641a976/src/heap/heap.h#L1106-L1136):

```js
node --max-executable-size=96 --max-old-space-size=128 --max-semi-space-size=1 app.js
```

for 32-bit and/or Android and

```js
node --max-executable-size=192 --max-old-space-size=256 --max-semi-space-size=2 app.js
```

for 64-bit non-android.

These would limit the heap totals to 225mb and 450mb respectively. It doesn't include memory usage outside JS. For instance buffers are allocated as "c memory" , not in the JavaScript heap.

Also you should know that the closer you are to the heap limit the more time is wasted in GC. E.g. if you are at 95% memory usage 90% of the CPU would be used for GC and 10% for running actual code (not real numbers but give the general idea). So you should be as generous as possible with the limits and never exceed say 16% of the maximum memory usage (I.E. `heapUsed/limit` should not be greater than `0.16`). 16% is just something I recall from some paper, it might not be the most optimal.

Flags:

- `--max-executable-size` the maximum size of heap reserved for executable code (the native code result of just-in-time compiled JavaScript).
- `--max-old-space-size` the maximum size of heap reserved for long term objects
- `--max-semi-space-size` the maximum size of heap reserved for short term objects