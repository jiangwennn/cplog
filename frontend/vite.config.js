import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import commonjs from 'vite-plugin-commonjs'
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    commonjs(),
  ],
  css: {
    postcss: {
      plugins: [tailwindcss, autoprefixer]
    }
  },
  resolve: {
    alias: {
      '/': path.resolve(__dirname, '.'),
      '@': path.resolve(__dirname, 'src'),
      '_wailsjs': path.resolve(__dirname, 'wailsjs'),
    }
  }
})
