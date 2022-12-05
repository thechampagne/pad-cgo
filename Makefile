CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/pad.a build/pad.so

.PHONY: all
all: $(LIBS)

build/pad.a: pad.go
	$(CGO) $(STATIC) -o build/pad.a $<

build/pad.so: pad.go
	$(CGO) $(SHARED) -o build/pad.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
