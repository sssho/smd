# smd
Get markdown table from spreadsheet style text on clipboard

## Usage
1. Copy spreadsheet style text (from Google spreadsheet, MS excel, etc)
1. Run `smd` (Markdown table text is written to clipboard)
1. Paste

## What is spreadsheet sytle text?
```
head0<TAB>head1<TAB>head2\n
data0<TAB>data1<TAB>data2\n
data3<TAB>data4<TAB>data5
```

## Requirment
`xclip` or `xsel` for Linux, Unix

## Run smd on Windows
If you are in Windows, it's convenient to use Windows `Run` to run `smd`.

Press <kbd>Win</kbd>–<kbd>r</kbd>, then type `smd` 

Please make sure that `smd.exe` is in a folder which is in `%PATH%`
