Get the emsdk repo

`git clone https://github.com/emscripten-core/emsdk.git`

# Enter that directory
cd emsdk

# Fetch the latest version of the emsdk (not needed the first time you clone)
git pull

# Download and install the latest SDK tools.
./emsdk install latest

# Make the "latest" SDK "active" for the current user. (writes .emscripten file)
./emsdk activate latest

# Activate PATH and other environment variables in the current terminal
source ./emsdk_env.sh

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

