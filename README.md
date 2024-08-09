# VS Code Profile Toolkit

A simple tool to manage `.code-profile` files.

## What are .code-profile files?

VS Code has a [profile](https://code.visualstudio.com/docs/editor/profiles) system, to manage multiple configurations.

When you export an existing profile to disk, you get a `.code-profile` file (a nested JSON file).

## Features

For now, the tool can do two things:

- `extract`, to unpack a `.code-profile` file into a specific folder,
- `archive`, to pack a profile folder into a `.code-profile` file.

## Why?

I have multiple profiles, and I want to easily compare them and write my own profiles by hand.
One day I might add a "diff" command to the tool to have a better view of the differences of each profile, but for now here's my workflow:

- In VS Code, export your profiles somewhere on your filesystem (as `.code-profile` files)
- Use `vs-prof-tk extract` to unpack the profiles in a folder
    - Example: `vs-prof-tk extract -i /work/Default.code-profile /work/extracted`
- Edit the files as you want
- Recreate a `.code-profile` file using `vs-prof-tk archive`
    - Example: `vs-prof-tk archive -i /work/extracted/Default /work/DefaultNew.code-profile`
- Now you can import your new profile in VS Code using the "Import profile" feature

## Building

You need a recent version of Go.
If you use `nix`, you can use the included `shell.nix` file.

```
go build -o vs-prof-tk
```