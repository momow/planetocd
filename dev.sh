#!/bin/bash

~/go/bin/CompileDaemon -directory=. -command="./planetocd -action=server" -pattern='(.+\.go|.+\.c|.+\.html|.+\.yaml|.+\.css|.+\.json|\.md)$'
