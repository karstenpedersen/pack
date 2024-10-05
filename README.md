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

We can now run `pack check` to see what files will get packaged and `pack` to zip them to `.pack-out/pack.zip`.

## Example Use Case

I created this tool to zip files for my university assignments. With `pack` I can compile LaTeX and then package its output together with my source code and other files.

