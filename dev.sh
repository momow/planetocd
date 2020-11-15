#!/bin/bash

~/go/bin/CompileDaemon -directory=. -command="./planetocd" -pattern='(.+\.go|.+\.c|.+\.html|.+\.yaml|.+\.css)$'
