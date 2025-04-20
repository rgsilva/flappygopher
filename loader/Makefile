PS2SDK = /usr/local/ps2dev/ps2sdk/
PS2DEV = /usr/local/ps2dev/
EE_PREFIX=mips64r5900el-ps2-elf-
#EE_PREFIX=mipsel-none-elf-

EE_BIN = main.elf
EE_OBJS = loader.o
EE_LIBS = -L$(PS2SDK)/ports/lib -L$(PS2DEV)/gsKit/lib/ -Lmodules/ds34bt/ee/ -Lmodules/ds34usb/ee/ -lpatches -lfileXio -lpad -ldebug -lmath3d -ljpeg -lfreetype -lgskit_toolkit -lgskit -ldmakit -lpng -lz -lmc -laudsrv -lelf-loader -laudsrv -lc
EE_INCS += -I$(PS2DEV)/gsKit/include -I$(PS2SDK)/ports/include -I$(PS2SDK)/ports/include/freetype2 -I$(PS2SDK)/ports/include/zlib
EE_CFLAGS = -I$(PS2DEV)/gsKit/include -mxgot
EE_LDFLAGS = -L$(PS2DEV)/gsKit/lib $(GO_OBJS) -Wl,-Map,output.map

build: freesio2.irx freepad.irx $(EE_OBJS)

freesio2.irx:
	cp $(PS2SDK)/iop/irx/freesio2.irx $@

freepad.irx:
	cp $(PS2SDK)/iop/irx/freepad.irx $@

# clean:
# 	rm -f *.ll *.o *.elf *.irx *.map *.asm

# build: red.o ps2.o go
# 	file *.o

# # elf: main.elf

# clean:
# 	rm -rf *.o *.elf *.ll

include $(PS2SDK)/samples/Makefile.pref
include $(PS2SDK)/samples/Makefile.eeglobal
