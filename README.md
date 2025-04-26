# Welcome

## What's this ?

This is a tool for calculating the stars count of github repos by specific account.

The result of calculation will be displayed by symbol ( defalut: * ).

## How to use ?

`go run main.go -account "rails" -breakline 30`:

```
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
******************************
```

## Arguments

### Necessary

- account: github account, such as wl02599509.

### Optional

- column: what data you want to calculate. ( 'stargazers_count' is the only available column. )

- symbol: what symbol you want to display. ( defalut: * )

- breakline: how much number of symbol you want to break the line. ( default: 0 )
