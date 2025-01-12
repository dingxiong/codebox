
" My principles of vim config
" 1. avoid too many plugins and complicated plugins like YouCompleteMe. I am not going to build vim like an IDE
" 2. Only use plugins that support vim 8 package installation 

" ======= basic setup =======
set number
set wildmenu
set wildmode=list:longest,full
syntax on

" ======= copy & paste =======
" use clipboard register "+
" nnoremap Y "+y


" ======= auto complete =======
" do not install any such plugin. Just use Ctrl + N/P 

" ======= spell check =======
set spelllang=en_us
autocmd FileType gitcommit setlocal spell
autocmd FileType tex setlocal spell

" ======= status bar =======
" install https://github.com/itchyny/lightline.vim
set laststatus=2


" ======= tab <> space =======
set expandtab
set tabstop=2
set shiftwidth=2

autocmd Filetype ruby setlocal ts=2 sts=2 sw=2
autocmd Filetype tex setlocal ts=2 sts=2 sw=2
autocmd Filetype sh setlocal ts=2 sts=2 sw=2
autocmd Filetype cpp setlocal ts=2 sts=2 sw=2
