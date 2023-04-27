# Contributing

## Release process

> Requires: `python3`

To build for release, simply run the `build.py`:

```shell
$ time python3 build.py
created out dir: './out'
I: detected 20 architecture operating system combinations, preparing build...
I: read config from build.conf:
 {'VERSION': '0.0.2-alpha', 'FEATURE': 'livepreview.1', 'FLAGS': '-w -s'}
I: prepared variables:
 {'VERSION': '0.0.2-alpha+livepreview.1', 'BUILD_AT': '2023-04-27T18:52:38.786325', 'BUILD_BY': 'xnacly-47723417+xNaCly@users.noreply.github.com'}
==============================
building for linux
building 0/20 [0.0.2-alpha_386_linux]
building 1/20 [bare:0.0.2-alpha_386_linux]
building 2/20 [0.0.2-alpha_amd64_linux]
building 3/20 [bare:0.0.2-alpha_amd64_linux]
building 4/20 [0.0.2-alpha_arm_linux]
building 5/20 [bare:0.0.2-alpha_arm_linux]
building 6/20 [0.0.2-alpha_arm64_linux]
building 7/20 [bare:0.0.2-alpha_arm64_linux]
==============================
building for windows
building 8/20 [0.0.2-alpha_386_windows]
building 9/20 [bare:0.0.2-alpha_386_windows]
building 10/20 [0.0.2-alpha_amd64_windows]
building 11/20 [bare:0.0.2-alpha_amd64_windows]
building 12/20 [0.0.2-alpha_arm_windows]
building 13/20 [bare:0.0.2-alpha_arm_windows]
building 14/20 [0.0.2-alpha_arm64_windows]
building 15/20 [bare:0.0.2-alpha_arm64_windows]
==============================
building for darwin
building 16/20 [0.0.2-alpha_amd64_darwin]
building 17/20 [bare:0.0.2-alpha_amd64_darwin]
building 18/20 [0.0.2-alpha_arm64_darwin]
building 19/20 [bare:0.0.2-alpha_arm64_darwin]
vvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
done...

________________________________________________________
Executed in   57.16 millis    fish           external
   usr time   28.55 millis  216.00 micros   28.33 millis
   sys t
```

This is fairly fast, due to multithreading in the python script.
