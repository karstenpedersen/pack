# pack

CLI tool to make it easier to archive files.

## Get started

Start by initializing pack:

```bash
pack init
```

This will create a `pack.config.json` file:

```json
{
    "name": "pack",
    "method": "zip",
    "outDir": ".pack-out",
    "include": [],
    "hooks": {}
}
```

You can then specify what files you want to include. Here I specify that I want to include all markdown files:

```json
{
    "name": "pack",
    "method": "zip",
    "outDir": ".pack-out",
    "include": [
        "*.md"
    ],
    "hooks": {}
}
```

We can now run `pack check` to see what files will get packaged and `pack run` to zip them to `.pack-out/pack.zip`.

## Example Use Case

I created this tool to zip files for my school assignments. With `pack run` I can compile LaTeX and then package its output together with my source code and other files.

## Commands

```bash
pack # help
pack --help # help
pack --version # prints version

pack init # initialize pack config
pack check # check what files will be packaged
pack run # package files
pack show # show packaged files
```
