 ███╗░░░███╗░█████╗░███╗░░██╗░░░░░██╗░█████╗░██████╗░░█████╗░  
 ████╗░████║██╔══██╗████╗░██║░░░░░██║██╔══██╗██╔══██╗██╔══██╗ 
 ██╔████╔██║███████║██╔██╗██║░░░░░██║███████║██████╔╝██║░░██║  
 ██║╚██╔╝██║██╔══██║██║╚████║██╗░░██║██╔══██║██╔══██╗██║░░██║  
 ██║░╚═╝░██║██║░░██║██║░╚███║╚█████╔╝██║░░██║██║░░██║╚█████╔╝  
 ╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝░╚════╝░╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░  

░██████╗███████╗████████╗██╗░░░██╗██████╗░
 ██╔════╝██╔════╝╚══██╔══╝██║░░░██║██╔══██╗
 ╚█████╗░█████╗░░░░░██║░░░██║░░░██║██████╔╝
 ░╚═══██╗██╔══╝░░░░░██║░░░██║░░░██║██╔═══╝░
 ██████╔╝███████╗░░░██║░░░╚██████╔╝██║░░░░░
 ╚═════╝░╚══════╝░░░╚═╝░░░░╚═════╝░╚═╝░░░░░
 


<img src="https://raw.githubusercontent.com/MicaelliMedeiros/micaellimedeiros/master/image/computer-illustration.png" alt="ilustração de um computador" min-width="400px" max-width="400px" width="400px" align="right">

<p align="left"> 
  Não gosto de gastar muito tempo instalando e configurando meu sistema de trabalho, mas prefiro acelerar o tempo para usá-lo de forma produtiva. Daí esses scripts que me ajudam a colocar em funcionamento um novo sistema o mais rápido possível.
</p>

<p align="left">
  <a href="#" title="Gmail">
  <img src="https://img.shields.io/badge/Linux-E34F26?style=for-the-badge&logo=linux&logoColor=black" alt="Gmail"/></a>
  <a href="#" title="LinkedIn">
  <img src="https://img.shields.io/badge/GoLang-100000?style=for-the-badge&logo=go&logoColor=white" alt="LinkedIn"/></a>
  <a href="#" title="WhatsApp">
  <img src="https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white" alt="WhatsApp"/></a>
  <a href="#" title="Facebook">
  <img src="https://img.shields.io/badge/Shell_Script-121011?style=for-the-badge&logo=gnu-bash&logoColor=white" alt="Facebook"/></a>
</p>

<p align="left">
    Dê uma olhada em <code>manjaro_install_apps.sh</code> , <code>manjaro_install_cli.sh</code> e <code>manjaro_install_web_dev.sh</code>. Antes de executar qualquer um deles, recomendo que você os abra em um editor de texto e revise o que será instalado, comente/descomente dependendo de suas necessidades, então <code>chmod +x *.sh</code> e execute.
</p>

###  🔧 Gerar o binário
Siga os passos abaixo para compilar o projeto e gerar o executável:

1. Instale o Go
Certifique-se de que você tem o Go instalado. Você pode verificar com:

``` bash
go version
```

Se não tiver, instale a partir de: https://golang.org/dl/

2. Compile o projeto
No terminal, vá até o diretório raiz do projeto e execute:

``` bash
go build -o manjaro-setup ./cmd/main.go
```

Isso irá gerar um binário chamado nome-do-binario na raiz do projeto.

3. (Opcional) Compilar para outros sistemas
Para gerar o binário para outro sistema operacional ou arquitetura:

``` bash
GOOS=linux GOARCH=amd64 go build -o manjaro-setup ./cmd/main.go
GOOS=windows GOARCH=amd64 go build -o manjaro-setup ./cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o manjaro-setup ./cmd/main.go
```

4.Execute o binário
No terminal:

```bash
./manjaro-setup
```

```shell
manjaro-setup.exe
```

# Pacman Packages

