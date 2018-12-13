#include "bindings.h"
#include <stdio.h>
#include <inttypes.h>
#include <Zydis/Zydis.h>


void printme(unsigned char *data, unsigned long length, long unsigned int runtime_address) {

    // Initialize decoder context
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);

    // Initialize formatter. Only required when you actually plan to do instruction
    // formatting ("disassembling"), like we do here
    ZydisFormatter formatter;
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);

    // Loop over the instructions in our buffer.
    // The runtime-address (instruction pointer) is chosen arbitrary here in order to better
    // visualize relative addressing
    ZyanUSize offset = 0;
    ZydisDecodedInstruction instruction;
    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(&decoder, data + offset, length - offset,
        &instruction))) {
        // Print current instruction pointer.
        printf("%016" PRIX64 "  ", runtime_address);

        // Format & print the binary instruction structure to human readable format
        char buffer[256];
        ZydisFormatterFormatInstruction(&formatter, &instruction, buffer, sizeof(buffer),
            runtime_address);
        puts(buffer);

        offset += instruction.length;
        runtime_address += instruction.length;
    }
}
