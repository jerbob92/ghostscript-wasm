#!/bin/bash
set -euo pipefail

OUT_DIR="$PWD/out"
ROOT="$PWD"
EMCC_FLAGS_DEBUG="-g"
EMCC_FLAGS_RELEASE="-O2"

export CPPFLAGS="-I$OUT_DIR/include"
export LDFLAGS="-L$OUT_DIR/lib"
export PKG_CONFIG_PATH="$OUT_DIR/lib/pkgconfig"
export EM_PKG_CONFIG_PATH="$PKG_CONFIG_PATH"
export CFLAGS="$EMCC_FLAGS_DEBUG"
export CXXFLAGS="$CFLAGS"
export TARGET_ARCH_FILE="$ROOT/arch_wasm.h"

mkdir -p "$OUT_DIR"

cd "$ROOT/lib/ghostscript"

# There is a bug in this version of Ghostscript that prevents passing in gcc to compile the build tools, replace the var manually.
sed -i "s/CCAUX=@CC@/CCAUX=gcc/g" base/Makefile.in

emconfigure ./autogen.sh \
  CFLAGSAUX= CPPFLAGSAUX= \
  --host="wasm32-unknown-linux" \
  --prefix="$OUT_DIR" \
  --disable-cups \
  --disable-dbus \
  --disable-gtk \
  --with-system-libtiff

export GS_LDFLAGS="\
-s ALLOW_MEMORY_GROWTH=1 \
-s WASM=1 \
-s ALLOW_MEMORY_GROWTH=1 \
-s STANDALONE_WASM=1 \
-sERROR_ON_UNDEFINED_SYMBOLS=0 \
-s USE_ZLIB=1"

emmake make \
  LDFLAGS="$LDFLAGS $GS_LDFLAGS" \
  -j install

exit 1
