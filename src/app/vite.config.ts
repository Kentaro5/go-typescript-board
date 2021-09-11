import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from 'path'

console.log(__dirname);
console.log(path.resolve(__dirname, 'main.ts'));
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    lib: {
      fileName: 'app',
      formats: ['es'],
      entry: path.resolve(__dirname, 'main.ts'),
    },
    outDir: 'dist',
    emptyOutDir: false,
  },
  resolve: {
    alias: [
      { find: /^~/, replacement: path.resolve(__dirname) },
      { find: /^\/~/, replacement: path.resolve(__dirname) },
    ],
  },
})
