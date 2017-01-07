# timeu
time utility.

## Usage

```
timeu [-s|-v] [time]
timeu [-s|-v] [time1] [+ | -] [time2]
```

## Examples
```
$ timeu 56
00:00:56

$ timeu -s 42:56
2576

$ timeu 11:42:56 + 2:30:59
14:13:55

$ timeu -s 11:42:56 + 2:30:59
51235

$ timeu 11:42:56 - 2:30:59
09:11:57
```
