# GoCommit

# Why

I need it this way.

It basically runs

`git add . and git commit -m "<type = first cli arg><Curernt Git Branch><message = second cli arg>"`

With two arguments

```GoCommit "fix" "Remove JavaScript from Hero CTA"```

Git log would show: *fix: [branchName] Remove JavaScript from Hero CTA*

OR with one argument

```GoCommit "Remove JavaScript from Hero CTA"```

`git add . and git commit -m "<message = second cli arg>"`

Git log would show: *Remove JavaScript from Hero CTA*

OR with none argument

```GoCommit```

`git add . and git commit -m "Update"`

Git log would show: *Update*


## How to run
./GoCommit "fix" "Remove JavaScript from Hero CTA"
## Good to know
Curernt Git Branch is currently trimmed by 8 (which will be later defined by a config file).

## Usage

```shell
go build
./GoCommit

# or

go run main.go
```

## In the Fute

- [ ] Customizable by config file
- [ ] Able to read message patterns by config file