# pdf2text

Usage

```
curl -i --data-binary @file.pdf http://127.0.0.1:5000/
```

Benchmark

```
ab -n 500 -c 20 -T "application/pdf" -p file.pdf http://127.0.0.1:5000/
```
