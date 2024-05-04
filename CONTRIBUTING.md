# THIS DOCUMENT IN DEVELOPMENT

## Naming

### Imports
You should use these aliases:

| Alias      | Import path                                        |
|------------|----------------------------------------------------|
| tf         | github.com/vitaliy-ukiru/telebot-filter/telefilter |
| tb or tele | gopkg.in/telebot.v3                                |

## Commit message

### Changed package
If you commit targets to package
in commit message you must indicate that package was changed.

If you change works with many packages split the commit into several 
or place all packages.

But change like code format, dependencies, etc can be without with package.

### Message format
Preferred format is [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
or similar formats.

If you commit includes breaking changes place `!` in commit message
and add "BREAKING CHANGES" in commit body with description and 
motivation for this.

Example:
```
feat(telefilter)!: change Filter return's type to error

BREAKING CHANGES: It helps to control flow of handling update 
```

Example without scoped package:
```
style: go fmt whole project
```

## Tests and documentation

Tests are good, but creating them for absolutely every action 
is rather pointless. Don't write a lot of tests just for more coverage. 
It's better to make them fewer, but in more relevant places.

Documentation is preferred thing. It will also be good if 
you add examples for new functionality, which will be displayed 
in the go doc.

