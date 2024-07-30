# templ-bulma

## Status: Unstable

Feel free to use this as is, but somethings still might move around as I continue to improve the project.
Please wait for the first release.

## Why

I wanted add building blocks for templ so I had to look at less html and more Go code.
Overall I found it to be helpful because everything now populates with intellasence. 
With the naming pattern it should be easy to find the related components and build on top of it.

## How to use 

This package needs to be added to your project as a `git submodule`.

```bash
git submodule init
git submodule add https://github.com/jtom38/templ-bulma
templ generate
```

When a new version comes out you can run the following command.
This will download the newest version available in this repo.

```bash
git submodule update --remote
templ generate
```

### Why a submodule?

I attempted to use `go get` and the files never got found.
I attempted to then generate the templates and upload them, did not work.
Go was not able to find the files I was hoping for.

Uploading compiled templates would also require that you keep your templ version locked to what this project uses.
That does not work for me at all.
As new versions come out you should be able to use the newest version reguardless of your templates.

When you run `templ generate` its only looking at templ files with in your project/directory.
So by moving to a submodule, the templates are now loaded into your project and you can geneate the temp.

## HTMX

This has some HXMX features and I expect to add more.  
Right now the form allows you to use a `hx-post` from the struct params that can be passed in.

## Closing

If you find this helpful, leave a ⭐️!
If you find any issues or improvement, please create a PR.
