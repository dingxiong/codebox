# CXX := /opt/homebrew/opt/llvm/bin/clang++
CXX := clang++


.PHONY: all clean

all: Hello.out

Hello.pcm: Hello.cppm
	$(CXX) -std=c++20 $^ --precompile -o $@

Hello.out: main.cc Hello.pcm 
	$(CXX) -fmodule-file=Hello=Hello.pcm $^ -o $@

clean:
	rm -f Hello.pcm Hello.out
