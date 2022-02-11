# cranlogs

Access [cranlogs](https://cranlogs.r-pkg.org/) from the 
command line.

## Install

```bash
go get github.com/devOpifex/cranlogs
```

or

```bash
go install github.com/devOpifex/cranlogs@latest
```

## Help

```bash
cranlogs -h
cranlogs daily -h
```

## Completion

```bash
cranlogs completion -h
```

## Examples

```bash
cranlogs top
cranlogs top -n=3
cranlogs trending -r=true
cranlogs daily --period="last-month" --package="echarts4r"
```
