# overflow me
Overflow me is a way to overflow a string into a specific data type to replace its value

```bash
./overflowme --char-arr-len 8 --char-arr-data "debin" --char-arr-filler "a" --overflow-target-data "1337" overflow-target-type "uint"
# replacement length is: 8

# payload: debinaaa\x39\x05\x00\x00
```
