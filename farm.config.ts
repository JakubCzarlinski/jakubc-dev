import { defineConfig, type JsPlugin } from "@farmfe/core";
import farmPluginPostcss from "@farmfe/js-plugin-postcss";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import { readFile } from "fs/promises";
import path from "path";

const hash = await new Promise<string>((r) => {
  // Read the hash from ./hash.txt
  readFile("./hash.txt", "utf-8")
    .then((hash) => r(hash.trim()))
    .catch((e) => console.error(e));
});

const entryFilename = `index.[ext]`;
const assetsFilename = `assets/[name]-${hash}.[ext]`;
const filename = `assets/[name]-${hash}.[ext]`;

function customHash(hash: string): JsPlugin {
  return {
    name: "custom-hash",
    priority: -9999,
    augmentResourceHash: {
      filters: {
        resourcePotTypes: ["js"],
      },
      executor: (param) => {
        return hash;
      },
    },
  };
}

export default defineConfig({
  vitePlugins: [
    svelte({ compilerOptions: { modernAst: true } }),
  ],
  plugins: [farmPluginPostcss({}), customHash(hash)],
  compilation: {
    mode: "production",
    // minify: false,
    external: [
    ],
    sourcemap: false,
    externalNodeBuiltins: true,
    resolve: {
      alias: {
        "@/": path.join(process.cwd(), "./"),
      },
    },
    output: {
      // clean: false,
      entryFilename,
      assetsFilename,
      filename,
    },
    partialBundling: {
      enforceResources: [
        {
          name: "index",
          test: [".+"],
        },
      ],
    },
  },
});
