#!/bin/bash

cd "gen/proto"
echo "📁 Generating protocol buffers from gen/proto..."
./generate.sh
cd "../.."

echo "🏁 Setup complete!"