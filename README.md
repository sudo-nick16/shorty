### Shorty

- Simple URL shortener.
- This web app was built to help me revise my skills in Go and Bazel. It is a simple, bare-bones app, but it serves its purpose well.

To run the api using bazel: 

```
FOO=BAR BIZ=BAZ bazel run //api:shorty
```

To build a docker image using bazel:

```
bazel build //api:shorty_image
```
