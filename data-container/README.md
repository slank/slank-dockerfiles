This is a minimal image for implementing the container-as-volume pattern. The
peristent volume is exposed at /persistent.

Use it thusly:

```
docker run -name myapp_data slank/data-container
docker run -name myapp -volumes-from myapp_data myapp-image
```

The included `noop.go` source file can be built by running `make`. It simply
exits 0 (success). Using this binary results in an image roughly 1/4 the size
of a busybox-based image.
