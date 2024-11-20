import { defineConfig, type RsbuildConfig } from "@rsbuild/core";
import { pluginSvelte } from "@rsbuild/plugin-svelte";
import { resolve } from "./require.ts";

export default defineConfig({
  plugins: [
    pluginSvelte(),
  ],
  html: {
    template: "./index.html",
  },
  source: {
    entry: {
      index: "./project/src/main.ts",
    },
    alias: {
      "@": resolve(__dirname, "./"),
    },
  },
  output: {
    // minify: true,
    distPath: {
      root: "dist",
      html: "./",
      js: "assets",
      jsAsync: "assets",
      css: "assets",
      cssAsync: "assets",
      svg: "assets",
      font: "assets",
      wasm: "assets",
      image: "assets",
      media: "assets",
      assets: "assets",
    },
    filename: {
      js: `[name]-[hash].js`,
      chunk: `[name]-[hash].js`,
      asset: `[name]-[hash].[ext]`,
    },
  },
  performance: {
    buildCache: true,
    removeConsole: true,
    // buildCache: true,
    // removeConsole: true,
    // chunkSplit: {
    //   splitChunks: {
    //     cacheGroups: {
    //       index: {
    //         name: "index",
    //         chunks: "all",
    //         test: (module) => {
    //           const id = module.identifier();
    //           if (id.includes("index") || id.includes("index-client")) {
    //             return true;
    //           }
    //           return (
    //             id.includes("node_modules/svelte/") &&
    //             (id.includes("internal") ||
    //               id.includes("store") ||
    //               id.includes("transition"))
    //           );
    //         },
    //       },
    //     },
    //   },
    // },
  },
} as RsbuildConfig);
