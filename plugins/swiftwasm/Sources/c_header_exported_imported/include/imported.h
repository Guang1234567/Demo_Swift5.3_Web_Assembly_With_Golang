
#ifndef SWIFTWASM_IMPORTED_H
#define SWIFTWASM_IMPORTED_H

#include <stdlib.h>
#include <stdint.h>
#if __wasm32__

__attribute__((__import_name__("fetch_code"))) extern int c_fetchCode(int input);

#endif

#endif //SWIFTWASM_IMPORTED_H
