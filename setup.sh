#!/usr/bin/env bash

pushd "gen/proto" || exit 1
echo "📁 Generating protocol buffers from gen/proto..."
./generate.sh
popd || exit 1

echo "🏁 Setup complete!"
