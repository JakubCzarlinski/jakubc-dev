import purgecss from "@fullhuman/postcss-purgecss";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import autoprefixer from "autoprefixer";
import { readFile } from "fs/promises";
import tailwindcss from "tailwindcss";
import { defineConfig } from "vite";
import { resolve } from "./require.ts";
import config from "./tailwind.config.ts";

const hash = await new Promise<string>((r) => {
  // Read the hash from ./hash.txt
  readFile(resolve(__dirname, "./hash.txt"), "utf-8")
    .then((hash) => r(hash.trim()))
    .catch(() => r(""));
});

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
      maxParallelFileOps: 128,
      output: {
        entryFileNames: `assets/[name]-${hash}.js`,
        chunkFileNames: `assets/[name]-${hash}.js`,
        assetFileNames: `assets/[name]-${hash}.[ext]`,
        manualChunks(id) {
          if (id.includes("index") || id.includes("index-client")) {
            return "index";
          }
          if (
            (id.includes("node_modules") &&
              id.includes("svelte") &&
              (id.includes("internal") ||
                id.includes("store") ||
                id.includes("transition"))) ||
            id.includes("index-client")
          ) {
            return "index";
          }
        },
      },
    },
  },
});
