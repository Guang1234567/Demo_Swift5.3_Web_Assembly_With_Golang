
#ifndef SWIFTWASM_EXPORTED_H
#define SWIFTWASM_EXPORTED_H

#include <stdlib.h>
#include <stdint.h>
#if __wasm32__

extern void helloFromHostImpl();

#endif

#endif //SWIFTWASM_EXPORTED_H
