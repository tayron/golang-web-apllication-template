#!/bin/bash

nomeAplicacao="hospeda.app"
nomeAplicacaoTesteUnitario="hospeda.app_test"

clear
echo "---------------------------------------------------------" && \
echo "Gerando executável dos testes unitário:" && \
echo "---------------------------------------------------------" && \
go test tests/* -c -o $nomeAplicacaoTesteUnitario && \
upx $nomeAplicacaoTesteUnitario && \

echo && \
echo "---------------------------------------------------------" && \
echo "Configurando aplicação:" && \
echo "---------------------------------------------------------" && \
export ENVIROMENT="developer" && \

echo && \
echo "---------------------------------------------------------" && \
echo "Rodando teste unitário:" && \
echo "---------------------------------------------------------" && \
./$nomeAplicacaoTesteUnitario && \

echo && \
echo "---------------------------------------------------------" && \
echo "Gerando executável da aplicação:" && \
echo "---------------------------------------------------------" && \
go build -o $nomeAplicacao -ldflags "-s -w" main.go && \
upx $nomeAplicacao