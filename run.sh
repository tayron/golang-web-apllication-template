#!/bin/bash
echo && \
echo "---------------------------------------------------------" && \
echo "Configurando aplicação:" && \
echo "---------------------------------------------------------" && \
export ENVIROMENT="developer" && \
export URL="http://localhost" && \
export PORT="3000" && \
go run *.go

