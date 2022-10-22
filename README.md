# Ideas
A Lightweight cli tool for keeping ideas. This is a clone of idea cli tool which was originally written in node.js https://github.com/IonicaBizau/idea.

## Installation

```sh
go install github.com/rmsubekti/ideas@latest
```

then run ideas --help to see what this cli tool can do.

```sh
$ ideas --help
NAME:
   ideas - A lighweight CLI tool for keeping your ideas.

USAGE:
   ideas [command] <ideas|state|id>

COMMANDS:
   init, i    Create new .ideas.json file in the current directory. Default: ~/.ideas.json
   create, c  Create new idea. Example: `ideas create CLI app`
   delete, d  Delete an idea. Example: `ideas delete 1`
   solve, s   Solve an idea. Example: `ideas solve 1`
   list, l    List all ideas. Example `ideas list solved`.(default: open)
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

All ideas is stored in your home directory. Run command `ideas init` store your ideas in the current directory, for example:

```sh
cd /to/your/project/directory
ideas init
```