# tsk

A task runner that only requires a `scripts` folder and a `Tskfile.yml` with one line.

## How it works

Put your scripts inside a `scripts` folder at the root of the project, and add
a `Tskfile.yml` next to it, like this:

```yaml
name: MyProject
```

Now, from anywhere inside your project, from root to any folder as deep as you
want, run `tsk`.

```
tsk MyProject
Usage: tsk [command] <subcommands...>

  build scripts/build.sh
  gendoc scripts/gendoc.sh
```

Running the scripts with `tsk` always runs them with the project's root folder
as working directory, doesn't care about where you run `tsk`.

## Features

- [X] Run bash files inside `scripts` folder.
- [ ] Make use of the script file's shebang to decide how to run it.
    - Right now, it must be a `.sh` file.
- [ ] Bash autocomplete support.

## Install

Having the go tools installed, build from source and install:

```bash
go get -u github.com/Sirikon/tsk/cmd/tsk
```
