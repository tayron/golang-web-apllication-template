# Go Web Apllication Server

Aplicação web básica feito com **go** contendo:
* Testes unitários
* Script para execução da aplicação localmente (Ambiente de desenvolvimento)
* Script para geração do binário da aplicação para uso em produção
* Script de deploy usado no ambiente de produção

## Para trabalhar com TimeZone personalizado
```sh
var cstSh, _ = time.LoadLocation("Asia/Shanghai") //ShangHai
fmt.Println("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
```

## Executar aplicação localmente sem usar container docker
```sh 
./run.sh
```

Retorno da execução acima deverá ser algo semelhante:

```sh
---------------------------------------------------------
Configurando aplicação:
---------------------------------------------------------
2021/08/03 18:12:50 listening on http://localhost:3000
```

## Gerar binário da aplicação
```sh 
./build.sh
```

Retorno da execução acima deverá ser algo semelhante:

```sh
---------------------------------------------------------
Gerando executável dos testes unitário:
---------------------------------------------------------
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2018
UPX 3.95        Markus Oberhumer, Laszlo Molnar & John Reiser   Aug 26th 2018

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   3211392 ->   1810964   56.39%   linux/amd64   hospeda.app_test              

Packed 1 file.

---------------------------------------------------------
Configurando aplicação:
---------------------------------------------------------

---------------------------------------------------------
Rodando teste unitário:
---------------------------------------------------------
PASS

---------------------------------------------------------
Gerando executável da aplicação:
---------------------------------------------------------
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2018
UPX 3.95        Markus Oberhumer, Laszlo Molnar & John Reiser   Aug 26th 2018

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   6647808 ->   2538320   38.18%   linux/amd64   hospeda.app                   

Packed 1 file.
tayron@tayron:~
```
## Executar aplicação utilizando container docker ou deploy no servidor de produção

Abra o arquivo docker-compose.yml e altere ```URL=http://localhost``` para o endereço do seu site, exemplo: ```URL=https://www.hospeda.app```.
Caso deseje testar o container localmente, basta deixar como ```URL=http://localhost```.

Logo em seguida execute:
```sh 
./deploy.sh
```

Retorno da execução acima deverá ser algo semelhante:

```sh
---------------------------------------------------------
Criando ambiente de teste:
---------------------------------------------------------

---------------------------------------------------------
Baixando atualizações:
---------------------------------------------------------
Already up to date.

---------------------------------------------------------
Configurando aplicação:
---------------------------------------------------------

---------------------------------------------------------
Rodando testes unitário:
---------------------------------------------------------
PASS

---------------------------------------------------------
Baixando atualizações no projeto principal:
---------------------------------------------------------
Already up to date.

---------------------------------------------------------
Subindo aplicação:
---------------------------------------------------------
/usr/lib/python3/dist-packages/requests/__init__.py:91: RequestsDependencyWarning: urllib3 (1.26.5) or chardet (3.0.4) doesn't match a supported version!
  RequestsDependencyWarning)
Creating network "go-web-apllication-server_default" with the default driver
Building hospeda.app
Step 1/7 : FROM golang
 ---> bfb42f2526a7
Step 2/7 : MAINTAINER Hospeda App <ti@hospeda.app>
 ---> Using cache
 ---> 1497cdb89403
Step 3/7 : RUN apt-get update
 ---> Using cache
 ---> 0fcc5ac0ba86
Step 4/7 : RUN apt-get install ffmpeg -y
 ---> Using cache
 ---> 092621c625cc
Step 5/7 : WORKDIR /home
 ---> Using cache
 ---> 91920fde7eef
Step 6/7 : ENTRYPOINT ./web_application
 ---> Using cache
 ---> 8f3024dfc86b
Step 7/7 : EXPOSE 80
 ---> Using cache
 ---> 2c8709d67c34
Successfully built 2c8709d67c34
Successfully tagged go-web-apllication-server_hospeda.app:latest
Creating hospeda.app ... done


---------------------------------------------------------
Restartando a aplicação:
---------------------------------------------------------
hospeda.app


---------------------------------------------------------
Obtendo ip da aplicação:
---------------------------------------------------------

| 172.28.0.2 


tayron@tayron:~/go/src/github.com/tayron/go-web-apllication-server$ ./deploy.sh 

---------------------------------------------------------
Criando ambiente de teste:
---------------------------------------------------------

---------------------------------------------------------
Baixando atualizações:
---------------------------------------------------------
Already up to date.

---------------------------------------------------------
Configurando aplicação:
---------------------------------------------------------

---------------------------------------------------------
Rodando testes unitário:
---------------------------------------------------------
PASS

---------------------------------------------------------
Baixando atualizações no projeto principal:
---------------------------------------------------------
Already up to date.

---------------------------------------------------------
Subindo aplicação:
---------------------------------------------------------
/usr/lib/python3/dist-packages/requests/__init__.py:91: RequestsDependencyWarning: urllib3 (1.26.5) or chardet (3.0.4) doesn't match a supported version!
  RequestsDependencyWarning)
Building hospeda.app
Step 1/5 : FROM golang
 ---> bfb42f2526a7
Step 2/5 : MAINTAINER Hospeda App <ti@hospeda.app>
 ---> Using cache
 ---> 1497cdb89403
Step 3/5 : WORKDIR /home
 ---> Running in df9906c9a674
Removing intermediate container df9906c9a674
 ---> 1d9ac2c092a6
Step 4/5 : ENTRYPOINT ./web_application
 ---> Running in 8ba96f21116b
Removing intermediate container 8ba96f21116b
 ---> f45de30c7eb0
Step 5/5 : EXPOSE 80
 ---> Running in d294bd129c96
Removing intermediate container d294bd129c96
 ---> 5b99cb450ab7
Successfully built 5b99cb450ab7
Successfully tagged go-web-apllication-server_hospeda.app:latest
Recreating hospeda.app ... done


---------------------------------------------------------
Restartando a aplicação:
---------------------------------------------------------
hospeda.app


---------------------------------------------------------
Obtendo ip da aplicação:
---------------------------------------------------------

| 172.28.0.2
```

A aplicação estará rodando e acessível através do endereço **http://172.28.0.2**
