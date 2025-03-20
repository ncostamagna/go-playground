// mycpp.cpp
#include <iostream>

extern "C" {
    void printInt(int x) {
        std::cout << "Int: " << x << std::endl;
    }

    void printStr(const char* s) {
        std::cout << "Str: " << s << std::endl;
    }
}
