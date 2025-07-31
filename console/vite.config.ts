// Path: /console/vite.config.ts
// Purpose: Configuration file for the Vite development server and build tool.

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    // This is required for file changes to be detected inside a Docker container
    watch: {
      usePolling: true,
    },
  },
})