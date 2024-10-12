import purgecss from "@fullhuman/postcss-purgecss";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";
import { defineConfig } from "vite";
import { resolve } from "./require.ts";
import config from "./tailwind.config.ts";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    svelte({ compilerOptions: { modernAst: true } }),
    // @ts-ignore
    purgecss({}),
  ],
  css: {
    postcss: {
      plugins: [tailwindcss({ config }), autoprefixer()],
    },
  },
  resolve: {
    alias: {
      "@": resolve(__dirname, "./"),
    },
  },
  // mode: "development",
  build: {
    // minify: false,
    rollupOptions: {
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`,
        manualChunks(id) {
          if (
            (id.includes("node_modules") &&
              id.includes("svelte") &&
              (id.includes("internal") ||
                id.includes("store") ||
                id.includes("transition"))) ||
            id.includes("index-client")
          ) {
            return "svelte";
          }
        },
      },
    },
  },
});
