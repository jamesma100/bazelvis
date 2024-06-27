# bazelvis
Bazel projects can get large and their dependencies can become seemingly intractable.
`bazelvis` allows you to visualise the dependency graph like a filesystem which you can navigate recursively.

I got the original idea from one of my favorite Vim features, which is being able to open directories and traverse them, returning to the parent if needed.
Turns out this UI pattern is quite suitable for any type of acyclic graph.

## Example
Here is a simple C++ project from the Bazel [tutorials page](https://bazel.build/start/cpp), which has the following dependency graph.

![cpp-tutorial-stage3](https://github.com/jamesma100/bazelvis/assets/44740178/6ea2a94d-deb9-4709-ad3f-9bffa516cece)

We can run `bazelvis` on the main target via:
```
bazelvis //main:hello-world
```
And there you have it!

![2024-06-19 20 15 49](https://github.com/jamesma100/bazelvis/assets/44740178/e7f913b5-8532-47d1-9163-429b7b2ecc7c)

## Setup
Just run `./build.sh` and move the generated Go binary under `./bin/bazelvis` to your `PATH`.
Then you can run `bazelvis //some:target` in any Bazel workspace.

### Keybindings
- `k` (or up-arrow): move cursor up
- `j`: (or down-arrow): move cursor down
- [TODO] `ctrl-f`: move down a page
- [TODO] `ctrl-b`: move up a page
