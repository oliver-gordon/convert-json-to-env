# convert-json-to-env

Wrote this to quickly convert files between JSON and .env for when I use hashicorp vault.

Pass a JSON file, it'll convert to an suitable .env.

{
    "name":"Bob"
}

will become

NAME=Bob

It'll convert anything just not correctly
i.e [1,2,3] becomes [1 2 3] in an .env file.

Pass `-print` before your file argument to just print the values to the console without saving to a file.

.env file will be stored as the currently time `hh-mm-ss`.