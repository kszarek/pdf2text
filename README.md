# pdf2text

Usage

```
curl -i --data-binary @file.pdf http://localhost:5000/
```

Benchmark

```
ab -n 500 -c 20 -T "application/pdf" -p file.pdf http://localhost:5000/
```
