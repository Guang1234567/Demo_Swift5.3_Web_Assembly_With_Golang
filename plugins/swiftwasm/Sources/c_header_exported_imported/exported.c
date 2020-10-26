
#include "exported.h"
#include "stdlib.h"

__attribute__((__export_name__("hello"))) void helloFromHost()
{
    helloFromHostImpl();
}
