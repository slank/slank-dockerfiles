This is a minimal image for implementing the container-as-volume pattern. The
peristent volume is exposed at /persistent.

Use it thusly:

```
docker run -name myapp_data slank/data-container
docker run -name myapp -volumes-from myapp_data myapp-image
```

The included `setup.go` source file can be built by running `make`. Using this
binary results in an image roughly 1/2 the size of a busybox-based image.

The image's entrypoint accepts two arguments: the numeric uid and gid for
/persistent. After setting ownership on the directory, it exits 0 (success).
