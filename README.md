# GoComfyCommit

# Why

I need it this way, to write semantic commit messages for different project, which different commit message patterns.

It basically runs

`git add . and git commit -m "..."` with a customized commit messages, to be even DRY here.

Use the config.example.json (rename to config.json) to add profiles.
Add an alias, to run this profile (ie. "-p exp")

```GoComfyCommit -t "docs" -m "Update Readme" -p exp```
(on branch main that does:
```
git add . and git commit -m "My CommitType: docs at Branch: mai > My Message: Update Readme"
```
)

In "commitMessage" you can write the pattern for you messages and place the placeholders (${t}, ${b}, ${m}), where you want.

```
${b} = that is the branch name, from where you run it
${m} = That is you commit messages, passed by flag -m
${t} = Just another text variable to place somewhere a type like "fix", "docs"
```

cropBranchFromTo: Here you can truncate the branch name. If you do not have "main", but maybe a the ticketname i.e "ticket-1234-remove-js-vars-from-shell" you can set it to [0,10] to just have "ticket-1234".

## All flags


-m (set commit message)
-t (set type (just text for commit message))
-p (select profile from config)
-c (crop/trim current git branch (from, to))


## In the Fute

- [ ] Add message description (to set links)
- [ ] Set different config path as default
- [ ] Add--help command
- [x] Customizable by config file
- [x] Able to read message patterns by config file