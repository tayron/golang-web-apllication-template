#!/bin/bash

diretorioAplicacao="$(pwd)"
nomeAplicacao="hospeda.app"
nomeAplicacaoTesteUnitario="hospeda.app_test"

echo && \
echo "---------------------------------------------------------" && \
echo "Criando ambiente de teste:" && \
echo "---------------------------------------------------------" && \
cp -R $diretorioAplicacao "$diretorioAplicacao.test"/ && \
cd "$diretorioAplicacao.test"/ && \

echo && \
echo "---------------------------------------------------------" && \
echo "Baixando atualizações:" && \
echo "---------------------------------------------------------" && \
git pull && \

echo && \
echo "---------------------------------------------------------" && \
echo "Configurando aplicação:" && \
echo "---------------------------------------------------------" && \
export ENVIROMENT="production" && \

echo && \
echo "---------------------------------------------------------" && \
echo "Rodando testes unitário:" && \
echo "---------------------------------------------------------" && \
./$nomeAplicacaoTesteUnitario && \

echo && \
echo "---------------------------------------------------------" && \
echo "Baixando atualizações no projeto principal:" && \
echo "---------------------------------------------------------" && \
cd $diretorioAplicacao && \
rm -rf "$diretorioAplicacao.test" && \
git pull && \

echo && \
echo "---------------------------------------------------------" && \
echo "Subindo aplicação:" && \
echo "---------------------------------------------------------" && \
docker-compose up --build -d && \

echo && \
echo && \
echo "---------------------------------------------------------" && \
echo "Restartando a aplicação:" && \
echo "---------------------------------------------------------" && \
docker restart $nomeAplicacao && \

echo && \
echo && \
echo "---------------------------------------------------------" && \
echo "Obtendo ip da aplicação:" && \
echo "---------------------------------------------------------" && \
echo && \
docker inspect -f '{{range .NetworkSettings.Networks}}| {{.IPAddress}} {{end}}' $nomeAplicacao && \
echo 
echo