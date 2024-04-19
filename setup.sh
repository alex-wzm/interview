#!/bin/bash -e

echo "📁 Generating protocol buffers from gen/proto..."
( cd "gen/proto" && ./generate.sh )

echo "🏁 Setup complete!"
