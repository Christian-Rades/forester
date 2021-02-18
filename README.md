# Forrester

Forrester is a tool to convert a sorted list of paths to a tree.
The tool receives a path per line over stdin and return a dot graph
to stdout.

# Usage

## Quick filesystem visualisation

`find . -type f | sort | forester | dot -Tsvg > files.svg`

## Tree of a symfony api

`for controller in (find . -type f -name '*Controller.php'); sed -n 's/^.*@Route("\([^"]*\).*$/\1/p' $controller; end | sort | forester | dot -Tsvg > api.svg`
(fish shell)
