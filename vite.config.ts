import purgecss from "@fullhuman/postcss-purgecss";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import autoprefixer from "autoprefixer";
import { exec } from "child_process";
import tailwindcss from "tailwindcss";
import { defineConfig } from "vite";
import { resolve } from "./require.ts";
import config from "./tailwind.config.ts";

// Get current git hash.

const gitHash = await new Promise<string>((resolve, reject) => {
  exec("git rev-parse HEAD", (error, stdout, stderr) => {
    if (error) {
      reject("Error getting git hash, " + error.message);
      return;
    }
    resolve(stdout.trim());
  });
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
      output: {
        entryFileNames: `assets/[name]-${gitHash}.js`,
        chunkFileNames: `assets/[name]-${gitHash}.js`,
        assetFileNames: `assets/[name]-${gitHash}.[ext]`,
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
