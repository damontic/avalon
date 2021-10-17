# Avalon

 ![Avalon](/assets/images/text/title.png)
Play Avalon with your friends.
![Merlin](/assets/images/character_merlin.png)

## Design
- Just one executable and a web resources directory.

## Backend
- Implemented using native ![golang](https://golang.org/).
- ![Current Kanban board](https://github.com/damontic/avalon/projects/1).
- Able to run either via HTTP or HTTPSby using a valid cert

## Frontend
- Served by the Backend logic. So it is a Monolithic app.
- Implemented using ![p5.js](https://p5js.org/).
- ![Current Kanban board](https://github.com/damontic/avalon/projects/2).

## pre-commit useful commands
For these to work you need to install the ![pre-commit](https://pre-commit.com/) tool.

### install git hooks
```
pre-commit install
```

### update git hooks
```
pre-commit autoupdate
```

### manually run pre-commit rules
```
pre-commit run --all-files
```