| Pacote               | Descrição                                          | Status     |
|----------------------|----------------------------------------------------|------------|
| atom                 | Editor de texto hackable e open source            | Ativo      |
| base-devel           | Pacotes básicos para compilação (make, gcc, etc.) | Ativo      |
| cmake                | Ferramenta de automação para build de projetos    | Ativo      |
| ctags                | Gera índices de tags para navegação de código     | Ativo      |
| dbeaver              | Cliente de banco de dados universal               | Ativo      |
| docker               | Plataforma para containers                        | Ativo      |
| erlang-dependencies  | Dependências para projetos Erlang/Elixir          | Ativo      |
| fish                 | Shell amigável e interativo                       | Ativo      |
| flatpak              | Sistema de empacotamento universal para Linux     | Desabilitado |
| mono                 | Plataforma .NET open source                       | Desabilitado |
| ncurses              | Biblioteca para interfaces TUI (text-based)       | Ativo      |
| neovim               | Editor de texto moderno baseado no Vim            | Desabilitado |
| notepadqq            | Editor de texto similar ao Notepad++ para Linux   | Ativo      |
| opera                | Navegador web rápido e leve                       | Ativo      |
| pamac                | Frontend gráfico para pacman                      | Ativo      |
| plank                | Dock leve para ambiente gráfico                   | Ativo      |
| postgresql           | Sistema gerenciador de banco de dados relacional  | Desabilitado |
| qbittorrent          | Cliente de torrents com interface Qt              | Ativo      |
| redis                | Banco de dados chave-valor em memória             | Desabilitado |
| snap                 | Sistema de pacotes da Canonical                   | Ativo      |
| steam                | Plataforma de jogos                               | Ativo      |
| sublime-text         | Editor de texto popular entre desenvolvedores     | Ativo      |
| telegram-desktop     | Aplicativo de desktop para o Telegram             | Ativo      |
| tig                  | Visualizador de repositórios Git no terminal      | Ativo      |
| tmux                 | Multiplexador de terminal                         | Ativo      |
| visual-studio-code   | Editor de código da Microsoft                     | Ativo      |
| wine                 | Executa aplicativos Windows no Linux              | Desabilitado |
| xclip                | Acesso ao clipboard via terminal                  | Ativo      |
| yay                  | AUR Helper para pacman                            | Ativo      |
| zsh                  | Shell poderoso e customizável                     | Desabilitado |

# Pamac Packages

| Pacote        | Descrição                                  |
|---------------|--------------------------------------------|
| slack         | Cliente oficial do Slack                  |
| virtual-box   | Virtualização de sistemas operacionais    |
| whatsapp      | Cliente de WhatsApp para desktop          |

# Yay (AUR) Packages

| Pacote               | Descrição                                  | Status     |
|----------------------|--------------------------------------------|------------|
| chrome               | Navegador web do Google                   | Ativo      |
| heroku-cli           | CLI da plataforma Heroku                  | Ativo      |
| slack-desktop        | Cliente desktop alternativo do Slack      | Ativo      |
| spotify              | Serviço de streaming de música            | Ativo      |
| teamviewer           | Ferramenta de acesso remoto               | Ativo      |
| visual-studio-code   | Editor de código da Microsoft             | Desabilitado |
| whatsapp             | Cliente de WhatsApp para desktop          | Ativo      |
| zoom                 | Plataforma de videoconferência            | Ativo      |

# Flatpak Packages (Opcional)

| Pacote             | Descrição                          |
|--------------------|------------------------------------|
| atom               | Editor de texto open source       |
| dbeaver            | Cliente de banco de dados         |
| franz              | Cliente de mensagens unificadas   |
| postman            | Testes de API                     |
| spotify            | Cliente de música                 |
| visual studio code | Editor de código                  |
| zoom               | Cliente de videoconferência       |

# Environment Setup

| Ferramenta | Descrição                                                   | Status     |
|------------|-------------------------------------------------------------|------------|
| asdf       | Gerenciador de versões para linguagens                      | Ativo      |
| docker     | Plataforma para containers                                  | Ativo      |
| elixir     | Linguagem funcional baseada em Erlang                       | Desabilitado |
| erlang     | Plataforma para aplicações distribuídas                     | Desabilitado |
| fish       | Shell moderno e interativo                                  | Ativo      |
| git        | Sistema de controle de versão                               | Ativo      |
| neovim     | Editor de texto baseado em Vim                              | Desabilitado |
| postgres   | Banco de dados relacional                                   | Desabilitado |
| ruby       | Linguagem de programação dinâmica                           | Ativo      |
| swappiness | Ajuste do uso de swap no sistema                           | Ativo      |
| teamviewer | Acesso remoto a dispositivos                                | Desabilitado |
| tmux       | Multiplexador de terminal                                   | Ativo      |
| zsh        | Shell poderoso e customizável                               | Desabilitado |