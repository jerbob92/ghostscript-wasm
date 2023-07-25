# Make sure the ghostscript submodule is loaded
git submodule update --init --recursive

# Patch the ghostscript submodule
cd lib/ghostscript
git apply ../../ghostscript.patch
cd ../..

# Get the emsdk repo
git clone https://github.com/emscripten-core/emsdk.git

# Enter that directory
cd emsdk

# Fetch the latest version of the emsdk (not needed the first time you clone)
git pull

# Checkout the correct version
git checkout 3.1.44

# Download and install the SDK tools.
./emsdk install 3.1.44

# Make the SDK version active for the current user. (writes .emscripten file)
./emsdk activate 3.1.44

# Activate PATH and other environment variables in the current terminal
source ./emsdk_env.sh

# Build the Emscripten freetype version and copy it into ghostscript
embuilder build freetype
rm -Rf ../lib/ghostscript/freetype
mkdir ../lib/ghostscript/freetype
cp -R upstream/emscripten/cache/ports/freetype/FreeType-version_1/* ../lib/ghostscript/freetype

# Go to the emscripten directory
cd upstream/emscripten

# Remove the cache dir
rm -Rf cache

# Apply our emscripten WASI patch
patch -p1 < ../../../emscripten.patch

# Go back to the root of the build dir
cd ../../../

# Create the prefix directory (I couldn't make this work by just setting the prefix in the compiled binary)
sudo mkdir -p /ghostscript
sudo chmod 777 /ghostscript

# Run the build script
./build.sh

# Copy the generated binaries into the wasm folder of the repository.
cp lib/ghostscript/bin/gs.wasm ../wasm
rm -Rf ../ghostscript
cp -R /ghostscript ../ghostscript
sudo rm -Rf /ghostscript

